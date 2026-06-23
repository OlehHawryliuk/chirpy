package main

import (
	"chirpy/internal/auth"
	"chirpy/internal/database"
	"encoding/json"
	"net/http"
	"time"
)

func (cfg *apiConfig) userLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	defer r.Body.Close()

	loginParams := UserLogin{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&loginParams); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := cfg.db.GetUserByEmail(r.Context(), loginParams.Email)
	if err != nil {
		w.WriteHeader(401)
		return
	}

	passwordMatch, err := auth.CheckPasswordHash(loginParams.Password, user.HashedPassword)
	if err != nil {
		w.WriteHeader(401)
		return
	}

	if !passwordMatch {
		invalidLogin := "Incorrect email or password"
		dat, err := json.Marshal(invalidLogin)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(401)
		w.Write(dat)
		return
	}

	duration := time.Duration(3600) * time.Second
	userToken, err := auth.MakeJWT(user.ID, cfg.secret, duration)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	refToken := auth.MakeRefreshToken()
	refrTokenParams := database.CreateRefreshTokenParams{
		Token:  refToken,
		UserID: user.ID,
	}
	_, err = cfg.db.CreateRefreshToken(r.Context(), refrTokenParams)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := LoginResponse{
		ID:           user.ID,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		Email:        user.Email,
		Token:        userToken,
		RefreshToken: refToken,
		IsChirpyRed:  user.IsChirpyRed.Bool,
	}

	dat, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Write(dat)
}
