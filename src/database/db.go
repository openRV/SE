package database

//
//import (
//	"database/sql"
//	"fmt"
//	"os"
//
//	_ "github.com/mattn/go-sqlite3"
//)
//
//func TestDB() {
//	os.Remove("./foo.db")
//
//	db, err := sql.Open("sqlite3", "./foo.db")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer db.Close()
//
//	sqlStmt := `
//	create table foo (id integer not null primary key, name text);
//	delete from foo;
//	`
//	_, err = db.Exec(sqlStmt)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	tx, err := db.Begin()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	stmt, err := tx.Prepare("insert into foo(id, name) values (?,?)")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer stmt.Close()
//
//	for i := 0; i < 100; i++ {
//		_, err = stmt.Exec(i, fmt.Sprintf("Hello world %03d", i))
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//	}
//	err = tx.Commit()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	rows, err := db.Query("select id,name from foo")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer rows.Close()
//
//	for rows.Next() {
//		var id int
//		var name string
//		err = rows.Scan(&id, &name)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//	}
//	err = rows.Err()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	stmt, err = db.Prepare("select name from foo where id = ?")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer stmt.Close()
//	var name string
//	err = stmt.QueryRow("3").Scan(&name)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(name)
//
//	_, err = db.Exec("delete from foo")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	_, err = db.Exec("insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	rows, err = db.Query("select id, name from foo")
//	if err != nil {
//		fmt.Println(err)
//	}
//	defer rows.Close()
//	for rows.Next() {
//		var id int
//		var name string
//		err = rows.Scan(&id, &name)
//		if err != nil {
//			fmt.Println(err)
//		}
//		fmt.Println(id, name)
//	}
//	err = rows.Err()
//	if err != nil {
//		fmt.Println(err)
//	}
//
//}
//
