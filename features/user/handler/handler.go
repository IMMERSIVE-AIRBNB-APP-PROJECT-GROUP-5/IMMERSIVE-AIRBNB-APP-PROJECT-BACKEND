package handler

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/AIRBNBAPP/app/middlewares"
	"github.com/AIRBNBAPP/features/user"
	"github.com/AIRBNBAPP/helper"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
)

type UserHandler struct {
	userService user.UserServiceInterface
}

func New(handler user.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: handler,
	}
}

func (handler *UserHandler) Register(c echo.Context) error {
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
	err := handler.userService.Register(userCore)
	if err != nil {
		if strings.Contains(err.Error(), "Gagal melakukan pendaftaran akun") {
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
		if strings.Contains(err.Error(), "Gagal melakukan login") {
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

func (handler *UserHandler) Profil(c echo.Context) error {
	// Mendapatkan ID pengguna yang login
	id, err := middlewares.ExtractTokenUserId(c)

	// Mengambil data pengguna berdasarkan ID
	results, err := handler.userService.Profil(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Gagal mendapatkan data pengguna"))
	}
	var userResponse []UserResponse
	for _, value := range results {
		userResponse = append(userResponse, UserResponse{
			User_name: value.User_name,
			Email:     value.Email,
			Phone:     value.Phone,
			Status:    UserStatus(value.Status),
		})
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Berhasil mendapatkan data pengguna", userResponse))
}

func (handler *UserHandler) UpdatedProfil(c echo.Context) error {
	// Mendapatkan ID pengguna yang login
	id, err := middlewares.ExtractTokenUserId(c)

	// Bind data pengguna yang baru dari request body
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

	err = handler.userService.UpdatedProfil(id, userCore)
	if err != nil {
		if strings.Contains(err.Error(), "Gagal memperbarui data pengguna") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("error updated data user, row affected = 0"))
		} else {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
		}
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("Berhasil memperbarui data pengguna"))
}

func (handler *UserHandler) DeleteAccount(c echo.Context) error {
	// Mendapatkan ID pengguna yang login
	id, err := middlewares.ExtractTokenUserId(c)

	// Hapus data pengguna dari database
	err = handler.userService.DeleteAccount(id)
	if err != nil {
		if strings.Contains(err.Error(), "Gagal menghapus data pengguna") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("error deleted data user, row affected = 0"))
		} else {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Pengguna tidak ditemukan"))
		}
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("Berhasil menghapus data pengguna"))
}

func (handler *UserHandler) ValidateHoster(c echo.Context) error {
	var (
		storageClient *storage.Client
	)

	bucket := "test_buckety_go"
	folder := "Users"

	ctx := appengine.NewContext(c.Request())

	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile("keys.json"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"error":   true,
		})
	}

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	Object := folder + "/" + file.Filename

	sw := storageClient.Bucket(bucket).Object(Object).NewWriter(ctx)

	if _, err := io.Copy(sw, src); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"error":   true,
		})
	}

	if err := sw.Close(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"error":   true,
		})
	}

	u, err := url.Parse("https://storage.googleapis.com/" + bucket + "/" + sw.Attrs().Name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"Error":   true,
		})
	}
	// Mendapatkan ID pengguna yang login
	id, err := middlewares.ExtractTokenUserId(c)

	userInput := user.Core{
		Url: u.String(),
	}

	err = handler.userService.ValidateHoster(id, userInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Gagal menambahkan foto"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("Selamat, anda memiliki akses menambahkan homestays"))
}
