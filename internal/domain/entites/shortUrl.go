package entites

import "time"

type ShortenUrl struct {
	CreatedAt  time.Time `json:"created_at"`
	ExpiresAt  time.Time `json:"expires_at"`
	RealUrl    string    `json:"real-url"`
	Identifier string    `json:"identifier"`
	Usages     int       `json:"usages"`
	ID         uint32    `json:"id"`
}

type InputUrl struct {
	RealUrl    string `json:"real-url"`
	Identifier string `json:"identifier"`
}
