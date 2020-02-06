package utils

import (
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ekomi-ltd/cronoscope/controllers"
)

var done chan bool = make(chan bool)
var ticker *time.Ticker

func sendData(builder *strings.Builder, config *CronoscopeConfig) {

	// fmt.Print(builder.String())

	var response *http.Response = nil
	var err error
	retries := config.PushRetries
	forSometime := time.Duration(config.PushRetriesInterval) * time.Second
	reader := strings.NewReader(builder.String())
	endpoint := net.JoinHostPort(config.PushergatewayHost, strconv.Itoa(config.PushergatewayPort))
	failed := true

	closeResponse := func() {
		if response != nil {
			response.Body.Close()
		}
	}

	for retries > 0 {
		response, err = http.Post(endpoint, "application/octet-stream", reader)

		if err != nil {
			retries--
			response = nil
			failed = true
			fmt.Printf("%d/%d Retry Faild to send data, Sleeping for %d seconds",
				(config.PushRetries - retries), config.PushRetries, config.PushRetriesInterval)
			time.Sleep(forSometime)
		} else {
			failed = true
			break
		}
	}

	if failed == true {
		fmt.Printf("All %d retries with %d second intervals failed\n", config.PushRetries, config.PushRetriesInterval)
	}

	defer closeResponse()

}

func startMonitoringAgent(config *CronoscopeConfig) {

	var builder strings.Builder
	memoryController := controllers.NewMemoryController()

	readAndSendMetrics := func() {
		builder.Reset()
		memoryController.Read(&builder)
		sendData(&builder, config)
	}

	// 1. Send at the start and end
	readAndSendMetrics()
	defer readAndSendMetrics()

	for {
		select {
		case <-done:
			fmt.Println("Timer stop return statement!")
			return
		case <-ticker.C:
			readAndSendMetrics() // 2. And of course, at configured internvals
		}

	}
}

func StartAgent(config *CronoscopeConfig) {
	ticker = time.NewTicker(time.Duration(config.PollingInterval) * time.Second)
	go startMonitoringAgent(config)
}

func StopAgent() {
	ticker.Stop()
	done <- true
	fmt.Println("Timer stopped")
}
