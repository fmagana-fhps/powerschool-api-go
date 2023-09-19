package models

type Counselor struct {
	ID        int    `json:"id,omitempty"`
	UserDcid  int    `json:"userDcid,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	IsPrimary bool   `json:"isPrimary,omitempty"`
}

type Counselors struct {
	Counselor []Counselor `json:"counselor,omitempty"`
}
