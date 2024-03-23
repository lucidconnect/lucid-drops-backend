package drops

import (
	"errors"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lucidconnect/inverse/dbutils"
	"github.com/lucidconnect/inverse/engine"
	"github.com/lucidconnect/inverse/graph/model"
	"github.com/lucidconnect/inverse/internal"
	"github.com/lucidconnect/inverse/mintwatcher"
	"github.com/lucidconnect/inverse/models"
	"github.com/lucidconnect/inverse/services"
	"github.com/lucidconnect/inverse/utils"
	"github.com/rs/zerolog/log"
)

func TempCreateItem(input *model.ItemInput, authDetails *internal.AuthDetails) (*model.Item, error) {
	creator, err := engine.GetCreatorByAddress(authDetails.Address)
	if err != nil {
		return nil, errors.New("creator has not been onboarded to create a new drop")
	}

	if input.Name == nil || input.Image == nil || input.DropID == nil || input.Description == nil {
		return nil, errors.New("pass in all Fields inorder to create a new item")
	}

	drop, err := engine.GetDropByID(*input.DropID)
	if err != nil {
		return nil, errors.New("drop not found")
	}

	if drop.CreatorID != creator.ID {
		return nil, errors.New("drop doesn't belong to the creator if the item")
	}

	newItem := &models.Item{
		Name:         *input.Name,
		Image:        *input.Image,
		Description:  *input.Description,
		DropID:       drop.ID,
		UserLimit:    input.UserLimit,
		EditionLimit: input.EditionLimit,
	}

	if input.ClaimFee != nil {
		newItem.ClaimFee = *input.ClaimFee
	}

	err = engine.CreateModel(newItem)
	if err != nil {
		return nil, errors.New("couldn't create new drop")
	}

	return newItem.ToGraphData(), nil
}

func CreateItem(input *model.ItemInput, authDetails *internal.AuthDetails) (*model.Item, error) {
	creator, err := engine.GetCreatorByAddress(authDetails.Address)
	if err != nil {
		return nil, errors.New("creator has not been onboarded to create a new drop")
	}

	if input.Name == nil || input.Image == nil || input.DropID == nil || input.Description == nil {
		return nil, errors.New("pass in all Fields inorder to create a new item")
	}

	drop, err := engine.GetDropByID(*input.DropID)
	if err != nil {
		return nil, errors.New("drop not found")
	}

	if drop.CreatorID != creator.ID {
		return nil, errors.New("drop doesn't belong to the creator if the item")
	}

	newItem := &models.Item{
		Name:         *input.Name,
		Image:        *input.Image,
		Description:  *input.Description,
		DropID:       drop.ID,
		UserLimit:    input.UserLimit,
		EditionLimit: input.EditionLimit,
	}

	if input.ClaimFee != nil {
		newItem.ClaimFee = *input.ClaimFee
	}
	tokenId := int64(1)
	newItem.TokenID = &tokenId

	err = engine.CreateModel(newItem)
	if err != nil {
		return nil, errors.New("couldn't create new drop")
	}

	// This was removed because all deployments are now triggered by the FE
	// go func() {
	// 	inverseAAServerURL := utils.UseEnvOrDefault("AA_SERVER", "https://inverse-aa.onrender.com")
	// 	inverseAPIBaseURL := utils.UseEnvOrDefault("API_BASE_URL", "https://inverse-backend.onrender.com")
	// 	client := &http.Client{}
	// 	if drop.AAContractAddress == nil {
	// 		log.Info().Msg("🪼TODO ADD SUPPORT FOR QUEING")
	// 		return
	// 	}
	// 	itemData, err := json.Marshal(map[string]interface{}{
	// 		"image":           fmt.Sprintf("%s/metadata/%s/%s", inverseAPIBaseURL, *drop.AAContractAddress, newItem.ID.String()),
	// 		"contractAddress": *drop.AAContractAddress,
	// 		"Network":         drop.BlockchainNetwork,
	// 	})

	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}

	// 	req, err := http.NewRequest(http.MethodPost, inverseAAServerURL+"/additem", bytes.NewBuffer(itemData))
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}

	// 	req.Header.Add("Content-Type", "application/json")
	// 	res, err := client.Do(req)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}

	// 	var isBase bool
	// 	if drop.BlockchainNetwork != nil {
	// 		isBase = *drop.BlockchainNetwork == model.BlockchainNetworkBase
	// 	}

	// 	if res.StatusCode == http.StatusOK {
	// 		go func() {
	// 			tokenID, err := FetchTokenUri(*drop.AAContractAddress, newItem.ID.String(), isBase)
	// 			if err != nil {
	// 				return
	// 			}

	// 			if tokenID == nil {
	// 				log.Info().Msgf("🚨 Token ID not found for Item %s", newItem.ID)
	// 				return
	// 			}

	// 			tokenIDint64 := int64(*tokenID)
	// 			newItem.TokenID = &tokenIDint64
	// 			err = engine.SaveModel(&newItem)
	// 			if err != nil {
	// 				log.Error().Msg(err.Error())
	// 			}
	// 		}()
	// 	}
	// 	defer res.Body.Close()
	// }()

	return newItem.ToGraphData(), nil
}

