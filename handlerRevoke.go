package main

import (
	"chirpy/internal/auth"
	"net/http"
)

func (cfg *apiConfig) handleRevoke(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		w.WriteHeader(401)
		return
	}

	_, err = cfg.db.RevokeToken(r.Context(), token)
	if err != nil {
		w.WriteHeader(401)
		return
	}
	w.WriteHeader(204)
}
