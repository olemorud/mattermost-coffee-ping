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

type mattermostReply struct {
	Response_type string `json:"response_type"`
	Text          string `json:"text"`
}

type mattermostMessage struct {
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

func omitEmptyStrings(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

var debug *log.Logger
var warning *log.Logger

func getCoffeeDrinkers() []string {
	username := os.Getenv("EGROUPS_USERNAME")
	password := os.Getenv("EGROUPS_PASSWORD")
	egroups_url := "https://foundservices.cern.ch:443/ws/egroups/v1/EgroupsWebService/"

	soapClient := soap.NewClient(egroups_url, soap.WithBasicAuth(username, password))
	egroupsClient := egroups.NewEgroupsService(soapClient)

	request := egroups.FindEgroupByIdRequest{Id: 10542497}
	reply, err := egroupsClient.FindEgroupById(&request)

	if err != nil {
		warning.Println("Could not get egroup members:", err)
		return []string{"@omorud"}
	}

	drinkers := make([]string, 0)

	for _, member := range reply.Result.Members {
		drinkers = append(drinkers, "@"+strings.ToLower(member.PrimaryAccount))
	}

	return drinkers
}

// Entrypoint of google cloud function:
//
//	Upon receiving a webhook from mattermost, respond to the webhook by tagging
//	everyone in e-group 10542497
func Coffee(writer http.ResponseWriter, request *http.Request) {
	// Create cloud logging client, default to stderr on failure
	project_id := os.Getenv("PROJECT_ID")
	client, err := logging.NewClient(context.Background(), project_id)

	if err != nil {
		log.Printf("Failed to create logging client: %v", err)
		debug = log.New(os.Stderr, "", log.LstdFlags)
		warning = log.New(os.Stderr, "", log.LstdFlags)
	} else {
		defer client.Close()
		logger := client.Logger("Coffee-log")
		debug = logger.StandardLogger(logging.Debug)
		warning = logger.StandardLogger(logging.Warning)
	}

	// Unmarshal incoming message
	incomingMsg := mattermostMessage{}
	err = json.NewDecoder(request.Body).Decode(&incomingMsg)

	if err != nil {
		warning.Printf("Failed to decode message: %v\n%+v\n", err, request.Body)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	unsafeWhitelist := strings.Split(os.Getenv("COLON_SEPARATED_TOKEN_WHITELIST"), ":")
	whitelist := omitEmptyStrings(unsafeWhitelist)

	if !contains(whitelist, incomingMsg.Token) {
		warning.Printf("Invalid request received with token %+v\n", incomingMsg.Token)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Respond to incoming message
	reply := mattermostReply{
		Response_type: "comment",
		Text:          "",
	}

	for _, name := range getCoffeeDrinkers() {
		reply.Text += name + " "
	}

	reply.Text += "| [get tagged by bot](https://e-groups.cern.ch/e-groups/EgroupsSubscription.do?egroupId=10542497)"

	url, ok := os.LookupEnv("GITHUB_URL")
	if ok {
		reply.Text += " | [view source](" + url + ")"
	}

	err = json.NewEncoder(writer).Encode(reply)

	if err != nil {
		warning.Printf("Failed to encode response: %v\n", err)
	}
	debug.Printf("Successfully responded to message: %+v\n", incomingMsg)
}
