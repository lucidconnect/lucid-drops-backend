package route

import (
	"net/http"

	"github.com/rs/zerolog/log"
	"inverse.so/engine/whitelist"
)

func PatreonCallBack(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	log.Info().Msgf("PatreonCallBack: %+v", *r)
	code := r.URL.Query().Get("code")

	// TODO: check if the token is valid
	authID, err := whitelist.ProcessPatreonCallback(&code)
	if err != nil {
		log.Error().Msgf("PatreonCallBack: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "http://localhost:3000/whitelist/patreon/"+*authID, http.StatusFound)
}
