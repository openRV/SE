// @Title user.go
// @Description 操作数据库中 user 表相关的函数及相关的数据类型
// @Author 杜沛然 ${DATE} ${TIME}
package database

import (
	"SE/src/interface/admin/index"
	"fmt"
	"time"
)

type User struct {
	Username string
	Password string
}

type UserSearchRet struct {
	Exist      bool
	Password   bool
	RegistDate string
	Role       string
	Avatar     string
}

type RegisterRet struct {
	Success bool
	Msg     string
}

type DeleteRet struct {
	Success bool
	Msg     string
}

type SelfDataRet struct {
	UserName string
	Password string
	Avatar   string
	Msg      string
}

type UpdateRet struct {
	Success bool
	Msg     string
}

func SearchUser(user User) UserSearchRet {

	stmt, err := DB.Prepare("select password , registDate , role , avatar from Users where userName = $1")
	if err != nil {
		fmt.Println(err)
		return UserSearchRet{
			Exist: false,
		}
	}
	defer stmt.Close()

	var ret = UserSearchRet{Exist: true, Password: true}
	var password string
	err = stmt.QueryRow(user.Username).Scan(&password, &ret.RegistDate, &ret.Role, &ret.Avatar)
	if err != nil {
		fmt.Println(err)
		return UserSearchRet{
			Exist: false,
		}
	}
	if user.Password != password {
		return UserSearchRet{
			Exist:    true,
			Password: false,
		}
	}

	return ret

}

func RegisterUser(user User) RegisterRet {

	stmt, err := DB.Prepare("SELECT * FROM Users WHERE userName = $1")
	if err != nil {
		fmt.Println(err)
		return RegisterRet{
			Success: false,
			Msg:     "database error",
		}
	}
	defer stmt.Close()

	err = stmt.QueryRow(user.Username).Scan(nil)
	if err == nil {
		fmt.Println(err)
		return RegisterRet{
			Success: false,
			Msg:     "username already been used",
		}
	}

	stmt, err = DB.Prepare("insert into Users(userName , password , registDate , role , avatar) values ($1,$2,$3,$4,$5)")
	if err != nil {
		fmt.Println(err)
		return RegisterRet{
			Success: false,
			Msg:     "database error",
		}
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.Password, time.Now().Format("2006-01-02 15:04:05"), "user", "https://ui-avatars.com/api/?name="+user.Username)
	if err != nil {
		fmt.Println(err)
		return RegisterRet{
			Success: false,
			Msg:     "insert process error",
		}
	}

	stmt, err = DB.Prepare("insert into Dir(dirId , dirName , owner , createDate , lastView ) values ($1,$2,$3,$4,$5)")
	if err != nil {
		fmt.Println(err)
		return RegisterRet{
			Success: false,
			Msg:     "database error",
		}
	}
	_, err = stmt.Exec(user.Username, user.Username, user.Username, time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		fmt.Println(err)
		return RegisterRet{
			Success: false,
			Msg:     "insert process error",
		}
	}

	return RegisterRet{
		Success: true,
	}

}

func DeteleUser(userName string) DeleteRet {
	stmt, err := DB.Prepare("select * from Users where userName = $1")
	if err != nil {
		fmt.Println(err)
		return DeleteRet{
			Success: false,
			Msg:     "database error",
		}
	}
	defer stmt.Close()

	row := stmt.QueryRow(userName).Scan(nil)
	if row == nil {
		fmt.Println(err)
		return DeleteRet{
			Success: false,
			Msg:     "There is not the user",
		}
	}

	stmt, err = DB.Prepare("delete from Users where userName = $1")
	_, err = stmt.Exec(userName)

	if err != nil {
		fmt.Println(err)
		return DeleteRet{
			Success: false,
			Msg:     "delete process error",
		}
	}
	return DeleteRet{
		Success: true,
	}
}

func GetSelfInfo(userName string) SelfDataRet {
	stmt, err := DB.Prepare("select * from Users where userName = $1")
	if err != nil {
		fmt.Println(err)
		return SelfDataRet{
			Msg: "select process err",
		}
	}
	defer stmt.Close()

	row := stmt.QueryRow(userName)
	var str1, str2, str3, str4, str5 string
	/*
		CREATE TABLE User(
		    userName VARCHAR not NULL PRIMARY KEY,
		    password VARCHAR not NULL,
		    registDate VARCHAR,
		    role CHAR(5),--区分管理员与普通用户
		    avatar VARCHAR
		);
	*/
	err = row.Scan(&str1, &str2, &str3, &str4, &str5)
	return SelfDataRet{
		UserName: str1,
		Password: str2,
		Avatar:   str5,
	}
}

