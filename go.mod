module github.com/nicotanzil/backend-gqlgen

// +heroku goVersion go1.15
go 1.15

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-redis/redis/v8 v8.7.1
	github.com/jackc/pgproto3/v2 v2.0.7 // indirect
	github.com/mailjet/mailjet-apiv3-go v0.0.0-20201009050126-c24bc15a9394
	github.com/vektah/gqlparser/v2 v2.1.0
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad
	golang.org/x/text v0.3.5 // indirect
	gorm.io/driver/postgres v1.0.7
	gorm.io/gorm v1.20.12
)
