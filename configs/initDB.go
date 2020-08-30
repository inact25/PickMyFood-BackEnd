package configs

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func InitDB(c *Conf) (*sql.DB, error) {
	log.Println("Attempting to access database....")
	cfg := &mysql.Config{
		User:                 c.Db.DbUser,
		Passwd:               c.Db.DbPass,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%v:%v", c.Db.DbHost, c.Db.DbPort),
		DBName:               c.Db.DbSchema,
		AllowNativePasswords: c.Db.AllowNativePasswords,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Panic(err)
	}

	//Ping = check database availability
	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	log.Println("Connected to the database....")
	return db, nil
}
