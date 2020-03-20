package utils

import (
	"log"
	"os"
	"os/exec"
	"syscall"
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

func replaceProcess(args ...string) {

	binary, lookErr := exec.LookPath(args[0])
	if lookErr != nil {
		panic(lookErr)
	}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}

}

// LaunchProcess will launch the process given on the command line and would
// return the os.Process pointer
func LaunchProcess(config *CronoscopeConfig) (p *os.Process) {

	if config.Disabled == true {
		log.Println("CRONOSCOPE_DISABLED - Current process will be replaced.")
		replaceProcess(os.Args[1:]...) // No code executes beyond this point anyway.
	}

	process, err := startProcess(os.Args[1:]...)

	if err != nil {
		log.Panicf(err.Error())
	}

	return process
}
