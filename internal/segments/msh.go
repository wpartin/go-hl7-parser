package segments

type MSH struct {
	FieldSeparator       string
	EncodingCharacters   string
	SendingApplication   string
	SendingFacility      string
	ReceivingApplication string
	ReceivingFacility    string
	DateTimeOfMessage    string
	Security             string
	MessageType          string
	MessageControlID     string
	ProcessingID         string
	VersionID            string
	SequenceNumber       string
	ContinuationPointer  string
}

func NewMSH() *MSH {
	return &MSH{}
}

func ParseMSH(parts []string) MSH {
	return MSH{
		FieldSeparator:       GetField(parts, 0),
		EncodingCharacters:   GetField(parts, 1),
		SendingApplication:   GetField(parts, 2),
		SendingFacility:      GetField(parts, 3),
		ReceivingApplication: GetField(parts, 4),
		ReceivingFacility:    GetField(parts, 5),
		DateTimeOfMessage:    GetField(parts, 6),
		Security:             GetField(parts, 7),
		MessageType:          GetField(parts, 8),
		MessageControlID:     GetField(parts, 9),
		ProcessingID:         GetField(parts, 10),
		VersionID:            GetField(parts, 11),
		SequenceNumber:       GetField(parts, 12),
		ContinuationPointer:  GetField(parts, 13),
	}
}
