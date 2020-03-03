package controllers

import (
	"fmt"
	"strings"
)

func writeMetric(buffer *strings.Builder, metric string, value string, labels string, type_ string, help string) {
	buffer.WriteString(fmt.Sprintf("# HELP %v %v\n", metric, help))
	buffer.WriteString(fmt.Sprintf("# TYPE %v %v\n", metric, type_))
	buffer.WriteString(metric)
	buffer.WriteString(labels)
	buffer.WriteString(fmt.Sprintf(" %s\n", value))
}
