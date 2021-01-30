package providers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/nicotanzil/backend-gqlgen/app/errorlist"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
	"os"
	"time"
)

func GenerateToken(user model.User) (string, error) {

	// Setting up the secret key
	os.Setenv("ACCESS_KEY", "AccessSecret")


	nowTime := time.Now().Unix()

	// Creating access accessToken header and payload
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"accountId": user.ID,
		"issuedAt": nowTime,	// issued since now
	})

	var err error

	token, err := accessToken.SignedString([]byte(os.Getenv("ACCESS_KEY"))) // signature

	if err != nil {
		//If there is an errorlist in creating the JWT accessToken
		return "", errorlist.JWT_Fail
	}

	return token, nil
}

//func GenerateRefreshToken(user model.User) (string, error) {
//
//	// Setting up the secret key
//	os.Setenv("REFRESH_KEY", "RefreshSecret")
//
//	nowTime := time.Now().Unix()
//
//	// Creating access accessToken header and payload
//	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
//		"accountId": user.ID,
//		"issuedAt": nowTime,	// issued since now
//	})
//
//	var err error
//
//	token, err := refreshToken.SignedString([]byte(os.Getenv("REFRESH_KEY"))) // signature
//
//	if err != nil {
//		//If there is an errorlist in creating the JWT accessToken
//		return "", errorlist.JWT_Fail
//	}
//
//	return token, nil
//}

func ParseToken(tokenParam string) (float64, error) {
	token, err := jwt.Parse(tokenParam, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_KEY")), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := claims["accountId"].(float64)
		return id, nil
	} else {
		return 0, err
	}
}

//func ParseRefreshToken(tokenParam string) (float64, error) {
//	token, err := jwt.Parse(tokenParam, func(token *jwt.Token) (interface{}, error) {
//		return []byte(os.Getenv("REFRESH_KEY")), nil
//	})
//
//	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
//		id := claims["accountId"].(float64)
//		return id, nil
//	} else {
//		return 0, err
//	}
//
//}