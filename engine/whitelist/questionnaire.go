package whitelist

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm/clause"
	"inverse.so/dbutils"
	"inverse.so/engine"
	"inverse.so/graph/model"
	"inverse.so/internal"
	"inverse.so/models"
)

func CreateQuestionnaireCriteriaForItem(authDetails *internal.AuthDetails, input *model.QuestionnaireCriteriaInput) (*model.Item, error) {
	creator, err := engine.GetCreatorByAddress(authDetails.Address)
	if err != nil {
		return nil, errors.New("creator has not been onboarded to create a new collection")
	}

	item, err := engine.GetItemByID(input.ItemID)
	if err != nil {
		return nil, errors.New("item not found")
	}

	if item.Criteria != nil {
		//Delete Existing criteria
		err := engine.DeleteCriteriaIfExists(item)
		if err != nil {
			return nil, err
		}
	}

	switch input.QuestionType {
	case model.QuestionTypeDirectAnswer:
		if input.OpenEndedInput == nil || len(input.OpenEndedInput) == 0 {
			return nil, errors.New("questionTypeDirectAnswer should contain at least have a OpenEndedInput")
		}

		dbQuestions := make([]*models.DirectAnswerCriteria, len(input.OpenEndedInput))
		for idx, questionInput := range input.OpenEndedInput {
			mappedChoices := make(map[string]bool)
			for _, choice := range questionInput.Answers {
				mappedChoices[strings.ToLower(choice)] = true
			}

			answerBytes, err := json.Marshal(mappedChoices)
			if err != nil {
				return nil, err
			}

			dbQuestions[idx] = &models.DirectAnswerCriteria{
				CreatorID:  creator.ID,
				ItemID:     item.ID,
				Question:   questionInput.Question,
				Answers:    string(answerBytes),
				QuestionID: uuid.NewV4(),
			}
		}

		insertionErr := dbutils.DB.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(dbQuestions, 100).Error
		if insertionErr != nil {
			return nil, insertionErr
		}

		directAnswerQuestionnaireType := model.ClaimCriteriaTypeDirectAnswerQuestionnaire
		item.Criteria = &directAnswerQuestionnaireType
		itemUpdateErr := engine.SaveModel(item)
		if itemUpdateErr != nil {
			return nil, itemUpdateErr
		}

		return item.ToGraphData(), nil

	case model.QuestionTypeMultiChoice:
		if input.MultiChoiceInput == nil || len(input.MultiChoiceInput) == 0 {
			return nil, errors.New("questionTypeMultiChoice should contain at least have MultiChoiceInput")
		}

		dbQuestions := make([]*models.MultiChoiceCriteria, len(input.MultiChoiceInput))
		for idx, questionInput := range input.MultiChoiceInput {
			mappedChoices := make(map[string]bool)
			for _, choice := range questionInput.Choices {
				mappedChoices[strings.ToLower(choice)] = (choice == questionInput.CorrectChoice)
			}

			choicesBytes, err := json.Marshal(mappedChoices)
			if err != nil {
				return nil, err
			}

			dbQuestions[idx] = &models.MultiChoiceCriteria{
				CreatorID:     creator.ID,
				ItemID:        item.ID,
				Question:      questionInput.Question,
				QuestionID:    uuid.NewV4(),
				Choices:       string(choicesBytes),
				CorrectChoice: questionInput.CorrectChoice,
			}
		}

		insertionErr := dbutils.DB.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(dbQuestions, 100).Error
		if insertionErr != nil {
			return nil, insertionErr
		}

		multiChoiceQuestionnaireType := model.ClaimCriteriaTypeMutliChoiceQuestionnaire
		item.Criteria = &multiChoiceQuestionnaireType
		itemUpdateErr := engine.SaveModel(item)
		if itemUpdateErr != nil {
			return nil, itemUpdateErr
		}

		return item.ToGraphData(), nil
	}

	return nil, nil
}

