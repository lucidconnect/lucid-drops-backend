package route

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

type TokenURIReq struct {
	ContractAddress string `json:"contractAddress"`
	ItemID          string `json:"itemID"`
	IsBase          bool   `json:"isBase"`
}

func (s *Server) FetchTokenUri(w http.ResponseWriter, r *http.Request) {
	var requestBody TokenURIReq
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		log.Err(err).Caller().Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tokenID, err := fetchTokenUri(requestBody.ContractAddress, requestBody.ItemID, requestBody.IsBase)
	if err != nil {
		return
	}

	if tokenID == nil {
		log.Info().Msgf("🚨 Token ID not found for Item %s", requestBody.ItemID)
		return
	}

	tokenIDint64 := int64(*tokenID)

	if err = json.NewEncoder(w).Encode(tokenIDint64); err != nil {
		log.Err(err).Caller().Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
