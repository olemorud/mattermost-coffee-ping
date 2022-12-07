package coffee

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type reply struct {
	Response_type string `json:"response_type"`
	Text          string `json:"text"`
}

type message struct {
	Message      string `json:"message"`
	Channel_id   string `json:"channel_id"`
	Channel_name string `json:"channel_name"`
	Team_domain  string `json:"team_domain"`
	Team_id      string `json:"team_id"`
	Post_id      string `json:"post_id"`
	Text         string `json:"text"`
	Timestamp    int64  `json:"timestamp"`
	Token        string `json:"token"`
	Trigger_word string `json:"trigger_word"`
	User_id      string `json:"user_id"`
	User_name    string `json:"user_name"`
}

// Coffee() replies to mattermost webhooks with the correct token value
func Coffee(writer http.ResponseWriter, request *http.Request) {
	mattermost_token := os.Getenv("MATTERMOST_TOKEN")
	test_token := os.Getenv("TEST_TOKEN")

	if mattermost_token == "" || test_token == "" {
		log.Fatal("Failed to load mattermost token")
	}

	incoming_msg := message{}
	err := json.NewDecoder(request.Body).Decode(&incoming_msg)

	if err != nil {
		log.Printf("json.NewDecoder: %v\n", err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	log.Printf("Coffee requested by user id: %+v", incoming_msg.User_id)
	log.Println(incoming_msg)

	if incoming_msg.Token == mattermost_token || incoming_msg.Token == test_token {
		json.NewEncoder(writer).Encode(&reply{Response_type: "comment", Text: "@omorud"})
	} else {
		log.Printf("Invalid request received with token: %v", incoming_msg.Token)
	}
}
