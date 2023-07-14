package route

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
	"inverse.so/engine/whitelist"
	"inverse.so/utils"
)

func PatreonCallBack(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	log.Info().Msgf("PatreonCallBack: %+v", *r)
	code := r.URL.Query().Get("code")

	// TODO: check if the token is valid
	authID, _, err := whitelist.ProcessPatreonCallback(&code, true)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf(utils.UseEnvOrDefault("FE_BASE_URL", "https://1c5f-89-39-106-222.ngrok-free.app/%s/%s"), "item/criteria/patreon", *authID), http.StatusFound)
}

func PatreonWhitelistCallBack(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	log.Info().Msgf("PatreonCallBack: %+v", *r)
	code := r.URL.Query().Get("code")

	// TODO: check if the token is valid
	authID, _, err := whitelist.ProcessPatreonCallback(&code, false)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf(utils.UseEnvOrDefault("FE_BASE_URL", "https://1c5f-89-39-106-222.ngrok-free.app/%s/%s"), "whitelist/patreon", *authID), http.StatusFound)
}
