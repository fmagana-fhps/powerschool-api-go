package models

type Addresses struct {
	Physical Physical `json:"physical,omitempty"`
	Mailing  Mailing  `json:"mailing,omitempty"`
}

type Physical struct {
	Street        string `json:"street,omitempty"`
	City          string `json:"city,omitempty"`
	StateProvince string `json:"state_province,omitempty"`
	PostalCode    string `json:"postal_code,omitempty"`
	GridLocation  string `json:"grid_location,omitempty"`
}

type Mailing struct {
	Street        string `json:"street,omitempty"`
	City          string `json:"city,omitempty"`
	StateProvince string `json:"state_province,omitempty"`
	PostalCode    string `json:"postal_code,omitempty"`
	GridLocation  string `json:"grid_location,omitempty"`
}
