package res

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, res any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(res)
}
