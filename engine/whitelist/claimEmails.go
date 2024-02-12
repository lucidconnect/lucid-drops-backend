package whitelist

import (
	"errors"
	"time"

	"github.com/lucidconnect/inverse/emails"
	"github.com/lucidconnect/inverse/engine"
	"github.com/lucidconnect/inverse/graph/model"
	"github.com/lucidconnect/inverse/models"
	"github.com/lucidconnect/inverse/utils"
	"github.com/rs/zerolog/log"
)

const (
	//emailOTPttl is the number of minutes the Email OTP will be valid
	emailOTPttl          = 10
	maximumOTPReattempts = 3
)

func StartEmailVerificationForClaim(input *model.EmailClaimInput) (*model.StartEmailVerificationResponse, error) {
	item, err := engine.GetItemByID(input.ItemID)
	if err != nil {
		return nil, errors.New("items not found")
	}

	if item.ClaimDeadline != nil {
		if time.Now().After(*item.ClaimDeadline) {
			return nil, errors.New("the item is no longer available to be claimed")
		}
	}

	if item.Criteria == nil {
		return nil, errors.New("the item can only be claimed via Email Verification")
	}

	switch *item.Criteria {
	case model.ClaimCriteriaTypeEmailWhiteList:
		_, err = engine.GetEmailClaimIDByItemAndEmail(&item.ID, input.EmailAddress)
		if err != nil {
			return nil, errors.New("you aren't authorized to claim this item. Please try again")
		}
	case model.ClaimCriteriaTypeEmailDomain:
		_, err = engine.GetEmailClaimIDByItemAndEmailSubDomain(&item.ID, input.EmailAddress)
		if err != nil {
			return nil, errors.New("you aren't authorized to claim this item. Please try again")
		}
	default:
		return nil, errors.New("the item can only be claimed via Email Verification")
	}

	generatedOTP := utils.RandomNumericRunes(5)

	action := "verification"
	expiry := "ten minutes"
	from := "verify@getabacus.app"

	err = emails.SendVerificationEmail(input.EmailAddress, from, generatedOTP, action, expiry)
	if err != nil {
		log.Err(err)
		return nil, errors.New("an error occurred while sending the verification email, please try again")
	}

	collection, err := engine.GetCollectionByID(item.CollectionID.String())
	if err != nil {
		log.Err(err)
		return nil, err
	}

	items, err := engine.GetCollectionItems(item.CollectionID.String())
	if err != nil {
		log.Err(err)
		return nil, err
	}

	// TODO use DB order or smart contract deploys to persist this on the item level
	var ItemIdOnContract int64
	for idx, collectionItem := range items {
		if collectionItem.ID.String() == input.ItemID {
			ItemIdOnContract = int64((idx) + 1)
		}
	}

	var smartContractAddress string
	if collection.AAContractAddress != nil {
		smartContractAddress = *collection.AAContractAddress
	}

	newEmailOTP := &models.EmailOTP{
		IssuedAt:                  time.Now().Unix(),
		ExpiresAt:                 time.Now().Add(time.Minute * time.Duration(emailOTPttl)).Unix(),
		ItemID:                    item.ID,
		UserEmail:                 input.EmailAddress,
		ExpectedOTP:               generatedOTP,
		Attempts:                  0,
		ItemIdOnContract:          ItemIdOnContract,
		CollectionContractAddress: smartContractAddress,
	}

	err = engine.CreateModel(newEmailOTP)
	if err != nil {
		log.Info().Msgf("🚨 Claim Model failed to created in DB %+v", newEmailOTP)
		return nil, errors.New("an error occured when generating your claim.Please contact support")
	}

	return &model.StartEmailVerificationResponse{
		OtpRequestID: newEmailOTP.ID.String(),
	}, nil
}

func CompleteEmailVerificationForClaim(input *model.CompleteEmailVerificationInput) (*model.CompleteEmailVerificationResponse, error) {
	otpDetails, err := engine.GetEmailOTPRecordByID(input.OtpRequestID)
	if err != nil {
		return nil, errors.New("email verification attempt could not be found")
	}

	if otpDetails.VerifiedAt != nil {
		return nil, errors.New("email has already been verified please procced to claim page")
	}

	if otpDetails.Attempts >= maximumOTPReattempts {
		return nil, errors.New("email verification attempts have been exceded")
	}

	timeToTime := time.Unix(otpDetails.ExpiresAt, 0)
	if time.Now().After(timeToTime) {
		return nil, errors.New("email verification OTP has expired try again")
	}

	if otpDetails.ExpectedOTP != input.Otp {
		otpDetails.Attempts++
		otpSaveError := engine.SaveModel(otpDetails)
		if otpSaveError != nil {
			log.Info().Msgf("🚨 OTP Model failed to updated in DB %+v", otpDetails)
			return nil, errors.New("an error occured when verifying the OTP")
		}

		return nil, errors.New("otp is invalid try again")
	}

	item, err := engine.GetItemByID(otpDetails.ItemID.String())
	if err != nil {
		return nil, errors.New("item not found")
	}

	now := time.Now().Unix()
	otpDetails.VerifiedAt = &now
	otpSaveError := engine.SaveModel(otpDetails)
	if otpSaveError != nil {
		log.Info().Msgf("🚨 OTP Model failed to updated in DB %+v", otpDetails)
		return nil, errors.New("an error when verifying the OTP")
	}

	if item.TokenID == nil {
		return nil, errors.New("The requested item is not ready to be claimed, please try again in a few minutes")
	}

	passResp, err := CreateMintPassForValidatedCriteriaItem(item.ID.String())
	if err != nil {
		return nil, errors.New("error creating mint pass")
	}

	return &model.CompleteEmailVerificationResponse{
		OtpRequestID: *passResp.PassID,
	}, nil
}
