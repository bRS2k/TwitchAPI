package twitchapi

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/parnurzeal/gorequest"
)

// the number of seconds that will be waited if a 503 response is gotten
var timeout = 2

// the number of times to retry a request before giving up
var numTries = 3

// This is the template request format for all requests to be made
var request = gorequest.New().Header.Add("Client-ID", readAPIConfig().clientID)

// APIConfig is the configuration information necessary to connect to the twitch API
type apiConfig struct {
	clientID string
}

// ReadAPIConfig reads the config information from api.cfg and returns a config struct
// with the information
func readAPIConfig() APIConfig {
	file, _ := os.Open("api.cfg")
	decoder := json.NewDecoder(file)
	apiConfig := apiConfig{}
	if err := decoder.Decode(&apiConfig); err != nil {
		fmt.Println("Error while decoding config file:\n", err)
	}
}
