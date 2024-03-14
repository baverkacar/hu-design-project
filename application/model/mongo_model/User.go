package mongo_model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"hu-design-project/application/model"
	"hu-design-project/application/model/dto"
	"time"
)

type User struct {
	UserID     primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FullName   string             `bson:"fullName" json:"fullName"`
	Email      string             `bson:"email" json:"email"`
	Password   string             `bson:"password" json:"password"`
	IsActive   bool               `bson:"isActive" json:"isActive"`
	CreatedAt  time.Time          `bson:"createdAt" json:"createdAt"`
	LastUpdate time.Time          `bson:"lastUpdate" json:"lastUpdate"`
}

func NewUser(userCreateModel model.UserCreateModel) *User {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userCreateModel.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return &User{
		FullName:   userCreateModel.FullName,
		Email:      userCreateModel.Email,
		Password:   string(hashedPassword),
		IsActive:   false,
		CreatedAt:  time.Now(),
		LastUpdate: time.Now(),
	}
}

func (u *User) ToDTO(verificationCode string) *dto.UserDTO {
	return &dto.UserDTO{
		UserID:           u.UserID.Hex(),
		FullName:         u.FullName,
		Email:            u.Email,
		IsActive:         u.IsActive,
		CreatedAt:        u.CreatedAt,
		LastUpdate:       u.LastUpdate,
		VerificationCode: verificationCode,
	}
}