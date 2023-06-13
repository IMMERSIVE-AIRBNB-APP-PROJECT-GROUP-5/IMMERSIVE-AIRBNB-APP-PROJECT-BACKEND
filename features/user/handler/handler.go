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

func (handler *UserHandler) Login(c echo.Context) error {
	// Memeriksa apakah email dan password inputan dapat di bind
	loginInput := AuthRequest{}
	errBind := c.Bind(&loginInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}
	// Memeriksa apakah email & password telah diinputkan di database
	userData, token, err := handler.userService.Login(loginInput.Email, loginInput.Password)
	if err != nil {
		if strings.Contains(err.Error(), "login failed") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
			// Memeriksa validasi di database dan validasi lainnya
		} else {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
		}
	}

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Selamat, anda berhasil login !", map[string]any{
		"nama":   userData.User_name,
		"email":  userData.Email,
		"status": userData.Status,
		"token":  token,
	}))
}
