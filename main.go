package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
)

const (
	HAUrl 	= "https://homeassistant.example.com"
	HAToken = "LONG_LIVED_TOKEN"
)

func HandleRequest(payload map[string]json.RawMessage) (map[string]json.RawMessage, error) {

	jsonPayload, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", HAUrl + "/api/alexa/smart_home", bytes.NewBuffer(jsonPayload))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer " + HAToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var jsonResponse map[string]json.RawMessage
	err = json.Unmarshal(body, &jsonResponse)

	if err != nil {
		return nil, err
	}

	return jsonResponse, nil
}

func main(){
	lambda.Start(HandleRequest)
}
