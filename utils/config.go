package utils

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kelseyhightower/envconfig"
)

// CronoscopeConfig is the application configuration
type CronoscopeConfig struct {
	Disabled            bool   `default: false`
	PollingInterval     int    `default:"10" split_words:"true"`
	PushergatewayHost   string `required:"true" split_words:"true"`
	PushergatewayPort   int    `default:"9091" split_words:"true"`
	PushRetries         int    `default:"3" split_words:"true"`
	PushRetriesInterval int    `default:"2" split_words:"true"`
	LabelJob            string `required:"true" split_words:"true"`
	LabelInstance       string `required:"true" split_words:"true"`
}

// ReadConfig reads the configuration from enviornment and validates.
// In case of an error, this function will quit the program
func ReadConfig() CronoscopeConfig {

	var config CronoscopeConfig

	isDisabled, isSet := os.LookupEnv("CRONOSCOPE_DISABLED")

	if isSet == true && strings.ToLower(isDisabled) == "true" {
		config.Disabled = true
		log.Printf("CRONOSCOPE_DISABLED was set to true. Not processing further environment variables.")
		return config
	}

	err := envconfig.Process("CRONOSCOPE", &config)

	if len(os.Args) < 2 {
		fmt.Println("usage: cronosocpe your-command-here")
		os.Exit(1)
	}

	if err != nil {
		log.Fatalf(err.Error())
	}

	retryTime := (config.PushRetriesInterval * config.PushRetries) + 3

	if config.PollingInterval < retryTime {
		fmt.Println("CRONOSCOPE_POLLING_INTERVAL is too short")
		fmt.Println("CRONOSCOPE_POLLING_INTERVAL should be at least 2 seconds more than product of CRONOSCOPE_PUSH_RETRIES x CRONOSCOPE_PUSH_RETRIES_INTERVAL")

		fmt.Println("CRONOSCOPE_POLLING_INTERVAL=", config.PollingInterval)
		fmt.Println("CRONOSCOPE_PUSH_RETRIES=", config.PushRetries)
		fmt.Println("CRONOSCOPE_PUSH_RETRIES_INTERVAL=", config.PushRetriesInterval)
		os.Exit(1)
	}
	return config

}
