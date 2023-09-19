package models

type ScheduleSetup struct {
	HomeRoom           string `json:"home_room,omitempty"`
	NextSchool         int    `json:"next_school,omitempty"`
	SchedNextYearGrade int    `json:"sched_next_year_grade,omitempty"`
}
