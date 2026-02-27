package models

import (
	"Sec_17-18/todoapp/config"
	"crypto/sha1"

	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const (
	tableNameUser    = "users"
	tableNameTodo    = "todos"
	tableNameSession = "sessions"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
	     id INTEGER PRIMARY KEY AUTOINCREMENT,
		 uuid STRING NOT NULL UNIQUE,
		 name STRING,
		 email STRING,
		 password STRING,
		 created_at DATETIME )`, tableNameUser)

	Db.Exec(cmdU)

	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
	      id INTEGER PRIMARY KEY AUTOINCREMENT,
		  content TEXT,
		  user_id INTEGER,
		  created_at DATETIME,
		  complete INTEGER(2))`, tableNameTodo)
	Db.Exec(cmdT)
	ensureTodoCompleteColumn()

	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		email STRING,
		user_id INTEGER,
		created_at DATETIME)`, tableNameSession)
	Db.Exec(cmdS)
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}

func ensureTodoCompleteColumn() {
	rows, err := Db.Query(`PRAGMA table_info(todos)`)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	var (
		cid        int
		name       string
		columnType string
		notnull    int
		defaultVal sql.NullString
		pk         int
	)

	for rows.Next() {
		err = rows.Scan(&cid, &name, &columnType, &notnull, &defaultVal, &pk)
		if err != nil {
			log.Fatalln(err)
		}
		if name == "complete" {
			return
		}
	}

	_, err = Db.Exec(`ALTER TABLE todos ADD COLUMN complete INTEGER(2)`)
	if err != nil {
		log.Fatalln(err)
	}
}
