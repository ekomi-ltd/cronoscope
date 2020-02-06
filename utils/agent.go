package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/ekomi-ltd/cronoscope/controllers"
)

func StartMonitoringAgent(ticker *time.Ticker, done <-chan bool) {

	var builder strings.Builder
	memoryController := controllers.NewMemoryController()

	readAndSendMetrics := func() {
		builder.Reset()
		memoryController.Read(&builder)
		fmt.Print(builder.String())
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
