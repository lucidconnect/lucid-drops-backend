package models

import (
	"github.com/lucidconnect/inverse/graph/model"
	"github.com/lucidconnect/inverse/utils"
)

type Creator struct {
	Base
	WalletAddress         string // embedded wallet address
	ExternalWalletAddress string
	AAWalletAddress       string
	InverseUsername       *string      `gorm:"default:null"`
	Thumbnail             *string      `gorm:"default:null"`
	FirstPayment          bool         `gorm:"default:false"`
	Image                 *string      `gorm:"default:null"`
	Bio                   *string      `gorm:"default:null"`
	Twitter               *string      `gorm:"default:null"`
	Instagram             *string      `gorm:"default:null"`
	Github                *string      `gorm:"default:null"`
	Warpcast              *string      `gorm:"default:null"`
	StripeCustomerID      *string      `gorm:"default:null"`
	SignerInfo            []SignerInfo `gorm:"foreignKey:CreatorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (c *Creator) ToGraphData() *model.CreatorDetails {
	return &model.CreatorDetails{
		Address:         c.WalletAddress,
		CreatorID:       c.ID.String(),
		InverseUsername: c.InverseUsername,
		FirstPayment:    c.FirstPayment,
		AaWallet: &c.AAWalletAddress,
	}
}

func (c *Creator) CreatorToProfileData() *model.UserProfileType {
	return &model.UserProfileType{
		CreatorID:       utils.GetStrPtr(c.ID.String()),
		InverseUsername: c.InverseUsername,
		Thumbnail:       c.Thumbnail,
		Image:           c.Image,
		Bio:             c.Bio,
		Socials: &model.Socials{
			Twitter:   c.Twitter,
			Instagram: c.Instagram,
			Github:    c.Github,
			Warpcast:  c.Warpcast,
		},
		AaWallet: &c.AAWalletAddress,
	}
}
