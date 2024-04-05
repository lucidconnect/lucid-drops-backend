package route

import (
	"net/http"
)

func (s *Server) MetadataHandler(w http.ResponseWriter, r *http.Request) {
	// contractAddress := chi.URLParam(r, "contract")
	// itemId := chi.URLParam(r, "itemid")

	// item, err := s.nftRepo.FindDropById(itemId)
	// if err != nil {
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")
	// if err := json.NewEncoder(w).Encode(item.ToOpenSeaMetadata()); err != nil {
	// 	return
	// }
}
