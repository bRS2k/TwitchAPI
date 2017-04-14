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
var request = gorequest.New()
var clientID = readAPIConfig().ClientID

// APIConfig is the configuration information necessary to connect to the twitch API
type apiConfig struct {
	ClientID string
}

// ReadAPIConfig reads the config information from api.cfg and returns a config struct
// with the information
func readAPIConfig() apiConfig {
	configFile, _ := os.Open("../config/apiconf.json")
	decoder := json.NewDecoder(configFile)
	apiConfig := apiConfig{}
	err := decoder.Decode(&apiConfig)
	if err != nil {
		fmt.Println("Error while decoding config file:\n", err)
	}
	fmt.Println(apiConfig)
	return apiConfig
}
