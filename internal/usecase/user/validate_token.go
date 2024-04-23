package user

import (
	"context"
	"fmt"
)

func (u *userUsecase) ValidateToken(ctx context.Context, id string) error {
	user, err := u.UserRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if user.AccessToken == "" || user.RefreshToken == "" {
		return fmt.Errorf("Unauthorized")
	}

	// TODO: add validation token by expired time

	return nil
}
