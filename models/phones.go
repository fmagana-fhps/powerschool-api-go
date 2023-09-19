package models

type Main struct {
	Number string `json:"number,omitempty"`
}

type Phones struct {
	Main Main `json:"main,omitempty"`
}
