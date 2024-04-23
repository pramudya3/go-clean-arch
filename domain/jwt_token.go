package domain

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type JWTAccessClaims struct {
	jwt.RegisteredClaims
	ID    string `json:"id"`
	Email string `json:"email"`
}

type JWTRefreshClaims struct {
	jwt.RegisteredClaims
	ID string `json:"id"`
}

type JWTToken struct {
	User      *User
	SecretJWT string
	ExpiryAT  int
	ExpiryRT  int
}

type TokenDetail struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (j *JWTToken) CreateAccessToken() (accessToken string, err error) {
	claims := &JWTAccessClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour * time.Duration(j.ExpiryAT))},
			ID:        uuid.NewString(),
		},
		ID:    j.User.ID.Hex(),
		Email: j.User.Email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.SecretJWT))
	if err != nil {
		return "", fmt.Errorf("failed to create token")
	}

	return t, nil
}

func (j *JWTToken) CreateRefreshToken() (refreshToken string, err error) {
	claimsRefresh := &JWTRefreshClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour * time.Duration(j.ExpiryRT))},
			ID:        uuid.NewString(),
		},
		ID: j.User.ID.String(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	rt, err := token.SignedString([]byte(j.SecretJWT))
	if err != nil {
		return "", fmt.Errorf("failed to create refresh token")
	}
	return rt, nil
}

func ExtractIDToken(token, secret string) (string, error) {
	tokenClaims, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := tokenClaims.Claims.(jwt.MapClaims)
	if !ok && !tokenClaims.Valid {
		return "", fmt.Errorf("invalid token")
	}

	return claims["id"].(string), nil
}

func IsTokenAuthorized(token, secret string) (bool, error) {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

type JWTTokenUsecase interface {
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
	ExtractIDToken(token string, secret string) (string, error)
}
