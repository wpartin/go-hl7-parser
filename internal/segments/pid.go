package segments

type PID struct {
	SetID                string
	PatientIDExternal    string
	PatientIDInternal    string
	AlternatePatientID   string
	PatientName          string
	MotherMaidenName     string
	DateOfBirth          string
	Sex                  string
	PatientAlias         string
	EthnicGroup          string
	PatientAddress       string
	CountyCode           string
	PhoneNumberHome      string
	PhoneNumberBusiness  string
	LanguagePatient      string
	MaritalStatus        string
	Religion             string
	PatientAccountNumber string
	SocialSecurityNumber string
	DriverLicenseNumber  string
}

func (p *PID) NewPID() *PID {
	return &PID{}
}

func ParsePID(parts []string) PID {
	parts = parts[1:] // Skip the segment identifier (PID)

	return PID{
		SetID:                GetField(parts, 0),
		PatientIDExternal:    GetField(parts, 1),
		PatientIDInternal:    GetField(parts, 2),
		AlternatePatientID:   GetField(parts, 3),
		PatientName:          GetField(parts, 4),
		MotherMaidenName:     GetField(parts, 5),
		DateOfBirth:          GetField(parts, 6),
		Sex:                  GetField(parts, 7),
		PatientAlias:         GetField(parts, 8),
		EthnicGroup:          GetField(parts, 9),
		PatientAddress:       GetField(parts, 10),
		CountyCode:           GetField(parts, 11),
		PhoneNumberHome:      GetField(parts, 12),
		PhoneNumberBusiness:  GetField(parts, 13),
		LanguagePatient:      GetField(parts, 14),
		MaritalStatus:        GetField(parts, 15),
		Religion:             GetField(parts, 16),
		PatientAccountNumber: GetField(parts, 17),
		SocialSecurityNumber: GetField(parts, 18),
		DriverLicenseNumber:  GetField(parts, 19),
	}
}
