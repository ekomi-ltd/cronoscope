package main

import (
	"fmt"
	"github.com/ekomi-ltd/cronoscope/controllers"
	"os"
	"os/exec"
	"strings"
	"time"
)

func Start(args ...string) (p *os.Process, err error) {
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

	// 1. Send at the very start
	readAndSendMetrics()

	for {
		select {
		case <-done:
			readAndSendMetrics() // 3. Send when quiting
			return
		case <-ticker.C:
			readAndSendMetrics() // 2. And of course, at configured internvals
		}

	}
}

func main() {

	done := make(chan bool)
	ticker := time.NewTicker(5 * time.Second)

	go StartMonitoringAgent(ticker, done)

	process, err := Start(os.Args[1:]...)

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	fmt.Println("Going to wait for the process.")
	process.Wait()
	fmt.Print("Process stopped")
	fmt.Print("Going to stop the timer")
	ticker.Stop()
	fmt.Print("Timer stopped")

}
