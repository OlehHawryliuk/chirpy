package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) getSingleChirp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	chirpId := r.PathValue("chirpID")
	id, err := uuid.Parse(chirpId)

	if err != nil {
		errResp := ErrorResponse{
			Error: "Error parsing uuid from user_id",
		}
		dat, err := json.Marshal(errResp)
		if err != nil {
			log.Printf("Error marshalling JSON: %s", err)
			return
		}
		w.WriteHeader(400)
		w.Write(dat)
		return
	}
	chirp, err := cfg.db.GetChirp(r.Context(), id)
	if err != nil {
		errResp := ErrorResponse{
			Error: "Error getting a chirp",
		}
		dat, err := json.Marshal(errResp)
		if err != nil {
			log.Printf("Error marshalling JSON: %s", err)
			return
		}
		w.WriteHeader(404)
		w.Write(dat)
		return
	}

	validChirp := ValidResponse{
		ID:         chirp.ID,
		Created_at: chirp.CreatedAt,
		Updated_at: chirp.UpdatedAt,
		Body:       chirp.Body,
		User_id:    chirp.UserID,
	}

	dat, err := json.Marshal(validChirp)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		return
	}
	w.WriteHeader(200)
	w.Write(dat)
}
