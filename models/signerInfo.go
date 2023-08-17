package models

import (
	"github.com/ethereum/go-ethereum/common"
	"inverse.so/graph/model"
)

type SignerInfo struct {
	Base
	CreatorID     string         `gorm:"type:uuid;index:idx_creator_id,unique;not null;"`
	WalletAddress common.Address `gorm:"type:varchar(255);index:idx_wallet_address,unique;not null;"`
	Signature     *string
	AltPublicKey  string
	AltPrivateKey string
	Provider      model.SignerProvider
}
