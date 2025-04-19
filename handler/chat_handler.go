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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"response": err.Error()})
		return
	}

	json.NewEncoder(w).Encode(model.ChatResponse{Response: reply})
}
