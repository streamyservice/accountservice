package dto

import "time"

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}
type JWTClaims struct {
	UserID   uint
	Username string
	Exp      time.Time `json:"exp"`
}
