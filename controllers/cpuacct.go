package controllers

import (
	"io/ioutil"
	"log"
	"strings"
)

// CPUAcctController represents a CPU accounting controller
type CPUAcctController struct {
	path       string
	metricType string
	labels     string
}

// NewCPUAcctController initialises a new controller to be used.
func NewCPUAcctController(labels string) *CPUAcctController {
	cc := CPUAcctController{
		labels:     labels,
		path:       "/sys/fs/cgroup/cpuacct/cpuacct.stat",
		metricType: "gauge",
	}

	return &cc
}

func (c *CPUAcctController) Read(b *strings.Builder) {
	data, err := ioutil.ReadFile(c.path)
	if err != nil {
		log.Printf("CPUAcctController - Failed to read %s", c.path)
	}
	stats := string(data)
	lines := strings.Split(stats, "\n")

	if len(lines) < 2 {
		log.Printf("CPUAcctController - Expected to contain 2 lines.")
		return
	}
	userLine := strings.Split(lines[0], " ")
	systemLine := strings.Split(lines[1], " ")

	if len(userLine) < 2 {
		log.Printf("CPUAcctController - Expected to contain user usage separated by space.")
		return
	}

	if len(systemLine) < 2 {
		log.Printf("CPUAcctController - Expected to contain system usage separated by space.")
		return
	}

	writeMetric(b, "cpuacct_stat_user", userLine[1], c.labels, c.metricType, "CPU time spent in user mode")
	writeMetric(b, "cpuacct_stat_system", systemLine[1], c.labels, c.metricType, "CPU time spent in kernel mode")
}
