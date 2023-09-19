package models

type Student struct {
	Expansions        string             `json:"@expansions,omitempty"`
	Extensions        string             `json:"@extensions,omitempty"`
	ExtensionData     ExtensionData      `json:"_extension_data,omitempty"`
	ID                int                `json:"id,omitempty"`
	LocalID           int                `json:"local_id,omitempty"`
	StateProvinceID   string             `json:"state_province_id,omitempty"`
	StudentUsername   string             `json:"student_username,omitempty"`
	Name              Name               `json:"name,omitempty"`
	Addresses         Addresses          `json:"addresses,omitempty"`
	Alerts            []Alerts           `json:"alerts,omitempty"`
	Phones            []Phones           `json:"phones,omitempty"`
	SchoolEnrollment  []SchoolEnrollment `json:"school_enrollment,omitempty"`
	EthnicityRace     []EthnicityRace    `json:"ethnicity_race,omitempty"`
	Contact           Contact            `json:"contact,omitempty"`
	ContactInfo       ContactInfo        `json:"contact_info,omitempty"`
	InitialEnrollment InitialEnrollment  `json:"initial_enrollment,omitempty"`
	ScheduleSetup     ScheduleSetup      `json:"schedule_setup,omitempty"`
	Fees              Fees               `json:"fees,omitempty"`
	Lunch             Lunch              `json:"lunch,omitempty"`
	Counselors        Counselors         `json:"counselors,omitempty"`
	GlobalID          string             `json:"global_id,omitempty"`
}

type Name struct {
	FirstName  string `json:"first_name,omitempty"`
	MiddleName string `json:"middle_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
}
