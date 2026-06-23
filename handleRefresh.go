package main

import (
	"chirpy/internal/auth"
	"encoding/json"
	"net/http"
	"os"
	"time"
)

func (cfg *apiConfig) handleRefresh(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		w.WriteHeader(401)
		return
	}

	user, err := cfg.db.GetUserFromRefreshToken(r.Context(), token)
	if err != nil {
		w.WriteHeader(401)
		return
	}

	tokenSecret := os.Getenv("JWT_SECRET")
	expirationTime := time.Hour
	jwt, err := auth.MakeJWT(user.ID, tokenSecret, expirationTime)
	if err != nil {
		w.WriteHeader(401)
		return
	}

	response := ResponseRefreshToken{
		RefrToken: jwt,
	}

	dat, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(401)
		return
	}

	w.WriteHeader(200)
	w.Write(dat)
}
