package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"go_todo/config"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
)

var Db *sql.DB

var err error

// const (
// 	tableNameUser    = "users"
// 	tableNameTodo    = "todos"
// 	tableNameSession = "sessions"
// )

func InitDb() {
	err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatalln(err)
	}
	appEnv := config.Config.AppEnv
	if appEnv == "production" {
		url := os.Getenv("DATABASE_URL")
		connection, _ := pq.ParseURL(url)
		connection += "sslmode=require"
		Db, err = sql.Open(config.Config.SQLDriver, connection)
		if err != nil {
			log.Fatalln(err)
		}
	} else if appEnv == "develop" {
		log.Println(appEnv)
		connection := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("PASSWORD"))
		Db, err = sql.Open(config.Config.SQLDriver, connection)
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		log.Fatalln("error" + appEnv)
	}

	// Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
	// 	id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 	uuid STRING NOT NULL UNIQUE,
	// 	name STRING,
	// 	email STRING,
	// 	password STRING,
	// 	created_at DATETIME)`, tableNameUser)

	// _, err = Db.Exec(cmdU)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
	// 	id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 	content TEXT,
	// 	user_id INTEGER,
	// 	created_at DATETIME)`, tableNameTodo)

	// _, err = Db.Exec(cmdT)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
	// 	id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 	uuid STRING NOT NULL UNIQUE,
	// 	email STRING,
	// 	user_id INTEGER,
	// 	created_at DATETIME)`, tableNameSession)

	// _, err = Db.Exec(cmdS)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
