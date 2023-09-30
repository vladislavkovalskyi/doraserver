package handlers

import (
	"encoding/json"
	"net/http"
	"sync"
)

type Message struct {
	Nickname string `json:"nickname"`
	Content  string `json:"content"`
}

var (
	messages []Message
	mutex    sync.Mutex
)

func HandleSendMessage(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var message Message

	err := decoder.Decode(&message)
	if err != nil {
		http.Error(w, "Request reading error", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	messages = append(messages, message)
	mutex.Unlock()

	w.WriteHeader(http.StatusCreated)
}

func HandleGetMessages(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	response, err := json.Marshal(messages)
	if err != nil {
		http.Error(w, "Error during serialization of messages in JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
