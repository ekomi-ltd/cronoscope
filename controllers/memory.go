package controllers

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"
)

type MemoryController struct {
	path    string
	metrics map[string]string
}

/// NewMemoryMetrics
/// creates a new memory metrics controller
func NewMemoryController() *MemoryController {
	memoryRoot := "/sys/fs/cgroup/memory"
	m := MemoryController{
		path: memoryRoot,
		metrics: map[string]string{
			"memory_limit_in_bytes":      path.Join(memoryRoot, "memory.limit_in_bytes"),
			"memory_max_usage_in_bytes":  path.Join(memoryRoot, "memory.max_usage_in_bytes"),
			"memory_soft_limit_in_bytes": path.Join(memoryRoot, "memory.soft_limit_in_bytes"),
			"memory_usage_in_bytes":      path.Join(memoryRoot, "memory.usage_in_bytes"),
		},
	}
	return &m
}

func (m *MemoryController) Read(b *strings.Builder) {

	for metric, file := range m.metrics {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println("Faild to read " + file)

		}
		b.WriteString(fmt.Sprintf("%s=%s", metric, data))
	}
}
