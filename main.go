package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ekomi-ltd/cronoscope/utils"
	"github.com/kelseyhightower/envconfig"
)

func main() {

	var config utils.CronoscopeConfig

	err := envconfig.Process("CRONOSCOPE", &config)

	fmt.Println("Parsed environment variables.")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("%+v\n", config)

	done := make(chan bool)
	ticker := time.NewTicker(5 * time.Second)

	go utils.StartMonitoringAgent(ticker, done)

	process, err := utils.StartProcess(os.Args[1:]...)

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
