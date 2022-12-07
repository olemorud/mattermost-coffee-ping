#!/usr/bin/env sh

data='{"Message":"coffee", "Channel_id":"abc", "Channel_name":"Town Square", "Team_domain":"abc", "Team_id":"abc", "Post_id":"post id", "Text":"coffee", "Timestamp":"0", "Token":'$TEST_TOKEN'}'

curl -m 70 -X POST https://europe-west1-rising-city-366608.cloudfunctions.net/coffee \
-H "Authorization: bearer $(gcloud auth print-identity-token)" \
-H "Content-Type: application/json" \
-d "$data"