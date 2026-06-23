package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerGetAllChirps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	chirpsArray := []ValidResponse{}

	authorIDString := r.URL.Query().Get("author_id")
	sortString := r.URL.Query().Get("sort")

	if authorIDString != "" {
		authorIDUUid, err := uuid.Parse(authorIDString)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		allChirps, err := cfg.db.GetChirpByID(r.Context(), authorIDUUid)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		if sortString == "desc" {
			sort.Slice(allChirps, func(i, j int) bool { return allChirps[i].CreatedAt.After(allChirps[j].CreatedAt) })
		} else {
			sort.Slice(allChirps, func(i, j int) bool { return allChirps[i].CreatedAt.Before(allChirps[j].CreatedAt) })
		}

		for _, chirp := range allChirps {
			validChirp := ValidResponse{
				ID:         chirp.ID,
				Created_at: chirp.CreatedAt,
				Updated_at: chirp.UpdatedAt,
				Body:       chirp.Body,
				User_id:    chirp.UserID,
			}

			chirpsArray = append(chirpsArray, validChirp)
		}

		dat, err := json.Marshal(chirpsArray)
		if err != nil {
			errResp := ErrorResponse{
				Error: "Error marshalling all chirps",
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

		w.WriteHeader(200)
		w.Write(dat)
		return
	}

	chirps, err := cfg.db.GetAllChirps(r.Context())
	if err != nil {
		errResp := ErrorResponse{
			Error: "Error getting all chirps",
		}
		dat, err := json.Marshal(errResp)
		if err != nil {
			log.Printf("Error marshalling JSON: %s", err)
			return
		}
		w.WriteHeader(500)
		w.Write(dat)
		return
	}

	if sortString == "desc" {
		sort.Slice(chirps, func(i, j int) bool { return chirps[i].CreatedAt.After(chirps[j].CreatedAt) })
	} else {
		sort.Slice(chirps, func(i, j int) bool { return chirps[i].CreatedAt.Before(chirps[j].CreatedAt) })
	}

	for _, chirp := range chirps {
		validChirp := ValidResponse{
			ID:         chirp.ID,
			Created_at: chirp.CreatedAt,
			Updated_at: chirp.UpdatedAt,
			Body:       chirp.Body,
			User_id:    chirp.UserID,
		}

		chirpsArray = append(chirpsArray, validChirp)
	}

	dat, err := json.Marshal(chirpsArray)
	if err != nil {
		errResp := ErrorResponse{
			Error: "Error marshalling all chirps",
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

	w.WriteHeader(200)
	w.Write(dat)
}
