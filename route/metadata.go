package route

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"inverse.so/engine"
)

func MetadataHandler(w http.ResponseWriter, r *http.Request) {
	// contractAddress := chi.URLParam(r, "contract")
	itemId := chi.URLParam(r, "itemid")

	item, err := engine.GetItemByID(itemId)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(item.ToOpenSeaMetadata()); err != nil {
		return
	}
}
