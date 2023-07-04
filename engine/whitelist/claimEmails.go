package whitelist

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rs/zerolog/log"
	"inverse.so/emails"
	"inverse.so/engine"
	"inverse.so/graph/model"
	"inverse.so/magic"
	"inverse.so/models"
	"inverse.so/utils"
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

	action := "onboarding"
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
			ItemIdOnContract = int64(len(items) - (idx + 1))
		}
	}

	var smartContractAddress string
	if collection.ContractAddress != nil {
		smartContractAddress = *collection.ContractAddress
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

	now := time.Now().Unix()
	otpDetails.VerifiedAt = &now
	otpSaveError := engine.SaveModel(otpDetails)
	if otpSaveError != nil {
		log.Info().Msgf("🚨 OTP Model failed to updated in DB %+v", otpDetails)
		return nil, errors.New("an error when verifying the OTP")
	}

	return &model.CompleteEmailVerificationResponse{
		OtpRequestID: input.OtpRequestID,
	}, nil
}

func GenerateSignatureForClaim(input *model.GenerateClaimSignatureInput) (*model.MintAuthorizationResponse, error) {
	otpDetails, err := engine.GetEmailOTPRecordByID(input.OtpRequestID)
	if err != nil {
		return nil, errors.New("email verification attempt could not be found")
	}

	if otpDetails.VerifiedAt == nil {
		return nil, errors.New("email has not been verified yet")
	}

	if IsThisAValidEthAddress(input.ClaimingAddress) {
		return nil, errors.New("the passed in address in not a valid Ethereum address")
	}

	otpDetails.ClaimingAddress = input.ClaimingAddress
	otpSaveError := engine.SaveModel(otpDetails)
	if otpSaveError != nil {
		log.Info().Msgf("🚨 OTP Model failed to updated in DB %+v", otpDetails)
		return nil, errors.New("an error when verifying the Claim")
	}

	go func() {
		inverseAAServerURL := utils.UseEnvOrDefault("INVERSE_AA_SERVER", "https://inverse-aa.onrender.com")

		client := &http.Client{}

		itemData, err := json.Marshal(map[string]interface{}{
			"receiptientAddresses": []string{input.ClaimingAddress},
			"items":                []int64{otpDetails.ItemIdOnContract},
			"contractAddress":      otpDetails.CollectionContractAddress,
		})

		if err != nil {
			fmt.Println(err)
			return
		}

		req, err := http.NewRequest(http.MethodPost, inverseAAServerURL+"/sendnfts", bytes.NewBuffer(itemData))
		if err != nil {
			fmt.Println(err)
			return
		}

		req.Header.Add("Content-Type", "application/json")
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer res.Body.Close()
	}()

	// TODO add back signature flow
	dataInPackedFormat := utils.EncodePacked(
		// utils.EncodeAddress("0x14723A09ACff6D2A60DcdF7aA4AFf308FDDC160"),
		utils.EncodeAddress(input.ClaimingAddress), // Some Addresss
		utils.EncodeUint256("123"),
		utils.EncodeBytesString(hex.EncodeToString([]byte("coffee and donuts"))),
		utils.EncodeUint256("1"),
	)

	rawData := hexutil.Encode(dataInPackedFormat)

	keccakOfTheMessageInBytes := crypto.Keccak256(dataInPackedFormat)

	signature := magic.SecretlySignThisMessage("\x19Ethereum Signed Message:\n32" + string(keccakOfTheMessageInBytes))

	return &model.MintAuthorizationResponse{
		PackedData:           rawData,
		MintingAbi:           "['function mint(address _to) public']",
		MintingSignature:     signature,
		SmartContractAddress: "0x34bE7f35132E97915633BC1fc020364EA5134863",
	}, nil
}
