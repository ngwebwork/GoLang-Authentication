package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email string `bson:"email" json:"email"`
	PasswordHash string `bson:"passwordhash" json:"-"`
	Role string `bson:"role" json:"role"`
	createdAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}

type PublicUser struct{
	ID string `json:"id"`
	Email string `json:"email"`
	Role string `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func ToPublic (u User)  PublicUser{
	return PublicUser{
		ID: u.ID.Hex(),
		Email: u.Email,
		Role: u.Role,
		CreatedAt: u.createdAt,
		UpdatedAt: u.UpdatedAt,
	}
}