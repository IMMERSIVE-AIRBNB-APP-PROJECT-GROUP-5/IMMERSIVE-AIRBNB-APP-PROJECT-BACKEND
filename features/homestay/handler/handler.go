package handler

import (
	"io"
	"net/http"
	"net/url"

	"cloud.google.com/go/storage"
	"github.com/AIRBNBAPP/app/middlewares"
	"github.com/AIRBNBAPP/features/homestay"
	"github.com/AIRBNBAPP/helper"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
)

type HomestayHandler struct {
	homestayService homestay.HomestayServiceInterface
}

func New(handler homestay.HomestayServiceInterface) *HomestayHandler {
	return &HomestayHandler{
		homestayService: handler,
	}
}

func (handler *HomestayHandler) CreateHomestay(c echo.Context) error {
	var (
		storageClient *storage.Client
	)

	bucket := "test_buckety_go"
	folder := "Homestays"

	ctx := appengine.NewContext(c.Request())

	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile("keys.json"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"error":   true,
		})
	}

	file, err := c.FormFile("url")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	objectID := uuid.New().String()
	Object := folder + "/" + objectID + file.Filename

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

	// Bind data pengguna yang baru dari request body
	userInput := UserRequest{}
	// bind, membaca data yg dikirimkan client via request body
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data user"))
	}

	userInputUrl := homestay.PictureCore{
		Url: u.String(),
	}

	// mapping dari request ke core
	userCore := homestay.Core{
		Homestay_name: userInput.Homestay_name,
		Price:         userInput.Price,
		Total_guest:   userInput.Total_guest,
		Description:   userInput.Description,
		City_name:     userInput.City_name,
		Address:       userInput.Address,
		Picture:       userInputUrl,
		Facilities:    make([]homestay.FacilityCore, len(userInput.Facility_Id)),
	}

	for i, id := range userInput.Facility_Id {
		userCore.Facilities[i] = homestay.FacilityCore{
			Facility_Id: id,
		}
	}

	err = handler.homestayService.CreateHomestay(id, userCore)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("Selamat, anda berhasil menambahkan data homestay"))
}
