package main

import (
	"chirpy/internal/auth"
	"chirpy/internal/database"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func (cfg *apiConfig) handlerCreatingRecord(w http.ResponseWriter, r *http.Request) {
	badwords := []string{"kerfuffle", "sharbert", "fornax"}
	w.Header().Set("Content-Type", "application/json")

	defer r.Body.Close()

	body := ReturnedBody{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		w.WriteHeader(500)
		errResp := ErrorResponse{
			Error: "Something went wrong",
		}

		dat, err := json.Marshal(errResp)
		if err != nil {
			log.Printf("Error marshalling JSON: %s", err)
			return
		}
		w.Write(dat)

		return
	}

	if len(body.Body) > 140 {
		w.WriteHeader(400)
		errResp := ErrorResponse{
			Error: "Chirp is too long",
		}
		dat, err := json.Marshal(errResp)
		if err != nil {
			log.Printf("Error marshalling JSON: %s", err)
			return
		}
		w.Write(dat)
		return

	}

	noBadWordsStr := strings.Split(body.Body, " ")
	strLow := strings.ToLower(body.Body)
	wordsSlice := strings.Split(strLow, " ")
	for _, badWord := range badwords {
		for i, word := range wordsSlice {
			if strings.ToLower(badWord) == word {
				noBadWordsStr[i] = "****"
			}
		}
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		w.WriteHeader(401)
		return
	}

	userId, err := auth.ValidateJWT(token, cfg.secret)
	if err != nil {
		w.WriteHeader(401)
		return
	}
	cleanedBody := strings.Join(noBadWordsStr, " ")
	params := database.CreateChirpParams{
		Body:   cleanedBody,
		UserID: userId,
	}
	record, err := cfg.db.CreateChirp(r.Context(), params)
	if err != nil {
		errResp := ErrorResponse{
			Error: "Error creating chirp",
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

	respBadWords := ValidResponse{
		ID:         record.ID,
		Created_at: record.CreatedAt,
		Updated_at: record.UpdatedAt,
		Body:       cleanedBody,
		User_id:    record.UserID,
	}
	d, err := json.Marshal(respBadWords)
	if err != nil {
		errResp := ErrorResponse{
			Error: "Error marshalling responce",
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
	w.WriteHeader(201)
	w.Write(d)
}
