package route

import (
	"net/http"

	"github.com/rs/zerolog/log"
	"inverse.so/engine/whitelist"
)

func TelegramCallBack(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	log.Info().Msgf("TelegramCallBack: %+v", *r)
	userId := r.URL.Query().Get("id")
	userName := r.URL.Query().Get("username")
	photoURL := r.URL.Query().Get("photo_url")
	hash := r.URL.Query().Get("hash")

	// TODO: check if the token is valid
	authID, err := whitelist.ProcessTelegramCallBack(userId, userName, hash, photoURL)
	if err != nil {
		log.Error().Msgf("TelegramCallBack: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "http://localhost:3000/whitelist/telegram/"+*authID, http.StatusFound)
}
