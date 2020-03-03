package utils

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ekomi-ltd/cronoscope/controllers"
)

// done channel used to indicate to the go routine that we are done
var done chan bool = make(chan bool)

// ticker is the heartbeat of the monitoring agent.
var ticker *time.Ticker

// sendData sends the data from thebuilder to the pushergateway endpoint
// as mentioned in the configuration. Retries are also read from the
// configuration
func sendData(builder *strings.Builder, config *CronoscopeConfig) {

	fmt.Print(builder.String())
	var response *http.Response = nil
	var err error
	retries := config.PushRetries
	forSometime := time.Duration(config.PushRetriesInterval) * time.Second
	reader := strings.NewReader(builder.String())
	host := net.JoinHostPort(config.PushergatewayHost, strconv.Itoa(config.PushergatewayPort))
	endpoint := fmt.Sprintf("http://%s/metrics/job/%s", host, config.LabelJob)
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
			log.Printf("%d/%d Retry Faild to send data, Sleeping for %d seconds",
				(config.PushRetries - retries), config.PushRetries, config.PushRetriesInterval)
			time.Sleep(forSometime)
		} else {
			failed = false
			retries = 0
			response.Body.Close()
		}
	}

	if failed == true {
		log.Printf("All %d retries with %d second intervals failed\n", config.PushRetries, config.PushRetriesInterval)
	}

	defer closeResponse()

}

// startMonitoringAgent stars the infinit loop of reading and sending the
// data over the wire on fixed intervals.
func startMonitoringAgent(config *CronoscopeConfig) {

	var builder strings.Builder
	memoryController := controllers.NewMemoryController(config.Labels)
	cpuacctController := controllers.NewCPUAcctController(config.Labels)

	readAndSendMetrics := func() {
		builder.Reset()
		cpuacctController.Read(&builder)
		memoryController.Read(&builder)
		sendData(&builder, config)
	}

	// 1. Send at the start and end
	readAndSendMetrics()
	defer readAndSendMetrics()

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			readAndSendMetrics() // 2. And of course, at configured internvals
		}

	}
}

// StartAgent initilizes the tickeer and starts the monitoring as a go routine.
func StartAgent(config *CronoscopeConfig) {

	if config.Disabled == true {
		log.Println("CRONOSCOPE_DISABLED - Monitoring agent will not be started.")
		return
	}

	ticker = time.NewTicker(time.Duration(config.PollingInterval) * time.Second)
	go startMonitoringAgent(config)
}

// StopAgent stops the earlier started agent by stopping the ticker and sending
// the done signal on the channel
func StopAgent(config *CronoscopeConfig) {

	if config.Disabled == true {
		return
	}

	ticker.Stop()
	done <- true
}
