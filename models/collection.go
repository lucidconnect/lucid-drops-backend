package models

import (
	"github.com/lucidconnect/inverse/graph/model"
	uuid "github.com/satori/go.uuid"
)

type Drop struct {
	Base
	CreatorID              uuid.UUID
	CreatorAddress         string
	Name                   string
	Image                  string `json:"image"`
	Thumbnail              string `json:"thumbnail"`
	Description            string `json:"description"`
	AAContractAddress      *string
	TransactionHash        *string
	AAWalletDeploymentHash *string
	BlockchainNetwork      *model.BlockchainNetwork
	Featured               bool `gorm:"default:false"`
	MintUrl                string
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

// We nolonger trigger AA-wallet deployments
// func (c *Drop) AfterCreate(tx *gorm.DB) (err error) {
// 	go func() {
// 		inverseAAServerURL := utils.UseEnvOrDefault("AA_SERVER", "https://inverse-aa.onrender.com")

// 		client := &http.Client{}

// 		dropData, err := json.Marshal(c)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 		log.Info().Msgf("🪼 Sending Request to AA server at %s and Data : %s", inverseAAServerURL, utils.AsJson(dropData))

// 		req, err := http.NewRequest(http.MethodPost, inverseAAServerURL+"/deploy", bytes.NewBuffer(dropData))
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 		req.Header.Add("Content-Type", "application/json")
// 		res, err := client.Do(req)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 		defer res.Body.Close()

// 		body, err := io.ReadAll(res.Body)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 		var tempDesination DeplyomenResponse
// 		err = json.Unmarshal(body, &tempDesination)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 		c.TransactionHash = &tempDesination.Hash
// 		err = tx.Save(c).Error
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 	}()
// 	return nil
// }

func (c *Drop) ToGraphData(items []*model.Item) *model.Drop {
	mappedDrop := &model.Drop{
		ID:              c.ID.String(),
		CreatorID:       c.CreatorID.String(),
		CreatedAt:       c.CreatedAt,
		Name:            c.Name,
		Description:     c.Description,
		Image:           c.Image,
		Thumbnail:       c.Thumbnail,
		ContractAddress: c.AAContractAddress,
		Network:         c.BlockchainNetwork,
		MintURL:         c.MintUrl,
	}

	if c.AAContractAddress != nil {
		mappedDrop.ContractAddress = c.AAContractAddress
	}

	if items != nil {
		mappedDrop.Items = items
	}

	return mappedDrop
}
