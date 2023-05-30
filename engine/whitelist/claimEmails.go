package whitelist

import (
	"errors"
	"time"

	"github.com/rs/zerolog/log"
	"inverse.so/emails"
	"inverse.so/engine"
	"inverse.so/graph/model"
	"inverse.so/models"
	"inverse.so/utils"
)

const (
	//emailOTPttl is the number of minutes the Email OTP will be valid
	emailOTPttl = 10
)

func StartEmailVerificationForClaim(input *model.EmailClaimInput) (*model.StartEmailVerificationResponse, error) {
	item, err := engine.GetItemByID(input.ItemID)
	if err != nil {
		return nil, errors.New("items not found")
	}

	if item.Criteria == nil || !(*item.Criteria == model.ClaimCriteriaTypeEmailDomain || *item.Criteria == model.ClaimCriteriaTypeEmailWhiteList) {
		return nil, errors.New("the item can only be claimed via Email Verification")
	}

	_, err = engine.GetEmailClaimIDByItemAndEmail(&item.ID, input.EmailAddress)
	if err != nil {
		return nil, errors.New("you aren't authorized to claim this item. Please try again")
	}

	// TODO add subdomain claim
	// claim, err = engine.GetEmailClaimIDByItemAndEmail(&item.ID, *&input.ClaimingAddress) if err != nil {
	// 	return nil, errors.New("you aren't authorized to claim this item. Please try again")
	// }

	generatedOTP := utils.RandomNumericRunes(5)

	action := "onboarding"
	expiry := "ten minutes"
	from := "verify@getabacus.app"

	err = emails.SendVerificationEmail(input.EmailAddress, from, generatedOTP, action, expiry)
	if err != nil {
		log.Err(err)
		return nil, errors.New("an error occurred while sending the verification email, please try again")
	}

	newEmailOTP := &models.EmailOTP{
		IssuedAt:        time.Now().Unix(),
		ExpiresAt:       time.Now().Add(time.Minute * time.Duration(emailOTPttl)).Unix(),
		ItemID:          item.ID,
		UserEmail:       input.EmailAddress,
		ClaimingAddress: input.ClaimingAddress,
		ExpectedOTP:     generatedOTP,
		Attempts:        0,
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
