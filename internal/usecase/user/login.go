package user

import (
	"context"

	"github.com/pramudya3/go-clean-arch/domain"
)

func (u *userUsecase) Login(ctx context.Context, login *domain.Login) (*domain.TokenDetail, error) {
	// TODO: generate access_token & refresh_token
	return nil, nil
}
