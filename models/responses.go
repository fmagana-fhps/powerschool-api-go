package models

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	LifeLength  string `json:"expires_in"`
	ExpiresAt   int64
}

type StudentResponse struct {
	Student Student `json:"student,omitempty"`
}
