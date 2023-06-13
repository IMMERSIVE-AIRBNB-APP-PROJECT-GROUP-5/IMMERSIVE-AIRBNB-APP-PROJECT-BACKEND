package routes

import (
	"github.com/AIRBNBAPP/app/middlewares"
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

	e.POST("/register", userHandlerAPI.CreateUser)
	e.POST("/login", userHandlerAPI.Login)
	e.GET("/users", userHandlerAPI.GetUserById, jwtMiddleware)
	e.PUT("/users", userHandlerAPI.UpdateUserById, jwtMiddleware)
	// e.POST("/users/role", userHandlerAPI.CreateUser, jwtMiddleware)
	// e.PUT("/users/role/:id", userHandlerAPI.UpdateUser, jwtMiddleware)
	// e.DELETE("/users/role/:id", userHandlerAPI.DeleteUser, jwtMiddleware)

	// classData := _classData.New(db)

	// classService := _classService.New(classData)

	// classHandlerAPI := _classHandler.New(classService)

	// //EndPointClass
	// e.POST("/classes", classHandlerAPI.CreateClass, middlewares.JWTMiddleware())
	// e.GET("/classes", classHandlerAPI.GetAllClass, middlewares.JWTMiddleware())
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
