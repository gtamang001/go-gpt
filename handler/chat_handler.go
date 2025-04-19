package handler

import (
	"encoding/json"
	"go-gpt/model"
	"go-gpt/service"
	"net/http"
)

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	var req model.ChatRequest
	json.NewDecoder(r.Body).Decode(&req)

	reply, err := service.GetChatGPTResponse(req.Message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(model.ChatResponse{Response: reply})
}
