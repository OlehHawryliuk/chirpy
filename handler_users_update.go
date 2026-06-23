package main

import (
	"chirpy/internal/auth"
	"chirpy/internal/database"
	"encoding/json"
	"net/http"
	"os"
)

type UpdateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (cfg *apiConfig) handlerUsersUpdate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		w.WriteHeader(401)
		return
	}

	tokenSecret := os.Getenv("JWT_SECRET")
	userId, err := auth.ValidateJWT(token, tokenSecret)
	if err != nil {
		w.WriteHeader(401)
		return
	}

	loginParams := UpdateRequest{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&loginParams)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	hashedPassword, err := auth.HashPassword(loginParams.Password)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	newUserParams := database.UpdateUserEmailAndPasswordParams{
		Email:          loginParams.Email,
		HashedPassword: hashedPassword,
		ID:             userId,
	}

	user, err := cfg.db.UpdateUserEmailAndPassword(r.Context(), newUserParams)

	readyToSend := User{
		Email:       user.Email,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		ID:          user.ID,
		IsChirpyRed: user.IsChirpyRed.Bool,
	}

	dat, err := json.Marshal(readyToSend)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(200)
	w.Write(dat)
}
