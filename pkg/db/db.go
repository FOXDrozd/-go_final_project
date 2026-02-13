package db

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

const schema = `
    CREATE TABLE scheduler (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        date CHAR(16) NOT NULL DEFAULT "",
        title VARCHAR(256) NOT NULL DEFAULT "",
        comment TEXT NOT NULL DEFAULT "",
        repeat VARCHAR(128) NOT NULL DEFAULT ""
    );

    CREATE INDEX idx_scheduler_date ON scheduler(date);   
`

func Init(dbFile string) error {
	_, err := os.Stat(dbFile)
	install := os.IsNotExist(err)

	DB, err = sql.Open("sqlite", dbFile)

	if err != nil {
		return ErrCannotInitDatabase
	}

	defer DB.Close()

	if install {
		_, err = DB.Exec(schema)
		if err != nil {
			return ErrCannotInitDatabase
		}
	}

	return nil
}
