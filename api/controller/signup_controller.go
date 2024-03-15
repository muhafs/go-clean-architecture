package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhafs/go-clean-architecture/bootstrap"
	"github.com/muhafs/go-clean-architecture/domain/entity"
	"github.com/muhafs/go-clean-architecture/domain/model/request"
	"github.com/muhafs/go-clean-architecture/domain/model/response"
	"github.com/muhafs/go-clean-architecture/domain/usecase"
	"github.com/muhafs/go-clean-architecture/helper/psw"
)

type SignupController struct {
	SignupUsecase usecase.SignupUsecase
	Env           *bootstrap.Env
}

func (sc *SignupController) Signup(c *gin.Context) {
	// parse request
	var request request.SignupRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Success: false, Message: err.Error()})
		return
	}

	// check user existence
	if _, err := sc.SignupUsecase.FindByEmail(request.Email); err == nil {
		c.JSON(http.StatusConflict, response.ErrorResponse{Success: false, Message: "User already exists with the given email"})
		return
	}

	// hash password
	hashedPassword, err := psw.Hash(request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Success: false, Message: err.Error()})
		return
	}
	request.Password = string(hashedPassword)

	// create user
	user := entity.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	if err := sc.SignupUsecase.Create(c, &user); err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	// create access token here

	// create refresh token here

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "",
		Data: map[string]any{
			"access_token":  "",
			"refresh_token": "",
		},
	})
}
