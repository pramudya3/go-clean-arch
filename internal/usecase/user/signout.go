package user

import (
	"context"
	"fmt"
)

func (u *userUsecase) Signout(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, u.Timeout)
	defer cancel()

	user, err := u.UserRepository.FindByID(ctx, id)
	if err != nil {
		return fmt.Errorf("%+v", err)
	}

	user.AccessToken = ""
	user.RefreshToken = ""
	if err := u.UserRepository.UpdateUser(ctx, user); err != nil {
		return fmt.Errorf("%+v", err)
	}

	return nil
}
