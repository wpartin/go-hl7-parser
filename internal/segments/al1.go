package segments

type AL1 struct {
	SetID               string
	AllergenTypeCode    string
	AllergenCode        string
	AllergySeverityCode string
	AllergyReaction     string
	IdentificationDate  string
}

func (a *AL1) NewAL1() *AL1 {
	return &AL1{}
}

func ParseAL1(fields []string) AL1 {
	if len(fields) > 0 && fields[0] == "AL1" {
		fields = fields[1:] // shift to remove segment type
	}

	return AL1{
		SetID:               GetField(fields, 0),
		AllergenTypeCode:    GetField(fields, 1),
		AllergenCode:        GetField(fields, 2),
		AllergySeverityCode: GetField(fields, 3),
		AllergyReaction:     GetField(fields, 4),
		IdentificationDate:  GetField(fields, 5),
	}
}
