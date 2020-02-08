package utils

import (
	"log"
	"os"
	"os/exec"
)

func startProcess(args ...string) (p *os.Process, err error) {
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

// LaunchProcess will launch the process given on the command line and would
// return the os.Process pointer
func LaunchProcess() (p *os.Process) {
	process, err := startProcess(os.Args[1:]...)

	if err != nil {
		log.Panicf(err.Error())
	}

	return process
}
