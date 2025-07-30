package segments

type PV1 struct {
	SetID                   string
	PatientClass            string
	AssignedPatientLocation string
	AdmissionType           string
	PreAdmitNumber          string
	PriorPatientLocation    string
	AttendingDoctor         string
	ReferringDoctor         string
	ConsultingDoctor        string
	HospitalService         string
	TemporaryLocation       string
	PreAdmitTestIndicator   string
	AdmitSource             string
	AmbulatoryStatus        string
	VipIndicator            string
	AdmittingDoctor         string
	PatientType             string
	VisitNumber             string
	FinancialClass          string
	ChargePriceIndicator    string
	CourtesyCode            string
	CreditRating            string
	ContractCode            string
	ContractEffectiveDate   string
	ContractAmount          string
	ContractPeriod          string
	InterestCode            string
	TransferToBadDebtCode   string
	BadDebtAgencyCode       string
	BadDebtTransferAmount   string
	BadDebtRecoveryAmount   string
	DeleteAccountIndicator  string
	DeleteAccountDate       string
	DischargeDisposition    string
	DischargedToLocation    string
	DietType                string
	ServicingFacility       string
	BedStatus               string
	AccountStatus           string
	PendingLocation         string
	PriorTemporaryLocation  string
	AdmitDateTime           string
	DischargeDateTime       string
	CurrentPatientBalance   string
	TotalCharges            string
	TotalAdjustments        string
	TotalPayments           string
}

func (p *PV1) NewPV1() *PV1 {
	return &PV1{}
}

func ParsePV1(parts []string) PV1 {
	parts = parts[1:] // Skip the segment identifier (PV1)

	return PV1{
		SetID:                   GetField(parts, 0),
		PatientClass:            GetField(parts, 1),
		AssignedPatientLocation: GetField(parts, 2),
		AdmissionType:           GetField(parts, 3),
		PreAdmitNumber:          GetField(parts, 4),
		PriorPatientLocation:    GetField(parts, 5),
		AttendingDoctor:         GetField(parts, 6),
		ReferringDoctor:         GetField(parts, 7),
		ConsultingDoctor:        GetField(parts, 8),
		HospitalService:         GetField(parts, 9),
		TemporaryLocation:       GetField(parts, 10),
		PreAdmitTestIndicator:   GetField(parts, 11),
		AdmitSource:             GetField(parts, 12),
		AmbulatoryStatus:        GetField(parts, 13),
		VipIndicator:            GetField(parts, 14),
		AdmittingDoctor:         GetField(parts, 15),
		PatientType:             GetField(parts, 16),
		VisitNumber:             GetField(parts, 17),
		FinancialClass:          GetField(parts, 18),
		ChargePriceIndicator:    GetField(parts, 19),
		CourtesyCode:            GetField(parts, 20),
		CreditRating:            GetField(parts, 21),
		ContractCode:            GetField(parts, 22),
		ContractEffectiveDate:   GetField(parts, 23),
		ContractAmount:          GetField(parts, 24),
		ContractPeriod:          GetField(parts, 25),
		InterestCode:            GetField(parts, 26),
		TransferToBadDebtCode:   GetField(parts, 27),
		BadDebtAgencyCode:       GetField(parts, 28),
		BadDebtTransferAmount:   GetField(parts, 29),
		BadDebtRecoveryAmount:   GetField(parts, 30),
		DeleteAccountIndicator:  GetField(parts, 31),
		DeleteAccountDate:       GetField(parts, 32),
		DischargeDisposition:    GetField(parts, 33),
		DischargedToLocation:    GetField(parts, 34),
		DietType:                GetField(parts, 35),
		ServicingFacility:       GetField(parts, 36),
		BedStatus:               GetField(parts, 37),
		AccountStatus:           GetField(parts, 38),
		PendingLocation:         GetField(parts, 39),
		PriorTemporaryLocation:  GetField(parts, 40),
		AdmitDateTime:           GetField(parts, 41),
		DischargeDateTime:       GetField(parts, 42),
		CurrentPatientBalance:   GetField(parts, 43),
		TotalCharges:            GetField(parts, 44),
		TotalAdjustments:        GetField(parts, 45),
		TotalPayments:           GetField(parts, 46),
	}
}
