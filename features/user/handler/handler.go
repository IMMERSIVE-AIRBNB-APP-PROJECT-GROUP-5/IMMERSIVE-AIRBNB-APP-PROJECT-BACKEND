package handler

import (
	"net/http"
	"strings"

	"github.com/AIRBNBAPP/features/user"
	"github.com/AIRBNBAPP/helper"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.UserServiceInterface
}

func New(handler user.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: handler,
	}
}

func (handler *UserHandler) CreateUser(c echo.Context) error {
	userInput := UserRequest{}
	// bind, membaca data yg dikirimkan client via request body
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data user"))
	}
	// mapping dari request ke core
	userCore := user.Core{
		User_name: userInput.User_name,
		Phone:     userInput.Phone,
		Email:     userInput.Email,
		Password:  userInput.Password,
	}
	err := handler.userService.CreateUser(userCore)
	if err != nil {
		if strings.Contains(err.Error(), "failed inserted data user") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("error inserted data user, row affected = 0"))
		} else {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
		}
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("Selamat, akun anda berhasil terdaftar !"))
}