func ValidateQuestionnaireCriteriaForItem(itemID string, input []*model.QuestionnaireAnswerInput) (*string, error) {
	item, err := engine.GetItemByID(itemID)
	if err != nil {
		return nil, errors.New("item not found")
	}

	if item.ClaimDeadline != nil {
		if time.Now().After(*item.ClaimDeadline) {
			return nil, errors.New("the item is no longer available to be claimed")
		}
	}

	if item.Criteria == nil {
		return nil, errors.New("item can be freely claimed")
	}

	switch *item.Criteria {
	case model.ClaimCriteriaTypeDirectAnswerQuestionnaire:
		var directQuestions []*models.DirectAnswerCriteria

		err := dbutils.DB.Where(&models.DirectAnswerCriteria{ItemID: item.ID}).Find(&directQuestions).Error
		if err != nil {
			return nil, errors.New("seems item doesn't have any direct questions")
		}

		if len(directQuestions) != len(input) {
			return nil, fmt.Errorf("provide anwsers for all (%d) questions", len(directQuestions))
		}

		mappedQuestionsAndChoices := make(map[string]map[string]bool, len(directQuestions))
		for _, q := range directQuestions {
			var unmarshelledAnswers map[string]bool

			json.Unmarshal([]byte(q.Answers), &unmarshelledAnswers)

			mappedQuestionsAndChoices[q.QuestionID.String()] = unmarshelledAnswers
		}

		answeredQuestions := make(map[string]bool)

		for _, potentialAnswers := range input {
			answeredQuestions[potentialAnswers.QuestionID] = true

			correctAnswers, found := mappedQuestionsAndChoices[potentialAnswers.QuestionID]
			if !found {
				return nil, fmt.Errorf("(%s) is not part of the item claim questions", potentialAnswers.QuestionID)
			}

			_, correct := correctAnswers[strings.ToLower(potentialAnswers.Answer)]
			if !correct {
				return nil, errors.New("wrong answer supplied for one of the questions")
			}
		}

		if len(answeredQuestions) != len(directQuestions) {
			return nil, errors.New("submitted duplicate questions")
		}

	case model.ClaimCriteriaTypeMutliChoiceQuestionnaire:
		var multiChoiceQuestions []*models.MultiChoiceCriteria

		err := dbutils.DB.Where(&models.MultiChoiceCriteria{ItemID: item.ID}).Find(&multiChoiceQuestions).Error
		if err != nil {
			return nil, errors.New("seems item doesn't have any multi choice questions")
		}

		if len(multiChoiceQuestions) != len(input) {
			return nil, fmt.Errorf("provide anwsers for all (%d) questions", len(multiChoiceQuestions))
		}

		mappedQuestionsAndAnswer := make(map[string]string, len(multiChoiceQuestions))
		for _, q := range multiChoiceQuestions {
			mappedQuestionsAndAnswer[q.QuestionID.String()] = q.CorrectChoice
		}

		answeredQuestions := make(map[string]bool)

		for _, potentialAnswers := range input {
			answeredQuestions[potentialAnswers.QuestionID] = true

			correctAnswer, found := mappedQuestionsAndAnswer[potentialAnswers.QuestionID]
			if !found {
				return nil, fmt.Errorf("(%s) is not part of the item claim", potentialAnswers.QuestionID)
			}

			if !strings.EqualFold(correctAnswer, potentialAnswers.Answer) {
				return nil, fmt.Errorf("wrong choice %s supplied for some of the questions %s", potentialAnswers.Answer, correctAnswer)
			}
		}

		if len(answeredQuestions) != len(multiChoiceQuestions) {
			return nil, errors.New("submitted duplicate questions")
		}
	default:
		return nil, errors.New("item cannot be claimed via this method")
	}

	collection, err := engine.GetCollectionByID(item.CollectionID.String())
	if err != nil {
		log.Err(err)
		return nil, err
	}

	var smartContractAddress string
	if collection.AAContractAddress != nil {
		smartContractAddress = *collection.AAContractAddress
	}

	if item.TokenID == nil {
		return nil, errors.New("The requested item is not ready to be claimed, please try again in a few minutes")
	}

	newMint := models.MintPass{
		ItemId:                    itemID,
		ItemIdOnContract:          *item.TokenID,
		CollectionContractAddress: smartContractAddress,
	}

	err = dbutils.DB.Create(&newMint).Error
	if err != nil {
		return nil, err
	}

	claimingID := newMint.ID.String()

	return &claimingID, nil
}
