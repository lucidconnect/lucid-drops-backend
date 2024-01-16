package onboarding

import (
	"inverse.so/engine"
	"inverse.so/graph/model"
	"inverse.so/internal"
	"inverse.so/services"
)

func FetchItemCreatorByCollectionId(collectionID string) (*model.CreatorDetails, error) {
	// add cache
	collection, err := engine.GetCollectionByID(collectionID)
	if err != nil {
		return nil, err
	}

	creator, err := engine.GetCreatorByID(collection.CreatorID.String())
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

	err = engine.SaveModel(creator)
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
