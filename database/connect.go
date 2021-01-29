package connect

import(
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

const DB_NAME="staem"
const DB_HOST="127.0.0.1"
const DB_PORT="5432"				//port on installation
const DB_USER="postgres"			//default is postgres
const DB_PASS="password"	//password on installation

func Connect()(*gorm.DB, error){

	dsn := "host=127.0.0.1 user=postgres password=password dbname=staem port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}