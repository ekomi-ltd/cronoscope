package utils

import (
	"fmt"
	"strings"
	"time"
	// "net/http"

	"github.com/ekomi-ltd/cronoscope/controllers"
)

var done chan bool= make(chan bool)
var ticker *time.Ticker

func sendData(builder *strings.Builder, config *CronoscopeConfig) {
	
	fmt.Print(builder.String())

	// retries := 3
	// reader := strings.NewReader(builder.String())

	// for retries > 0 {
	// 	response, err := http.Post(config.PushergatewayHost, "application/octet-stream", reader)

	// 	if err != nil {
	// 		retries--
	// 	} else {
	// 		break;
	// 	}
	// }

	// defer response.Body.Close()

}


func startMonitoringAgent(config * CronoscopeConfig) {

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

func StartAgent(config * CronoscopeConfig) {
	ticker = time.NewTicker(time.Duration(config.PollingInterval) * time.Second)
	go startMonitoringAgent(config)
}

func StopAgent(){
	ticker.Stop()
	done <- true
	fmt.Println("Timer stopped")
}