func DeleteItem(itemID string, authDetails *internal.AuthDetails) (*model.Item, error) {
	creator, err := engine.GetCreatorByAddress(authDetails.Address)
	if err != nil {
		return nil, errors.New("creator has not been onboarded to create a new item")
	}

	item, err := engine.GetItemByID(itemID)
	if err != nil {
		return nil, errors.New("drop not found")
	}

	drop, err := engine.GetDropByID(item.DropID.String())
	if err != nil {
		return nil, errors.New("drop not found")
	}

	if creator.ID != drop.CreatorID {
		return nil, errors.New("the drop doesn't belong to this creator")
	}

	err = engine.SoftDeleteModel(item)
	if err != nil {
		return nil, errors.New("couldn't delete the item")
	}

	return item.ToGraphData(), nil
}

func UpdateItem(itemID string, input *model.ItemInput, authDetails *internal.AuthDetails) (*model.Item, error) {
	creator, err := engine.GetCreatorByAddress(authDetails.Address)
	if err != nil {
		return nil, errors.New("creator has not been onboarded to create a new item")
	}

	item, err := engine.GetItemByID(itemID)
	if err != nil {
		return nil, errors.New("drop not found")
	}

	if input.Name != nil {
		item.Name = *input.Name
	}

	if input.Image != nil {
		item.Image = *input.Image
	}

	if input.Description != nil {
		item.Description = *input.Description
	}

	if input.DropID != nil {
		drop, err := engine.GetDropByID(*input.DropID)
		if err != nil {
			return nil, errors.New("drop not found")
		}

		if creator.ID != drop.CreatorID {
			return nil, errors.New("the drop doesn't belong to this creator")
		}

		item.DropID = drop.ID
	}

	err = engine.SaveModel(item)
	if err != nil {
		return nil, errors.New("couldn't create new item")
	}

	return item.ToGraphData(), nil
}

func FetchAuthotizedSubdomainsForItem(itemID string) ([]string, error) {
	item, err := engine.GetItemByID(itemID)
	if err != nil {
		return nil, errors.New("items not found")
	}

	if !item.ShowEmailDomainHints || item.Criteria == nil || *item.Criteria != model.ClaimCriteriaTypeEmailDomain {
		return nil, nil
	}

	subdomains, err := engine.GetAuthorizedSubdomainsForItem(itemID)
	if err != nil {
		return nil, errors.New("authorized subdomains not found")
	}

	mappedDomains := make([]string, len(subdomains))
	for idx, subDomainEntry := range subdomains {
		mappedDomains[idx] = subDomainEntry.BaseDomain
	}

	return mappedDomains, nil
}

func FetchDropItems(dropID string, includeDelelted bool, authDetails *internal.AuthDetails) ([]*model.Item, error) {
	// All Drop data will be public for now

	// creator, err := engine.GetCreatorByAddress(authDetails.Address)
	// if err != nil {
	// 	return nil, errors.New("creator has not been onboarded")
	// }

	var err error
	var items []*models.Item

	if includeDelelted {
		items, err = engine.GetDropItemsIncludeDeleted(dropID)
		if err != nil {
			return nil, errors.New("items not found")
		}

	} else {
		items, err = engine.GetDropItems(dropID)
		if err != nil {
			return nil, errors.New("items not found")
		}
	}

	mappedItems := make([]*model.Item, len(items))
	for idx, item := range items {
		// mintPasses, _ := FetchMintPassesForDrops(dropID)
		// item.MintPasses = mintPasses
		mappedItems[idx] = item.ToGraphData()
	}

	return mappedItems, nil
}

