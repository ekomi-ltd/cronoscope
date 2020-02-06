package main

import (
	"fmt"
	"os"
	"github.com/ekomi-ltd/cronoscope/utils"
)

func main() {


	config := utils.ReadConfig()

	utils.StartAgent(&config)
	
	process, err := utils.StartProcess(os.Args[1:]...)

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	fmt.Println("Going to wait for the process.")
	process.Wait()
	utils.StopAgent()
}
