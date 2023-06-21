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
	authID, err := whitelist.ProcessTwitterCallback(&oauthToken, &oauthVerifier)
	if err != nil {
		log.Error().Msgf("TwitterCallBack: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "https://inverse.wtf/whitelist/twitter/"+*authID, http.StatusFound)
}
