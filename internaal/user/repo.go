package user

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repo struct {
	col *mongo.Collection	
}

func NewRepo(db *mongo.Database) *Repo{
	return &Repo{col: db.Collection("users")}
} 

func (r *Repo) FindByEmail(ctx context.Context, email string) (User, error){
	email = strings.ToLower(strings.TrimSpace(email))

	filter := bson.M{"email": email}

	var u User

	err := r.col.FindOne(ctx, filter).Decode(&u)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments){
			return User{}, mongo.ErrNoDocuments
		}
		return User{}, fmt.Errorf("Find by email failed: %w",err)
	}
	return u, nil
}

func (r *Repo) Create(ctx context.Context, u User) (User, error) {
	res, err := r.col.InsertOne(ctx, u)
	if err != nil{
		return User{},fmt.Errorf("Insert user failed: %w", err)
	}
	
	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return User{},fmt.Errorf("Insert user failed and Inserted id is not Object ID")
	}

	u.ID = id
	return u, nil

}