func UpadateInfo(oldUserName string, params index.SetInfoParams) UpdateRet {
	stmt, err := DB.Prepare("update Users set username = $1,password = $2,avatar = $3 where userName = $4")
	if err != nil {
		fmt.Println(err)
		return UpdateRet{
			Success: false,
			Msg:     "database err",
		}
	}
	defer stmt.Close()

	_, err = stmt.Exec(params.UserName, params.Password, params.Avatar, oldUserName)
	if err != nil {
		fmt.Println(err)
		return UpdateRet{
			Success: false,
			Msg:     "update process err",
		}
	}
	return UpdateRet{
		Success: true,
	}
}

func RegisterAdmin(userName string, password string) RegisterRet {
	stmt, err := DB.Prepare("insert into Users(userName , password , registDate , role , avatar) values ($1,$2,$3,'admin',$4)")
	if err != nil {
		fmt.Println(err)
		return RegisterRet{
			Success: false,
			Msg:     "database  err",
		}
	}
	defer stmt.Close()

	_, err = stmt.Exec(userName, password, time.Now().Format("2006-01-02 15:04:05"), "https://ui-avatars.com/api/?name="+userName)
	if err != nil {
		fmt.Println(err)
		return RegisterRet{
			Success: false,
			Msg:     "insert process err",
		}
	}

	return RegisterRet{Success: true}
}

type AllDataRet struct {
	Data [][2]string
	Msg  string
}

func GetAllUser() AllDataRet {
	stmt, err := DB.Prepare("select userName,password from Users where role='user'")
	if err != nil {
		fmt.Println(err)
		return AllDataRet{Msg: "database err"}
	}

	rows, err := stmt.Query()
	if err != nil {
		fmt.Println(err)
		return AllDataRet{Msg: "select process err"}
	}
	defer rows.Close()

	var data [][2]string

	for rows.Next() {
		var str1, str2 string
		err = rows.Scan(&str1, &str2)
		if err != nil {
			fmt.Println(err)
			return AllDataRet{Msg: "data get proceess err"}
		}
		data = append(data, [2]string{str1, str2})
	}

	return AllDataRet{
		Data: data,
	}
}

type UserNumRet struct {
	Success bool
	UserNum int
	Msg     string
}

func GetUserNum() UserNumRet {
	row, err := DB.Query("select userName from Users where role ='user'")
	if err != nil {
		fmt.Println(err)
		return UserNumRet{
			Success: false,
			Msg:     "database error",
		}
	}
	defer row.Close()

	var ret UserNumRet
	ret.UserNum = 0
	for row.Next() {
		ret.UserNum += 1
	}
	ret.Msg = "user num is ok"
	ret.Success = true
	return ret
}

type NewUserNumRet struct {
	Success bool
	D_data  []index.D_UserIncreaseData
	M_data  []index.M_UserIncreaseData
	Msg     string
}

func GetNewUserNum() NewUserNumRet {
	rows, err := DB.Query("select registDate from Users order by registDate")
	if err != nil {
		fmt.Println(err)
		return NewUserNumRet{
			Success: false,
			Msg:     "database error",
		}
	}
	defer rows.Close()
	var D_data []index.D_UserIncreaseData
	var M_data []index.M_UserIncreaseData

	for rows.Next() {
		var date string
		rows.Scan(&date)

		D_date := date[0:10]

		if D_data == nil {
			tmp1 := index.D_UserIncreaseData{Date: D_date, Num: 0}
			D_data = append(D_data, tmp1)
		}
		D_nonExist := false
		for i, D_element := range D_data {
			if D_element.Date == D_date {
				D_data[i].Num += 1
				break
			}
			if i == len(D_data)-1 {
				D_nonExist = true
			}
		}
		if D_nonExist {
			tmp2 := index.D_UserIncreaseData{Date: D_date, Num: 1}
			D_data = append(D_data, tmp2)
		}

		M_date := date[0:7]

		if M_data == nil {
			tmp3 := index.M_UserIncreaseData{Month: M_date, Num: 0}
			M_data = append(M_data, tmp3)
		}
		M_nonExist := false
		for i, _ := range M_data {
			if M_data[i].Month == M_date {
				M_data[i].Num += 1
				break
			}
			fmt.Println(i, len(M_data))
			if i == len(M_data)-1 {

				M_nonExist = true
			}
		}
		if M_nonExist {
			tmp4 := index.M_UserIncreaseData{Month: M_date, Num: 1}
			M_data = append(M_data, tmp4)
		}
	}

	return NewUserNumRet{Success: true, D_data: D_data, M_data: M_data, Msg: "new user num is ok"}
}
