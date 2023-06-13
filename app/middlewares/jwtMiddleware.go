package middlewares

import (
	"errors"

	"time"

	"github.com/AIRBNBAPP/app/config"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.JWT([]byte(config.SECRET_JWT))
}

func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(config.SECRET_JWT))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
func ExtractTokenUserId(c echo.Context) (int, error) {
	user := c.Get("user")
	if user == nil {
		return 0, errors.New("failed to extract user ID from token")
	}
	token, ok := user.(*jwt.Token)
	if !ok {
		return 0, errors.New("failed to extract user ID from token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("failed to extract user ID from token")
	}
	userID := int(claims["userId"].(float64))
	return userID, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(config.SECRET_JWT), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
