package utils

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
)

type CronoscopeConfig struct {
	PollingInterval     int    `default:"10" split_words:"true"`
	PushergatewayHost   string `required:"true" split_words:"true"`
	PushergatewayPort   int    `default:"9091" split_words:"true"`
	PushRetries         int    `default:"3" split_words:"true"`
	PushRetriesInterval int    `default:"2" split_words:"true"`
}

func ReadConfig() CronoscopeConfig {
	var config CronoscopeConfig
	err := envconfig.Process("CRONOSCOPE", &config)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
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
