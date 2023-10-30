package models

import "net/http"

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	LifeLength  string `json:"expires_in"`
	ExpiresAt   int64
}

type Response[T any] struct {
	StatusCode int
	Header     http.Header
	Body       T
}

type Tables struct {
	TransportationRecord TransportionRecord `json:"transportation,omitempty"`
}

type SchemaRecord struct {
	Id     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Tables Tables `json:"tables,omitempty"`
}

type PowerqueryResponse struct {
	Name       string           `json:"name,omitempty"`
	Records    []map[string]any `json:"record,omitempty"`
	Extensions string           `json:"@extensions,omitempty"`
}

type SchemaResponse struct {
	InsertCount int      `json:"insert_count,omitempty"`
	UpdateCount int      `json:"update_count,omitempty"`
	DeleteCount int      `json:"delete_count,omitempty"`
	Result      []Result `json:"result,omitempty"`
}

type SuccessMessage struct {
	ID  int    `json:"id,omitempty"`
	Ref string `json:"ref,omitempty"`
}

type Result struct {
	Status         string         `json:"status,omitempty"`
	Action         string         `json:"action,omitempty"`
	SuccessMessage SuccessMessage `json:"success_message,omitempty"`
}
