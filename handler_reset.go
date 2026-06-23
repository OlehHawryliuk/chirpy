package main

import (
	"net/http"
)

func (cfg *apiConfig) resetHandler(w http.ResponseWriter, r *http.Request) {
	if cfg.platform != "dev" {
		w.WriteHeader(403)
		return
	}

	cfg.fileserverHits.Store(0)

	err := cfg.db.DeleteAllUsers(r.Context())
	if err != nil {
		w.WriteHeader(500)

		return
	}
}
