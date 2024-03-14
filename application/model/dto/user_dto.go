package dto

import "time"

type UserDTO struct {
	UserID           string    `json:"userId"`
	FullName         string    `json:"fullName"`
	Email            string    `json:"email"`
	IsActive         bool      `bson:"isActive" json:"isActive"`
	VerificationCode string    `bson:"verificationCode" json:"verificationCode,omitempty"`
	CreatedAt        time.Time `bson:"createdAt" json:"createdAt"`
	LastUpdate       time.Time `bson:"lastUpdate" json:"lastUpdate"`
}
