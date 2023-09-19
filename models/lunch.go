package models

type Lunch struct {
	ID       int    `json:"id,omitempty"`
	Balance1 int    `json:"balance1,omitempty"`
	Balance2 int    `json:"balance2,omitempty"`
	Balance3 int    `json:"balance3,omitempty"`
	Balance4 int    `json:"balance4,omitempty"`
	LastMeal string `json:"last_meal,omitempty"`
}
