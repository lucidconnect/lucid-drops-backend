package models

import (
	"encoding/json"

	uuid "github.com/satori/go.uuid"
	"inverse.so/graph/model"
)

type DirectAnswerCriteria struct {
	BaseWithoutPrimaryKey
	QuestionID uuid.UUID
	ItemID     uuid.UUID `gorm:"primaryKey"`
	CreatorID  uuid.UUID `gorm:"primaryKey"`
	Question   string    `gorm:"primaryKey"`
	// Answers contains a marshalled hashmap of all corrent answers
	Answers   string
	ClaimCode bool `gorm:"default:false"`
}

func (d *DirectAnswerCriteria) ToGraphData() *model.QuestionnaireType {
	question := &model.QuestionnaireType{
		Question:     d.Question,
		QuestionID:   d.QuestionID.String(),
		QuestionType: model.QuestionTypeDirectAnswer,
		ClaimCode:    &d.ClaimCode,
	}

	return question
}

type MultiChoiceCriteria struct {
	BaseWithoutPrimaryKey
	QuestionID uuid.UUID
	ItemID     uuid.UUID `gorm:"primaryKey"`
	CreatorID  uuid.UUID `gorm:"primaryKey"`
	Question   string    `gorm:"primaryKey"`
	// Choices contains a marshalled hashmap of all choices and the correctness
	Choices       string
	CorrectChoice string
}

func (m *MultiChoiceCriteria) ToGraphData() *model.QuestionnaireType {
	var questionsMapping map[string]bool

	question := &model.QuestionnaireType{
		Question:     m.Question,
		QuestionID:   m.QuestionID.String(),
		QuestionType: model.QuestionTypeMultiChoice,
	}

	err := json.Unmarshal([]byte(m.Choices), &questionsMapping)
	if err == nil {
		choices := make([]string, len(questionsMapping))
		i := 0
		for choice := range questionsMapping {
			choices[i] = choice
			i++
		}

		question.Choices = choices
	}

	return question
}
