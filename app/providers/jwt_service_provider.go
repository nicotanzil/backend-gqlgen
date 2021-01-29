package providers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/nicotanzil/backend-gqlgen/app/errorlist"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
	"github.com/twinj/uuid"
	"os"
	"time"
)

var (
	ACCESS_KEY  = []byte("AccessSecret")
	REFRESH_KEY = []byte("RefreshSecret")
)

func GenerateToken(user model.User) (*model.TokenDetail, error) {

	// Setting up the secret key
	os.Setenv("ACCESS_KEY", "AccessSecret")
	os.Setenv("REFRESH_KEY", "RefreshSecret")


	nowTime := time.Now().Unix()
	td := &model.TokenDetail{}

	// Access Token
	td.AccessExpires = time.Now().Add(time.Minute * 5).Unix()
	td.AccessUUID = uuid.NewV4().String()

	//Refresh Token
	td.RefreshExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUUID = uuid.NewV4().String()

	// Creating access accessToken header and payload
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"accountId": user.ID,
		"accessUuid": td.AccessUUID,
		"expiresAt": td.AccessExpires,	// expired after 5 minutes
		"issuesAt": nowTime,	// issued since now
	})

	var err error

	td.AccessToken, err = accessToken.SignedString([]byte(os.Getenv("ACCESS_KEY"))) // signature

	if err != nil {
		//If there is an errorlist in creating the JWT accessToken
		return nil, errorlist.JWT_Fail
	}

	//Creating refresh accessToken header and payload
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"accountId": user.ID,
		"refreshUuid": td.RefreshUUID,
		"expiresAt": td.RefreshExpires,	// expired after 5 minutes
		"issuesAt": nowTime,	// issued since now
	})

	td.RefreshToken, err = refreshToken.SignedString([]byte(os.Getenv("REFRESH_KEY"))) // signature

	if err != nil {
		//If there is an errorlist in creating the JWT accessToken
		return nil, errorlist.JWT_Fail
	}

	return td, nil
}

//func ParseToken(td *model.TokenDetail) (int, error) {
//	token, err := jwt.Parse(td.AccessToken, func(token *jwt.Token) (interface{}, error) {
//		return os.Getenv("ACCESS_KEY"), nil
//	})
//
//	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
//		id := claims["accountId"].(string)
//		return id, nil
//	} else {
//		return 0, err
//	}
//}