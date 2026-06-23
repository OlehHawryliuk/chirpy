package main

import (
	"chirpy/internal/auth"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handerWebhooks(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		w.WriteHeader(401)
		return
	}
	if apiKey != cfg.polkaKey {
		w.WriteHeader(401)
		return
	}
	webhookData := WebHooksRequest{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&webhookData)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	if webhookData.Event != "user.upgraded" {
		w.WriteHeader(204)
		return
	}

	userId, err := uuid.Parse(webhookData.Data.UserId)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	err = cfg.db.UpgradeUserToRed(r.Context(), userId)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(204)
}
