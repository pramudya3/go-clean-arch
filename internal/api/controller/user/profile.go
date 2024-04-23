package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pramudya3/go-clean-arch/domain"
)

func GetProfile(uc domain.UserUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetString("x-user-id")

		if err := uc.ValidateToken(c, id); err != nil {
			c.JSON(http.StatusUnauthorized, domain.ResponseFailed("Unathorized"))
			return
		}

		profile, err := uc.GetUserByID(c, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ResponseFailed(err.Error()))
			return
		}

		c.JSON(http.StatusOK, domain.ResponseSuccess(profile))
	}
}
