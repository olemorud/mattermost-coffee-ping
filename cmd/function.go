package coffee

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/logging"
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

func contains(haystack []string, needle string) bool {
	for _, shiny := range haystack {
		if shiny == needle {
			return true
		}
	}
	return false
}

// Coffee() replies to mattermost webhooks with the correct token value
func Coffee(writer http.ResponseWriter, request *http.Request) {
	project_id := os.Getenv("PROJECT_ID")

	if project_id == "" {
		log.Fatal("Environment variable PROJECT_ID is empty")
	}

	// Create cloud logging client
	ctx := context.Background()
	client, err := logging.NewClient(ctx, project_id)

	logger := client.Logger("Coffee-log")
	var stdlog *log.Logger

	if err != nil {
		defer client.Close()
		stdlog = logger.StandardLogger(logging.Debug)
	} else {
		log.Printf("Failed to create logging client: %v", err)
		stdlog = log.New(os.Stderr, "", log.LstdFlags)
	}

	// Load valid tokens to reply to
	tokens := make([]string, 0)

	for _, envvar := range []string{"MATTERMOST_TOKEN", "TEST_TOKEN"} {
		token := os.Getenv(envvar)
		if token != "" {
			tokens = append(tokens, token)
		} else {
			stdlog.Printf("Environment variable %v is empty\n", token)
		}
	}

	// Unmarshal message and reply
	incoming := message{}
	err = json.NewDecoder(request.Body).Decode(&incoming)

	if err != nil {
		stdlog.Printf("Failed to decode message: %v\n", err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if !contains(tokens, incoming.Token) {
		stdlog.Printf("Invalid request received with token %+v", incoming.Token)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(writer).Encode(&reply{Response_type: "comment", Text: "@omorud"})

	if err != nil {
		stdlog.Printf("Failed to encode response: %v\n", err)
	}
}
