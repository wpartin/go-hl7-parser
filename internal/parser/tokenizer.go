package parser

import (
	"strings"
)

const (
	componentSeparator = '^'
	fieldSeparator     = '|'
)

func TokenizeSegment(raw string) []string {
	// Split the segment by the field separator
	fields := strings.Split(raw, string(fieldSeparator))

	// Further split each field by the component separator
	for i, field := range fields {
		fields[i] = strings.TrimSpace(field)
	}

	return fields
}