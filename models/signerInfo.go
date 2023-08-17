package models

import (
	"inverse.so/graph/model"
)

type SignerInfo struct {
	Base
	CreatorID     string `gorm:"type:uuid;index:idx_creator_id,unique;not null;"`
	WalletAddress string `gorm:"type:varchar(255);index:idx_wallet_address,unique;not null;"`
	Signature     *string
	AltPublicKey  string
	AltPrivateKey string
	Provider      model.SignerProvider
}