func FetchItemByID(itemID string) (*model.Item, error) {
	item, err := engine.GetItemByID(itemID)
	if err != nil {
		return nil, errors.New("item not found")
	}

	// mintPasses, _ := FetchMintPassesForItems(item.ID.String())
	// item.MintPasses = mintPasses

	return item.ToGraphData(), nil
}

func FetchQuestionsByItemID(itemID string) ([]*model.QuestionnaireType, error) {
	item, err := engine.GetItemByID(itemID)
	if err != nil {
		return nil, errors.New("item not found")
	}

	questions, err := engine.GetItemQuestionsByItem(item)
	if err != nil {
		return nil, errors.New("questions not found")
	}

	return questions, nil
}

func FetchFeaturedItems() ([]*model.Item, error) {
	items, err := engine.GetFeaturedItems()
	if err != nil {
		return nil, errors.New("items not found")
	}

	mappedItems := make([]*model.Item, len(items))

	for idx, item := range items {
		mappedItems[idx] = item.ToGraphData()
	}

	return mappedItems, nil
}

func SetItemClaimDeadline(itemID string, deadline string) (*model.Item, error) {
	item, err := engine.GetItemByID(itemID)
	if err != nil {
		return nil, errors.New("item not found")
	}

	dateForrmatted, err := time.Parse(time.RFC3339Nano, deadline)
	if err != nil {
		return nil, err
	}

	item.ClaimDeadline = &dateForrmatted
	err = engine.SaveModel(item)
	if err != nil {
		return nil, errors.New("couldn't save item")
	}

	return item.ToGraphData(), nil
}

func FetchMintPassesForDrops(dropId string) ([]models.MintPass, error) {

	var mintPasses []models.MintPass
	err := dbutils.DB.Model(&models.MintPass{}).Where("drop_id = ?", dropId).Find(&mintPasses).Error
	if err != nil {
		return nil, err
	}

	return mintPasses, nil
}

func FetchTokenUri(contractAddress, itemID string, isBase bool) (*int, error) {
	inverseAPIBaseURL := os.Getenv("API_BASE_URL")
	rpcProvider := utils.UseEnvOrDefault("POLYGON_RPC_PROVIDER", "https://polygon-mainnet.g.alchemy.com/v2/wH3GkDxLOS4h8O7hmIPWqvmOvE4VIqWn")
	if isBase {
		rpcProvider = utils.UseEnvOrDefault("BASE_RPC_PROVIDER", "https://base-mainnet.g.alchemy.com/v2/2jx1c05x5vFN7Swv9R_ZJKKAXZUfas8A")
	}

	client, err := ethclient.Dial(rpcProvider)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}

	addressToAddress := common.HexToAddress(contractAddress)
	x, err := mintwatcher.NewMintwatcher(addressToAddress, client)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}

	opts := &bind.CallOpts{}
	// TODO make counter more dynamic
	for i := 1; i <= 10; i++ {
		expectedURI := fmt.Sprintf("%s/metadata/%s/%s", inverseAPIBaseURL, contractAddress, itemID)
		integer := big.NewInt(int64(i))
		uri, err := x.Uri(opts, integer)
		if err != nil {
			log.Error().Msg(err.Error())
			log.Info().Msgf("🔍 Fetching token URI for %s/%s", contractAddress, itemID)
		}

		if uri == expectedURI {
			return &i, nil
		}
	}
	return nil, errors.New("token ID not found")
}

func FetchNftHolders(item *model.Item) ([]string, error) {
	var alchemyOpts []services.Option
	apiKeyOpt := services.WithApiKey(os.Getenv("ALCHEMY_API_KEY"))
	urlOpt := services.WithUrl(os.Getenv("ALCHEMY_URL"))
	alchemyOpts = append(alchemyOpts, apiKeyOpt, urlOpt)
	alchemyClient, err := services.NewAlchemyClient(alchemyOpts...)
	if err != nil {
		log.Err(err).Send()
		return nil, err
	}

	holders, err := alchemyClient.GetOwnersForNft(item.DropAddress, "1")
	if err != nil {
		log.Err(err).Send()
		return nil, err
	}
	return holders, nil
}
