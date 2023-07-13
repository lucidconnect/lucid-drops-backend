package route

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
	"inverse.so/engine/whitelist"
)

func PatreonCallBack(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	log.Info().Msgf("PatreonCallBack: %+v", *r)
	code := r.URL.Query().Get("code")

	// TODO: check if the token is valid
	authID, campaigns, err := whitelist.ProcessPatreonCallback(&code, true)
	if err != nil {
		log.Error().Msgf("PatreonCallBack: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if campaigns != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(campaigns)
	}

	http.Redirect(w, r, "http://localhost:3000/item/criteria/patreon/"+*authID, http.StatusFound)
}

func PatreonWhitelistCallBack(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	log.Info().Msgf("PatreonCallBack: %+v", *r)
	code := r.URL.Query().Get("code")

	// TODO: check if the token is valid
	authID, campaigns, err := whitelist.ProcessPatreonCallback(&code, false)
	if err != nil {
		log.Error().Msgf("PatreonCallBack: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if campaigns != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(campaigns)
	}

	http.Redirect(w, r, "http://localhost:3000/whitelist/patreon/"+*authID, http.StatusFound)
}
