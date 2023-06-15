package routes

import (
	"github.com/AIRBNBAPP/app/middlewares"
	_homestayData "github.com/AIRBNBAPP/features/homestay/data"
	_homestayHandler "github.com/AIRBNBAPP/features/homestay/handler"
	_homestayService "github.com/AIRBNBAPP/features/homestay/service"
	_userData "github.com/AIRBNBAPP/features/user/data"
	_userHandler "github.com/AIRBNBAPP/features/user/handler"
	_userService "github.com/AIRBNBAPP/features/user/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoute(e *echo.Echo, db *gorm.DB) {
	userData := _userData.New(db)
	userService := _userService.New(userData)
	userHandlerAPI := _userHandler.New(userService)

	// Register middleware
	jwtMiddleware := middlewares.JWTMiddleware()

	// User Routes

	e.POST("/register", userHandlerAPI.Register)
	e.POST("/login", userHandlerAPI.Login)
	e.GET("/users", userHandlerAPI.Profil, jwtMiddleware)
	e.PUT("/users", userHandlerAPI.UpdatedProfil, jwtMiddleware)
	e.DELETE("/users", userHandlerAPI.DeleteAccount, jwtMiddleware)
	e.POST("/users/validate", userHandlerAPI.ValidateHoster, jwtMiddleware)
	// e.POST("/users/role", userHandlerAPI.CreateUser, jwtMiddleware)
	// e.PUT("/users/role/:id", userHandlerAPI.UpdateUser, jwtMiddleware)
	// e.DELETE("/users/role/:id", userHandlerAPI.DeleteUser, jwtMiddleware)

	homestayData := _homestayData.New(db)

	homestayService := _homestayService.New(homestayData)

	homestayHandlerAPI := _homestayHandler.New(homestayService)

	//EndPointClass
	e.POST("/homestays", homestayHandlerAPI.CreateHomestay, jwtMiddleware)
	e.GET("/homestays", homestayHandlerAPI.GetAllHomestay)
	// e.PUT("/classes/:id", classHandlerAPI.UpdateClassById, middlewares.JWTMiddleware())

	// menteeData := _menteeData.New(db)

	// menteeService := _menteeService.New(menteeData)

	// menteeHandlerAPI := _menteeHandler.New(menteeService)

	// //EndPointMentee
	// e.POST("/mentees", menteeHandlerAPI.CreateMentee, middlewares.JWTMiddleware())
	// e.GET("/mentees", menteeHandlerAPI.GetAllMentee, middlewares.JWTMiddleware())
	// e.PUT("/mentees/:id", menteeHandlerAPI.UpdateMentee, middlewares.JWTMiddleware())
	// e.DELETE("/mentees/:id", menteeHandlerAPI.DeleteMentee, middlewares.JWTMiddleware())

	// logData := _logData.New(db)

	// logService := _logService.New(logData)

	// logHandlerAPI := _logHandler.New(logService)

	// // Log Routes
	// e.POST("/logs", logHandlerAPI.CreateLog, middlewares.JWTMiddleware())
	// e.GET("/logs", logHandlerAPI.GetLogsByMenteeID, middlewares.JWTMiddleware())
}
