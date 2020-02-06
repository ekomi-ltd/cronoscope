package utils

type CronoscopeConfig struct {
	PollingInterval   int    `default:"10" split_words:"true"`
	PushergatewayHost string `required:"true" split_words:"true"`
	PushergatewayPort int    `default:"9091" split_words:"true"`
}
