package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SignUp struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string             `bson:"name"`
	Email        string             `bson:"email"`
	Password     string             `bson:"password"`
	AccessToken  string             `bson:"access_token"`
	RefreshToken string             `bson:"refresh_token"`
}

type UserUsecase interface {
	Signup(ctx context.Context, user *SignUp) (*TokenDetail, error)
	Login(ctx context.Context, login *Login) (*TokenDetail, error)
	RefreshToken(token string) (*TokenDetail, error)
	Signout(ctx context.Context, id string) error

	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByID(ctx context.Context, id string) (*User, error)
	FetchUsers(ctx context.Context) ([]*User, error)

	UpdateUser(ctx context.Context, user *User) error

	ValidateToken(ctx context.Context, id string) error
}

type UserRepository interface {
	InsertUser(ctx context.Context, user *User) error

	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByID(ctx context.Context, id string) (*User, error)
	FetchAllUsers(ctx context.Context) ([]*User, error)

	UpdateUser(ctx context.Context, user *User) error
}
