package controllers

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"
)

type MemoryController struct {
	path    string
	metrics map[string]map[string]string
}

/// NewMemoryMetrics
/// creates a new memory metrics controller
func NewMemoryController() *MemoryController {
	memoryRoot := "/sys/fs/cgroup/memory"
	m := MemoryController{
		path: memoryRoot,
		metrics: map[string]map[string]string{
			"memory_limit_in_bytes": map[string]string{
				"help": "Limit of memory usage",
				"type": "gauge",
				"path": path.Join(memoryRoot, "memory.limit_in_bytes"),
			},
			"memory_max_usage_in_bytes": map[string]string{
				"help": "Max memory usage recorded",
				"type": "gauge",
				"path": path.Join(memoryRoot, "memory.max_usage_in_bytes"),
			},

			"memory_soft_limit_in_bytes": map[string]string{
				"help": "Soft limit of memory usage",
				"type": "gauge",
				"path": path.Join(memoryRoot, "memory.soft_limit_in_bytes"),
			},

			"memory_usage_in_bytes": {
				"help": "Current usage for memory",
				"type": "gauge",
				"path": path.Join(memoryRoot, "memory.usage_in_bytes"),
			},
		},
	}
	return &m
}

func (m *MemoryController) Read(b *strings.Builder) {

	for metric, config := range m.metrics {
		data, err := ioutil.ReadFile(config["path"])
		if err != nil {
			fmt.Println("Faild to read " + config["path"])
		}
		b.WriteString(fmt.Sprintf("# TYPE %v %v\n", metric, config["type"]))
		b.WriteString(fmt.Sprintf("# HELP %v %v\n", metric, config["help"]))
		b.WriteString(fmt.Sprintf("%v %v\n", metric, string(data)))
	}
}
