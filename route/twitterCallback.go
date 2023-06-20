package route

import (
	"net/http"

	"github.com/rs/zerolog/log"
	"inverse.so/engine/whitelist"
)

func TwitterCallBack(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	log.Info().Msgf("TwitterCallBack: %+v", *r)
	oauthToken := r.URL.Query().Get("oauth_token")
	oauthVerifier := r.URL.Query().Get("oauth_verifier")

	// TODO: check if the token is valid

	//process claim stuff
	// itemID := chi.URLParam(r, "itemID")
	// twitterHandle := chi.URLParam(r, "twitterHandle")

	// log.Info().Msgf("TwitterCallBack: %s %s %s", itemID, twitterHandle, oauthToken)
	whitelist.ProcessTwitterCallback(&oauthToken, &oauthVerifier)
}
