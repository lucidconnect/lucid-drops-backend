package route

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog/log"
)

type JsonMetadata struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Address     string `json:"address"`
	Image       string `json:"image"`
	Properties  any    `json:"properties"`
}

func (s *Server) MetadataHandler(w http.ResponseWriter, r *http.Request) {
	dropId := chi.URLParam(r, "dropId")
	tokenId := chi.URLParam(r, "id")

	metadata, err := s.nftRepo.ReadMetadataByDropId(dropId, tokenId)
	if err != nil {
		log.Err(err).Caller().Send()
		return
	}

	item, err := s.nftRepo.FindItemById(metadata.ItemId.String())
	if err != nil {
		log.Err(err).Caller().Send()
		return
	}

	jsonMd := JsonMetadata{
		Name:        metadata.Name,
		Description: metadata.Description,
		Image:       item.Image,
		Properties:  metadata.Properties,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(jsonMd); err != nil {
		return
	}
}
