package models

type Contact struct {
	EmergencyContactName1 string `json:"emergency_contact_name1:omitempty"`
	EmergencyPhone1       string `json:"emergency_phone1:omitempty"`
	EmergencyContactName2 string `json:"emergency_contact_name2:omitempty"`
	EmergencyPhone2       string `json:"emergency_phone2:omitempty"`
	GuardianEmail         string `json:"guardian_email:omitempty"`
	GuardianFax           string `json:"guardian_fax:omitempty"`
	Mother                string `json:"mother:omitempty"`
	Father                string `json:"father:omitempty"`
	DoctorName            string `json:"doctor_phone:omitempty"`
	DoctorPhone           string `json:"doctor_name:omitempty"`
}

type ContactInfo struct {
	Email string `json:"email,omitempty"`
}
