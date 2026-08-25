package web

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}
func methodGuard(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		http.Error(w, "method not allowed", 405)
		return false
	}
	return true
}
