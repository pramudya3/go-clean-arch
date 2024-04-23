package user

import "github.com/pramudya3/go-clean-arch/domain"

func (u *userUsecase) RefreshToken(token string) (*domain.TokenDetail, error) {
	//TODO: update access_token & refresh_token
	return nil, nil
}
