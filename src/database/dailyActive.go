// @Title dailyActive.go
// @Description 统计日活跃人数
// @Author 矫晓佳 ${DATE} ${TIME}
package database

import (
	"database/sql"
	"fmt"
)

// AddtoDailyLogin 在登陆时将用户添加日活跃表中(DailyLogin)的返回结果
type AddDailyLoginRet struct {
	Success bool   //插入是否成功
	Msg     string //错误信息
}

// @title AddtoDailyLogin
// @description 用于向日活跃表中添加用户
// @author 矫晓佳 ${DATE} ${TIME}
// @param userName string “要添加的用户的用户名”
// @return _ AddDailyLoginRet “包含是否成功及错误信息”
func AddtoDailyLogin(userName string) AddDailyLoginRet {
	//首先在 DailyLogin 中查询该用户今天是否登陆过
	_, err := DB.Exec("select * from DailyLogin where userName = $1", userName)
	if err == sql.ErrNoRows {
		//没有登陆过则执行插入操作
		DB.Exec("insert into DailyLogin values ($1)", userName)
		return AddDailyLoginRet{Success: true}
	} else {
		//已经登陆过则直接返回已经成功插入
		return AddDailyLoginRet{Success: true}
	}
}

// ActiveUserNumRet 统计活跃用户的返回结果
type ActiveUserNumRet struct {
	Success bool   //统计是否成功
	Msg     string //如果不成功，填写封装后的错误信息
	Num     int    //统计结果：今日活跃人数
}

// @title GetActiveUserNUm
// @description 用于统计今日活跃人数
// @auth 矫晓佳 ${DATE} ${TIME}
// @return _ ActiveUserNumRet "包含了今日活跃人数的结构体"
func GetActiveUserNum() ActiveUserNumRet {
	//从 DailyLogin 表中获取登陆过的 用户名
	rows, err := DB.Query("select userName from DailyLogin")
	if err != nil {
		fmt.Println(err)
		return ActiveUserNumRet{Success: false, Msg: "database error"}
	}
	defer rows.Close()

	//对结果进行遍历计数
	activeNum := 0
	for rows.Next() {
		activeNum += 1
	}
	//返回结果
	return ActiveUserNumRet{Success: true, Num: activeNum, Msg: "Active user num is ok"}
}
