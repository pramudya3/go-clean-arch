package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pramudya3/go-clean-arch/domain"
)

func Signup(uc domain.UserUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		payload := &domain.SignUp{}

		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, domain.ResponseFailed(err.Error()))
			return
		}

		if _, err := uc.GetUserByEmail(c, payload.Email); err == nil {
			c.JSON(http.StatusConflict, domain.ResponseFailed("email already exist"))
			return
		}

		td, err := uc.Signup(c, payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ResponseFailed(err.Error()))
			return
		}

		c.JSON(http.StatusCreated, domain.ResponseSuccess(td))
	}
}
