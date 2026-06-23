package main

import (
	"chirpy/internal/auth"
	"net/http"
	"os"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerDeleteChirp(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		w.WriteHeader(401)
		return
	}
	chirpId := r.PathValue("chirpID")
	uuidChirpId, err := uuid.Parse(chirpId)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	secretForToken := os.Getenv("JWT_SECRET")
	user, err := auth.ValidateJWT(token, secretForToken)
	if err != nil {
		w.WriteHeader(401)
		return
	}

	chirp, err := cfg.db.GetChirp(r.Context(), uuidChirpId)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	if chirp.UserID != user {
		w.WriteHeader(403)
		return
	}

	err = cfg.db.DeleteChirpById(r.Context(), chirp.ID)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(204)
}
