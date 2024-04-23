package user

import (
	"context"
	"fmt"

	"github.com/pramudya3/go-clean-arch/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (u *userUsecase) Signup(ctx context.Context, user *domain.SignUp) (*domain.TokenDetail, error) {
	newUser := &domain.User{
		ID:       primitive.NewObjectID(),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	jwtToken := &domain.JWTToken{
		User:      newUser,
		SecretJWT: u.SecretJWT,
		ExpiryAT:  u.ExpAT,
		ExpiryRT:  u.ExpRT,
	}

	at, err := jwtToken.CreateAccessToken()
	if err != nil {
		return nil, fmt.Errorf("generate access token error, %w", err)
	}

	rt, err := jwtToken.CreateRefreshToken()
	if err != nil {
		return nil, fmt.Errorf("generate refresh token error, %w", err)
	}

	newUser.AccessToken = at
	newUser.RefreshToken = rt
	if err := u.UserRepository.InsertUser(ctx, newUser); err != nil {
		return nil, fmt.Errorf("failed to create new user: %w", err)
	}

	return &domain.TokenDetail{
		AccessToken:  at,
		RefreshToken: rt,
	}, nil
}
