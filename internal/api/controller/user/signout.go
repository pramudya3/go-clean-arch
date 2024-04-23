package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pramudya3/go-clean-arch/domain"
)

func Signout(uc domain.UserUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetString("x-user-id")

		if err := uc.ValidateToken(c, id); err != nil {
			c.JSON(http.StatusUnauthorized, domain.ResponseFailed("Unathorized"))
			return
		}

		if err := uc.Signout(c, id); err != nil {
			c.JSON(http.StatusInternalServerError, domain.ResponseFailed(err.Error()))
			return
		}

		c.Status(http.StatusOK)
	}
}
