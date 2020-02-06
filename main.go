package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	//"time"
	//"controllers"
	"github.com/ekomi-ltd/cronoscope/controllers"

	"os"
)

// type MemoryMetrics struct {
// 	path string
// 	metrics map[string]string
// }

// /// NewMemoryMetrics
// /// creates a new memory metrics controller
// func NewMemoryMetrics() * MemoryMetrics {
// 	path := "/sys/fs/cgroup/memory"
// 	m := MemoryMetrics{
// 		path: path,
// 		metrics: map[string]string {
// 			"memory_limit_in_bytes": path + "/" + "memory.limit_in_bytes",
// 			"memory_max_usage_in_bytes": path + "/" + "memory.max_usage_in_bytes",
// 			"memory_soft_limit_in_bytes": path + "/" + "memory.soft_limit_in_bytes",
// 			"memory_usage_in_bytes": path + "/" + "memory.usage_in_bytes",
// 		},
// 	}
// 	return &m
// }

// func (m * MemoryMetrics) read(b *strings.Builder) {

// 	for metric, file := range m.metrics {
// 		data, err := ioutil.ReadFile(file)
// 		if err != nil {
// 			fmt.Println("Faild to read " + file)

// 		}
// 		b.WriteString(fmt.Sprintf("%s=%s", metric, data))
// 	}
// }

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

func pollAndSendMetrics(quit <-chan bool, done chan<- bool) {
	var builder strings.Builder
	memoryController := controllers.NewMemoryController()

	i := 0

	for {
		select {
		case done := <- quit:
			if done {
				fmt.Println("--- Routine will stop! ---")
				return;
			}
		default:
			builder.Reset()
			memoryController.Read(&builder)
			fmt.Print(builder.String())
			time.Sleep(1 * time.Second)
			i = i + 1
			if i == 5 {
				done <- true
				return;
			}
		}
		
	}
}

func main() {

	// Start the go routine
	// Start the process
	// Wait for process
	// Wait for routine
	// Add logging

	process, err := Start(os.Args[1:]...)

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	process.Wait()

	quit := make(chan bool)
	done := make(chan bool)

	fmt.Println("About to start polling metrics")
	go pollAndSendMetrics(quit, done)
	// fmt.Println("Routine started")
	// fmt.Println("Going to sleep for 10 seconds")
	// time.Sleep(10 * time.Second)
	// fmt.Println("Waking up and going to stop routine")
	// quit <- true
	// fmt.Println("Signal sent to stop routine")
	// fmt.Println("Will sleep for 5 seconds")
	// time.Sleep(5 * time.Second)
	// fmt.Println("Done and goodbye!")
	what := <- done
	fmt.Println("Polling done", what)



	// Start go routine

	// for {
	// 	builder.Reset()
	// 	memorMetrics.read(&builder)
	// 	fmt.Println(builder.String())
	// 	time.Sleep(2 * time.Second)
	// }

	// args := []string{"main.py"}
	// var procAttr os.ProcAttr
	// path, err := exec.LookPath("python")
	// if err != nil {
	// 	fmt.Println(path)
	// }
	// fmt.Println(path)

	// procAttr.Files = []*os.File{os.Stdin,
	// 	os.Stdout, os.Stderr} 

	// process, err := os.StartProcess(path, args, &procAttr)

	// if err != nil {
	// 	fmt.Print(err)
	// 	os.Exit(1)
	// }

	// process.Wait()
	// fmt.Println("Done waiting.")

	// if proc, err := Start("python", "/tmp/main.py"); err == nil {
	// 	proc.Wait()
	// }

}
