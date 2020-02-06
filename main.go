package main

import (
	"github.com/ekomi-ltd/cronoscope/utils"
)

func main() {
	config := utils.ReadConfig()
	utils.StartAgent(&config)
	process := utils.LaunchProcess()
	process.Wait()
	utils.StopAgent()
}
