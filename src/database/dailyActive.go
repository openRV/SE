package database

import (
	"database/sql"
	"fmt"
)

type AddDailyLoginRet struct {
	Success bool
	Msg     string
}

func AddtoDailyLogin(userName string) AddDailyLoginRet {
	_, err := DB.Exec("select * from DailyLogin where userName = $1", userName)
	if err == sql.ErrNoRows {
		DB.Exec("insert into DailyLogin values ($1)", userName)
		return AddDailyLoginRet{Success: false, Msg: "database error"}
	} else {
		return AddDailyLoginRet{Success: true}
	}
}

type ActiveUserNumRet struct {
	Success bool
	Msg     string
	Num     int
}

func GetActiveUserNum() ActiveUserNumRet {
	rows, err := DB.Query("select userName from DailyLogin")
	if err != nil {
		fmt.Println(err)
		return ActiveUserNumRet{Success: false, Msg: "database error"}
	}
	defer rows.Close()

	activeNum := 0
	for rows.Next() {
		activeNum += 1
	}

	return ActiveUserNumRet{Success: true, Num: activeNum, Msg: "Active user num is ok"}
}
