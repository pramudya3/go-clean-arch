package user

import (
	"context"
	"time"

	"github.com/pramudya3/go-clean-arch/domain"
)

type userUsecase struct {
	UserRepository domain.UserRepository
	Timeout        time.Duration
	ExpAT          int
	ExpRT          int
	SecretJWT      string
}

func (u *userUsecase) FetchUsers(ctx context.Context) ([]*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.Timeout)
	defer cancel()

	return u.UserRepository.FetchAllUsers(ctx)
}

func (u *userUsecase) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.Timeout)
	defer cancel()

	return u.UserRepository.FindByEmail(ctx, email)
}

func (u *userUsecase) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.Timeout)
	defer cancel()

	return u.UserRepository.FindByID(ctx, id)
}

func (u *userUsecase) UpdateUser(ctx context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(ctx, u.Timeout)
	defer cancel()

	return u.UserRepository.UpdateUser(ctx, user)
}

func NewUserUsecase(userRepository domain.UserRepository, env *domain.Env) domain.UserUsecase {
	return &userUsecase{
		UserRepository: userRepository,
		Timeout:        time.Duration(env.ContextTimeout) * time.Second,
		ExpAT:          env.ExpiryAccessToken,
		ExpRT:          env.ExpiryRefreshToken,
		SecretJWT:      env.SecretJWT,
	}
}
