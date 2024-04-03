package onboarding

import (
	"github.com/lucidconnect/inverse/engine"
	"github.com/lucidconnect/inverse/graph/model"
	"github.com/lucidconnect/inverse/internal"
	"github.com/lucidconnect/inverse/services"
)

func FetchItemCreatorByDropId(dropID string) (*model.CreatorDetails, error) {
	// add cache
	drop, err := engine.GetDropByID(dropID)
	if err != nil {
		return nil, err
	}

	creator, err := engine.GetCreatorByID(drop.CreatorID.String())
	if err != nil {
		return nil, err
	}
	return creator.ToGraphData(), nil
}

func EditUserProfile(input model.EditUserProfileInputType, authDetails *internal.AuthDetails) (*model.UserProfileType, error) {
	creator, err := engine.GetCreatorByAddress(authDetails.Address)
	if err != nil {
		return nil, err
	}

	signer, err := engine.GetAltSignerByCreatorID(creator.ID.String())
	if err != nil {
		return nil, err
	}

	if input.InverseUsername != nil {
		creator.InverseUsername = input.InverseUsername
	}

	if input.Bio != nil {
		creator.Bio = input.Bio
	}

	if input.Image != nil {
		creator.Image = input.Image
	}

	if input.Thumbnail != nil {
		creator.Thumbnail = input.Thumbnail
	}

	if input.Socials != nil {
		if input.Socials.Twitter != nil {
			creator.Twitter = input.Socials.Twitter
		}

		if input.Socials.Instagram != nil {
			creator.Instagram = input.Socials.Instagram
		}

		if input.Socials.Github != nil {
			creator.Github = input.Socials.Github
		}

		if input.Socials.Warpcast != nil {
			creator.Warpcast = input.Socials.Warpcast
		}
	}
	if input.AaWallet != nil {
		creator.AAWalletAddress = *input.AaWallet
		signer.WalletAddress = *input.AaWallet
	}
	err = engine.SaveModel(creator)
	if err != nil {
		return nil, err
	}
	err = engine.SaveModel(signer)
	if err != nil {
		return nil, err
	}

	return creator.CreatorToProfileData(), nil
}

func FindOrCreateStripeCustomerID(authDetails *internal.AuthDetails) (*string, error) {
	creator, err := engine.GetCreatorByAddress(authDetails.Address)
	if err != nil {
		return nil, err
	}

	if creator.StripeCustomerID != nil {
		return creator.StripeCustomerID, nil
	}

	// create stripe customer
	customer, err := services.CreateStripeCustomerID(creator.ID.String())
	if err != nil {
		return nil, err
	}

	creator.StripeCustomerID = &customer.ID
	err = engine.SaveModel(creator)
	if err != nil {
		return nil, err
	}

	return creator.StripeCustomerID, nil
}
