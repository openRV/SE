package database

func Test() {
	sqlStmt := `
	create table foo (id integer not null primary key, name text);
	delete from foo;
	`
	_, err := DB.Exec(sqlStmt)
	if err != nil {
		return
	}
}
