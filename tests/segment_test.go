package tests

import (
	"reflect"
	"testing"

	"github.org/wpartin/go/hl7-parser/internal/parser"
)

func TestAL1Segment(t *testing.T) {
	segment := "\rAL1|1||^Penicillin||Moderate|20050101\r"

	tests := []struct {
		name     string
		message  string
		expected string
	}{
		{"SetID", segment, "1"},
		{"AllergenTypeCode", segment, ""},
		{"AllergenCode", segment, "^Penicillin"},
		{"AllergySeverityCode", segment, ""},
		{"AllergyReaction", segment, "Moderate"},
		{"IdentificationDate", segment, "20050101"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			message, err := parser.ParseMessage(test.message)

			if err != nil {
				t.Errorf("failed with %v", err.Error())
			}

			al1s := message.AL1
			r := reflect.ValueOf(al1s[0])
			if r.FieldByName(test.name).String() != test.expected {
				t.Errorf("expected %s to be %v, got %v", test.name, test.expected, r.FieldByName(test.name).Interface())
			}
		})
	}
}

func TestMSHSegment(t *testing.T) {
	segment := "MSH|^~\\&|SendingApp|SendingFac|ReceivingApp|ReceivingFac|20230101||ADT^A01|12345|P|2.3\r"

	tests := []struct {
		name     string
		message  string
		expected string
	}{
		{"FieldSeparator", segment, "|"},
		{"EncodingCharacters", segment, "^~\\&"},
		{"SendingApplication", segment, "SendingApp"},
		{"SendingFacility", segment, "SendingFac"},
		{"ReceivingApplication", segment, "ReceivingApp"},
		{"ReceivingFacility", segment, "ReceivingFac"},
		{"DateTimeOfMessage", segment, "20230101"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			message, err := parser.ParseMessage(test.message)

			if err != nil {
				t.Errorf("failed with %v", err.Error())
			}

			msh := message.MSH
			r := reflect.ValueOf(msh)
			if r.FieldByName(test.name).String() != test.expected {
				t.Errorf("expected %s to be %v, got %v", test.name, test.expected, r.FieldByName(test.name).Interface())
			}
		})
	}
}

func TestPV1Segment(t *testing.T) {
	segment := "PV1|1|I|2000^2012^01||||1234^Doe^John||SUR|||||||||1234567890||||||||||||||||||||||||20230101\r"

	tests := []struct {
		name     string
		message  string
		expected string
	}{
		{"SetID", segment, "1"},
		{"PatientClass", segment, "I"},
		{"AssignedPatientLocation", segment, "2000^2012^01"},
		{"AttendingDoctor", segment, "1234^Doe^John"},
		{"VisitNumber", segment, "1234567890"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			message, err := parser.ParseMessage(test.message)

			if err != nil {
				t.Errorf("failed with %v", err.Error())
			}

			pv1 := message.PV1
			r := reflect.ValueOf(pv1)
			if r.FieldByName(test.name).String() != test.expected {
				t.Errorf("expected %s to be %v, got %v", test.name, test.expected, r.FieldByName(test.name).Interface())
			}
		})
	}
}

func TestPIDSegment(t *testing.T) {
	segment := "PID|1||123456^^^Hospital^MR||Doe^John^A||19800101|M|||123 Main St^^Hometown^IL^12345||(555)555-5555|||||987654321|N\r"

	tests := []struct {
		name     string
		message  string
		expected string
	}{
		{"SetID", segment, "1"},
		{"PatientIDInternal", segment, "123456^^^Hospital^MR"},
		{"PatientName", segment, "Doe^John^A"},
		{"DateOfBirth", segment, "19800101"},
		{"Sex", segment, "M"},
		{"PatientAddress", segment, "123 Main St^^Hometown^IL^12345"},
		{"PhoneNumberHome", segment, "(555)555-5555"},
		{"PatientAccountNumber", segment, "987654321"},
		{"SocialSecurityNumber", segment, "N"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			message, err := parser.ParseMessage(test.message)

			if err != nil {
				t.Errorf("failed with %v", err.Error())
			}

			pid := message.PID
			r := reflect.ValueOf(pid)
			if r.FieldByName(test.name).String() != test.expected {
				t.Errorf("expected %s to be %v, got %v", test.name, test.expected, r.FieldByName(test.name).Interface())
			}
		})
	}
}
