package models

type InitialEnrollment struct {
	DistrictEntryDate       string `json:"district_entry_date,omitempty"`
	DistrictEntryGradeLevel int    `json:"district_entry_grade_level,omitempty"`
	SchoolEntryDate         string `json:"school_entry_date,omitempty"`
	SchoolEntryGradeLevel   int    `json:"school_entry_grade_level,omitempty"`
}
