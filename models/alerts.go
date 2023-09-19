package models

type Alerts struct {
	Legal      Legal      `json:"legal,omitempty"`
	Discipline Discipline `json:"discipline,omitempty"`
	Medical    Medical    `json:"medical,omitempty"`
	Other      Other      `json:"other,omitempty"`
}

type Legal struct {
	Description string `json:"description,omitempty"`
	ExpiresDate string `json:"expires_date,omitempty"`
}

type Discipline struct {
	Description string `json:"description,omitempty"`
	ExpiresDate string `json:"expires_date,omitempty"`
}

type Medical struct {
	Description string `json:"description,omitempty"`
	ExpiresDate string `json:"expires_date,omitempty"`
}

type Other struct {
	Description string `json:"description,omitempty"`
	ExpiresDate string `json:"expires_date,omitempty"`
}
