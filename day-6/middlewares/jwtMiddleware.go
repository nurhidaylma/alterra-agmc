package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/nurhidaylma/alterra-agmc/day-6/config"
)

func CreateToken(userID int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenInString, err := token.SignedString([]byte(config.GetValue(config.SECRET_JWT)))
	if err != nil {
		return "", err
	}

	return tokenInString, nil
}

func ExtractTokenUserID(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userID := int(claims["user_id"].(float64))

		return userID
	}

	return 0
}
