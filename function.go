
import (
	"encoding/json"
	"log"
	"net/http"
)

type reply struct {
	Response_type string `json:"response_type"`
	Text          string `json:"text"`
}

// HelloWorld prints the JSON encoded "message" field in the body
// of the request or "Hello, World!" if there isn't one.
func Coffee(writer http.ResponseWriter, request *http.Request) {
	var data struct {
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

	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		log.Printf("json.NewDecoder: %v\n", err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	log.Printf("%+v", data.User_id)
	if data.Token == "moha5q8xqjg9prbq7q73676z9y" || data.Token == "33szr8spqi8tx81ymdsxpq5mjh" {
		json.NewEncoder(writer).Encode(&reply{Response_type: "comment", Text: "@omorud @ksarband @psvihra"})
	}
}

