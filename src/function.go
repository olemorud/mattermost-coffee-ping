package coffee

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"cloud.google.com/go/logging"
	"example.com/cloudfunction/egroups"
	"github.com/hooklift/gowsdl/soap"
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

func getCoffeeDrinkers(logger *log.Logger) []string {
	username := os.Getenv("EGROUPS_USERNAME")
	password := os.Getenv("EGROUPS_PASSWORD")

	soapClient := soap.NewClient("https://foundservices.cern.ch:443/ws/egroups/v1/EgroupsWebService/", soap.WithBasicAuth(username, password))
	egroupsClient := egroups.NewEgroupsService(soapClient)

	request := egroups.FindEgroupByIdRequest{Id: 10542497}
	reply, err := egroupsClient.FindEgroupById(&request)

	if err != nil {
		logger.Println("Could not get egroup members:", err)
		return []string{"@omorud"}
	}

	drinkers := make([]string, 0)

	for _, member := range reply.Result.Members {
		drinkers = append(drinkers, "@"+strings.ToLower(member.PrimaryAccount))
	}

	return drinkers
}

// Coffee replies to mattermost webhooks containing valid token values
func Coffee(writer http.ResponseWriter, request *http.Request) {
	// Create cloud logging client or default to stderr
	project_id := os.Getenv("PROJECT_ID")
	ctx := context.Background()
	client, err := logging.NewClient(ctx, project_id)
	var stdlog *log.Logger
	var warning *log.Logger

	if err != nil {
		log.Printf("Failed to create logging client: %v", err)
		stdlog = log.New(os.Stderr, "", log.LstdFlags)
	} else {
		defer client.Close()
		logger := client.Logger("Coffee-log")
		stdlog = logger.StandardLogger(logging.Debug)
		warning = logger.StandardLogger(logging.Warning)
	}

	// Load valid mattermost tokens to reply to from env
	tokens := make([]string, 0)

	for _, envvar := range []string{"MATTERMOST_TOKEN", "TEST_TOKEN"} {
		token := os.Getenv(envvar)
		if token != "" {
			tokens = append(tokens, token)
		} else {
			stdlog.Printf("Environment variable %v is not found/empty\n", token)
		}
	}

	// Unmarshal incoming message and reply
	incoming := message{}
	err = json.NewDecoder(request.Body).Decode(&incoming)

	if err != nil {
		warning.Printf("Failed to decode message: %v\n", err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if !contains(tokens, incoming.Token) {
		warning.Printf("Invalid request received with token %+v\n", incoming.Token)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	text := ""

	for _, name := range getCoffeeDrinkers(stdlog) {
		text += name + " "
	}

	url, ok := os.LookupEnv("GITHUB_URL")
	if ok {
		text += " (now open source: " + url + ")"
	}

	err = json.NewEncoder(writer).Encode(&reply{Response_type: "comment", Text: text})

	if err != nil {
		warning.Printf("Failed to encode response: %v\n", err)
	}

	stdlog.Printf("Successfully responded to message: %+v\n", incoming)
}
