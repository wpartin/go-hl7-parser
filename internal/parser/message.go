package parser

import (
	"strings"

	"github.org/wpartin/go/hl7-parser/internal/segments"
)

type Message struct {
	AL1 []segments.AL1
	MSH segments.MSH
	PID segments.PID
	PV1 segments.PV1
}

func ParseMessage(raw string) (*Message, error) {
	lines := strings.Split(raw, "\r")
	var msg Message

	for _, line := range lines {
		parts := TokenizeSegment(line)
		switch parts[0] {
		case "AL1":
			msg.AL1 = append(msg.AL1, segments.ParseAL1(parts))
		case "MSH":
			msg.MSH = segments.ParseMSH(parts)
			msg.MSH.FieldSeparator = strings.Split(line, "MSH")[1][:1]
		case "PID":
			msg.PID = segments.ParsePID(parts)
		case "PV1":
			msg.PV1 = segments.ParsePV1(parts)
		}
	}

	return &msg, nil
}