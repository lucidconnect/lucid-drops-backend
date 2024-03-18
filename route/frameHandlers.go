package route

import (
	"encoding/json"
	"net/http"

	"github.com/lucidconnect/inverse/engine"
	"github.com/lucidconnect/inverse/engine/whitelist"
	"github.com/lucidconnect/inverse/graph/model"
	"github.com/lucidconnect/inverse/models"
	"github.com/rs/zerolog/log"
)

func CreateMintPass(w http.ResponseWriter, r *http.Request) {
	dropId := r.URL.Query().Get("dropId")
	walletAddress := r.URL.Query().Get("wallet")
	drop, _ := engine.GetDropByID(dropId)

	var pass *model.ValidationRespoonse
	var err error
	if drop.FarcasterCriteria != (models.FarcasterCriteria{}) {
		pass, err = whitelist.ValidateFarcasterCriteriaForDrop(walletAddress, dropId)
		if err != nil {
			log.Err(err).Caller().Msg("GenerateSignatureForClaim")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		pass, err = whitelist.CreateMintPassForNoCriteriaDrop(dropId, walletAddress)
		if err != nil {
			log.Err(err).Caller().Msg("GenerateSignatureForClaim")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	if err = json.NewEncoder(w).Encode(pass); err != nil {
		log.Err(err).Caller().Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func CreateMintPassForNoCriteriaItem(w http.ResponseWriter, r *http.Request) {
	dropId := r.URL.Query().Get("dropId")
	walletAddress := r.URL.Query().Get("wallet")
	pass, err := whitelist.CreateMintPassForNoCriteriaDrop(dropId, walletAddress)
	if err != nil {
		log.Err(err).Caller().Msg("GenerateSignatureForClaim")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err = json.NewEncoder(w).Encode(pass); err != nil {
		log.Err(err).Caller().Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GenerateSignatureForClaim(w http.ResponseWriter, r *http.Request) {
	passId := r.URL.Query().Get("passId")
	claimingAddress := r.URL.Query().Get("claimingAddress")
	input := model.GenerateClaimSignatureInput{
		OtpRequestID:    passId,
		ClaimingAddress: claimingAddress,
	}
	sig, err := whitelist.GenerateSignatureForFarcasterClaim(&input)
	if err != nil {
		log.Err(err).Caller().Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(sig); err != nil {
		log.Err(err).Caller().Send()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
