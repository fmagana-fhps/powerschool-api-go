package models

type FullTimeEquivelancy struct {
	Fteid int    `json:"fteid,omitempty"`
	Name  string `json:"name,omitempty"`
}

type SchoolEnrollment struct {
	StatusCode              int                   `json:"status_code,omitempty"`
	EnrollStatus            string                `json:"enroll_status,omitempty"`
	EnrollStatusDescription string                `json:"enroll_status_description,omitempty"`
	GradeLevel              string                `json:"grade_level,omitempty"`
	EntryDate               string                `json:"entry_date,omitempty"`
	ExitDate                string                `json:"exit_date,omitempty"`
	SchoolNumber            int                   `json:"school_number,omitempty"`
	EntryCode               string                `json:"entry_code,omitempty"`
	EntryComment            string                `json:"entry_comment,omitempty"`
	ExitCode                string                `json:"exit_code,omitempty"`
	ExitComment             string                `json:"exit_comment,omitempty"`
	DistrictOfResidence     string                `json:"district_of_residence,omitempty"`
	Track                   string                `json:"track,omitempty"`
	FullTimeEquivelancy     []FullTimeEquivelancy `json:"full_time_equivelancy,omitempty"`
}
