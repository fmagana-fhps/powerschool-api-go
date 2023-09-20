package models

type TransportionRecord struct {
	Studentid           string `json:"studentid"`
	Description         string `json:"description,omitempty"`
	Schoolid            string `json:"schoolid,omitempty"`
	Startdate           string `json:"startdate,omitempty"`
	Enddate             string `json:"enddate,omitempty"`
	Departuretime       string `json:"departuretime,omitempty"`
	Arrivaltime         string `json:"arrivaltime,omitempty"`
	FromTo              string `json:"fromto,omitempty"`
	Type                string `json:"type,omitempty"`
	Monday              string `json:"monday,omitempty"`
	Tuesday             string `json:"tuesday,omitempty"`
	Wednesday           string `json:"wednesday,omitempty"`
	Thursday            string `json:"thursday,omitempty"`
	Friday              string `json:"friday,omitempty"`
	Distance            string `json:"distance,omitempty"`
	Distanceindicator   string `json:"distanceindicator,omitempty"`
	Linkingcode         string `json:"linkingcode,omitempty"`
	Routenumber         string `json:"routenumber,omitempty"`
	Busnumber           string `json:"busnumber,omitempty"`
	Drivername          string `json:"drivername,omitempty"`
	Contactnumber       string `json:"contactnumber,omitempty"`
	Stopnumber          string `json:"stopnumber,omitempty"`
	Address             string `json:"address,omitempty"`
	Specialinstructions string `json:"specialinstructions,omitempty"`
}
