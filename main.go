package main

import (
	"github.com/ekomi-ltd/cronoscope/utils"
)

func main() {
	config := utils.ReadConfig()
	utils.StartAgent(&config)
	process := utils.LaunchProcess(&config)
	process.Wait()
	utils.StopAgent(&config)
}
