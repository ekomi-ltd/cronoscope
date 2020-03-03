package controllers

import (
	"fmt"
	"strings"
)

func writeMetric(buffer *strings.Builder, metricsPrefix string, metric string, value string, labels string, type_ string, help string) {
	buffer.WriteString(fmt.Sprintf("# HELP %v%v %v\n", metricsPrefix, metric, help))
	buffer.WriteString(fmt.Sprintf("# TYPE %v%v %v\n", metricsPrefix, metric, type_))
	buffer.WriteString(fmt.Sprintf("%v%v", metricsPrefix, metric))
	buffer.WriteString(labels)
	buffer.WriteString(fmt.Sprintf(" %s\n", value))
}
