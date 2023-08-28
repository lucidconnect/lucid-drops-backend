package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"inverse.so/graph/model"
	"inverse.so/utils"
)

type Collection struct {
	Base
	CreatorID              uuid.UUID
	CreatorAddress         string
	Name                   string
	Image                  string `json:"image"`
	Thumbnail              string `json:"thumbnail"`
	Description            string `json:"description"`
	AAContractAddress      *string
	ContractAddress        *string
	TransactionHash        *string
	AAWalletDeploymentHash *string
	Featured               bool `gorm:"default:false"`
}

type DeplyomenResponse struct {
	Type          int    `json:"type"`
	ChainID       int    `json:"chainId"`
	Nonce         int    `json:"nonce"`
	To            string `json:"to"`
	Data          string `json:"data"`
	Hash          string `json:"hash"`
	From          string `json:"from"`
	Confirmations int    `json:"confirmations"`
}

func (c *Collection) AfterCreate(tx *gorm.DB) (err error) {
	go func() {
		inverseAAServerURL := utils.UseEnvOrDefault("INVERSE_AA_SERVER", "https://inverse-aa.onrender.com")

		client := &http.Client{}

		collectionData, err := json.Marshal(c)
		if err != nil {
			fmt.Println(err)
			return
		}

		log.Info().Msgf("🪼 Sending Request to AA server at %s and Data : %s", inverseAAServerURL, utils.AsJson(collectionData))

		req, err := http.NewRequest(http.MethodPost, inverseAAServerURL+"/deploy", bytes.NewBuffer(collectionData))
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

		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		var tempDesination DeplyomenResponse
		err = json.Unmarshal(body, &tempDesination)
		if err != nil {
			fmt.Println(err)
			return
		}

		c.TransactionHash = &tempDesination.Hash
		err = tx.Save(c).Error
		if err != nil {
			fmt.Println(err)
			return
		}

	}()
	return nil
}

func (c *Collection) ToGraphData() *model.Collection {
	return &model.Collection{
		ID:              c.ID.String(),
		Name:            c.Name,
		Description:     c.Description,
		Image:           c.Image,
		Thumbnail:       c.Thumbnail,
		ContractAddress: c.ContractAddress,
	}
}
