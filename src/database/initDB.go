package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB(Type string, Path string) error {

	_, err := os.Stat(Path)
	if os.IsNotExist(err) {
		fmt.Println("database does not exist, creating one at: ", Path)
	} else {
		fmt.Println("found database at: ", Path, " , type: ", Type)
	}

	db, err := sql.Open(Type, Path)
	if err != nil {
		return nil
	}

	DB = db

	return nil
}
