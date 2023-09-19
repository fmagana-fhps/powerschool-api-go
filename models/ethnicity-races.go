package models

type Races struct {
	DistrictRaceCode             string `json:"district_race_code,omitempty"`
	DistrictRaceCodeDsescription string `json:"district_race_code_dsescription,omitempty"`
	FederalRaceCodeCategory      string `json:"federal_race_code_category,omitempty"`
}

type EthnicityRace struct {
	SchedulingReportingEthnicity string  `json:"scheduling_reporting_ethnicity,omitempty"`
	FederalEthnicity             string  `json:"federal_ethnicity,omitempty"`
	FederalRaceDeclineIndicator  string  `json:"federal_race_decline_indicator,omitempty"`
	Races                        []Races `json:"races,omitempty"`
}
