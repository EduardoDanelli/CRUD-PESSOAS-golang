package database

import (
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" // Postgres Drive
)

// SetupDB é o método responsável por abrir a conexão com o banco de dados
func SetupDB() *gorm.DB {
	user := "postgres"
	pass := "5628250"
	dbname := "postgres"
	host := "localhost"
	logMode, _ := strconv.ParseBool("true")

	settings := "host=" + host + " user=" + user + " password=" + pass + " dbname=" + dbname + " sslmode=disable"

	db, err := gorm.Open("postgres", settings)
	if err != nil {
		panic(err)
	}

	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}

	db.BlockGlobalUpdate(true)
	db.LogMode(logMode)

	return db
}

