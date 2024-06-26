package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Username string `json:"username"`
	UserId   int64  `json:"userId"`
}
type JwtPayload struct {
	User User `json:"user"`
}

const jwtSecretKey = "secret" // TODO: use env variable

func GenerateJwtToken(payload JwtPayload) string {
	claims := jwt.MapClaims{
		"user": map[string]interface{}{
			"username": payload.User.Username,
			"userId":   payload.User.UserId,
		},
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		panic(err)
	}
	return tokenString
}

func VerifyJwtToken(tokenString string) (JwtPayload, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensures that the type of token.Method is *jwt.SigningMethodHMAC (HS256 is actually a variant of HMAC)
		_, isSameMethod := token.Method.(*jwt.SigningMethodHMAC)
		if !isSameMethod {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(jwtSecretKey), nil
	})
	if err != nil {
		return JwtPayload{}, err
	}
	if !parsedToken.Valid {
		return JwtPayload{}, jwt.ErrSignatureInvalid
	}
	claims := parsedToken.Claims.(jwt.MapClaims)
	return JwtPayload{
		User: User{
			Username: claims["user"].(map[string]interface{})["username"].(string),
			UserId:   int64(claims["user"].(map[string]interface{})["userId"].(float64)),
		},
	}, nil
}
