package utils

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
)

type CronoscopeConfig struct {
	PollingInterval   int    `default:"10" split_words:"true"`
	PushergatewayHost string `required:"true" split_words:"true"`
	PushergatewayPort int    `default:"9091" split_words:"true"`
	PushRetries int	`default:"3" split_words:"true"`
}

func ReadConfig() (CronoscopeConfig){
	var config CronoscopeConfig
	err := envconfig.Process("CRONOSCOPE", &config)

	fmt.Println("Parsed environment variables.")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return config

}