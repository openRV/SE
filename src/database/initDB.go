package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

type Database struct {
	Path           string `toml:"path"`
	Type           string `toml:"type"`
	SqlSchemeInit  string `toml:"sqlSchemeInit"`
	SqlSchemeCheck string `toml:"sqlSchemeCheck"`
}

func InitDB(conf Database) error {

	_, err := os.Stat(conf.Path)
	if os.IsNotExist(err) {
		fmt.Println("database does not exist, creating one at: ", conf.Path)
		if InitScheme(conf) != nil {
			return err
		}
	} else {
		fmt.Println("found database at: ", conf.Path, " , type: ", conf.Type)
		if ok, err := CheckScheme(conf); err != nil {
			fmt.Println("check database error, exiting")
			return err
		} else if !ok {
			fmt.Println("database scheme error, creating a new one")
			postFix := 1
			for {
				postFix = postFix + 1
				if os.Rename(conf.Path, conf.Path+fmt.Sprint(postFix)) == nil {
					fmt.Println("last database: " + conf.Path + " renamed as: " + conf.Path + fmt.Sprint(postFix))
					break
				}

			}
			if InitScheme(conf) != nil {
				return err
			}
		}

	}

	db, err := sql.Open(conf.Type, conf.Path)
	if err != nil {
		return nil
	}

	DB = db

	return nil
}

func InitScheme(conf Database) error {

	file, err := os.Open(conf.SqlSchemeInit)
	if err != nil {
		fmt.Println("read database scheme error, exiting")
		fmt.Println(err)
		return err
	}
	defer file.Close()

	bytes, _ := ioutil.ReadAll(file)
	sqlStmt := string(bytes)

	db, err := sql.Open(conf.Type, conf.Path)
	if err != nil {
		fmt.Println("init database error, exiting")
		fmt.Println(err)
		return err
	}

	_, err = db.Exec(sqlStmt)

	if err != nil {
		fmt.Println("init database error, exiting")
		fmt.Println(err)
		return err
	}

	return nil
}

func CheckScheme(conf Database) (bool, error) {

	// check scheme

	return true, nil
}
