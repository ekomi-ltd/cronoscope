package main

import (
	"fmt"
	"github.com/ekomi-ltd/cronoscope/controllers"
	"os"
	"os/exec"
	"strings"
	"time"
)

func StartProcess(args ...string) (p *os.Process, err error) {
	if args[0], err = exec.LookPath(args[0]); err == nil {
		var procAttr os.ProcAttr
		procAttr.Files = []*os.File{os.Stdin,
			os.Stdout, os.Stderr}
		p, err := os.StartProcess(args[0], args, &procAttr)
		if err == nil {
			return p, nil
		}
	}
	return nil, err
}

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

func main() {

	//	CRONOSCOPE_POLLING_INTERVAL
	// 	CRONOSCOPE_PUSHERGATEWAY_HOST
	//	CRONOSCOPE_PUSHERGATEWAY_PORT
	//  CRONOSCOPE_ADDITIONAL_LABELS

	done := make(chan bool)
	ticker := time.NewTicker(5 * time.Second)

	go StartMonitoringAgent(ticker, done)

	process, err := StartProcess(os.Args[1:]...)

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	fmt.Println("Going to wait for the process.")
	process.Wait()
	fmt.Println("Process stopped")
	fmt.Println("Going to stop the timer")
	ticker.Stop()
	done <- true
	fmt.Println("Timer stopped")

}
