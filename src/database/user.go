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

	stmt, err := DB.Prepare("select password , registDate , role , avatar from User where userName = ?")
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

	stmt, err := DB.Prepare("select * from User where userName = ?")
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

	stmt, err = DB.Prepare("insert into User(userName , password , registDate , role , avatar) values (?,?,?,?,?)")
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
	return RegisterRet{
		Success: true,
	}

}

func DeteleUser(userName string) DeleteRet {
	stmt, err := DB.Prepare("select * from User where userName = ?")
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

	stmt, err = DB.Prepare("delete from User where userName = ?")
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
	stmt, err := DB.Prepare("select * from User where userName = ?")
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
	stmt, err := DB.Prepare("update User set username = ?,password = ?,avatar = ? where userName = ?")
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
	stmt, err := DB.Prepare("insert into User(userName , password , registDate , role , avatar) values (?,?,?,'admin',?)")
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
	stmt, err := DB.Prepare("select userName,password from User where role='user'")
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
	row, err := DB.Query("select count(userName) from User")
	if err != nil {
		fmt.Println(err)
		return UserNumRet{
			Success: false,
			Msg:     "database error",
		}
	}
	defer row.Close()

	var ret UserNumRet
	row.Scan(&ret.UserNum)
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
	rows, err := DB.Query("select registDate from User")
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
			D_data[0].Date = D_date
			D_data[0].Num = 1
		}
		D_nonExist := false
		for i, D_element := range D_data {
			if D_element.Date == D_date {
				D_element.Num += 1
			}
			if i == len(D_data)-1 {
				D_nonExist = true
			}
		}
		if D_nonExist {
			tmp := index.D_UserIncreaseData{Date: D_date, Num: 1}
			D_data = append(D_data, tmp)
		}

		M_date := date[0:7]

		if D_data == nil {
			D_data[0].Date = D_date
			D_data[0].Num = 1
		}
		M_nonExist := false
		for i, M_element := range M_data {
			if M_element.Month == M_date {
				M_element.Num += 1
			}
			if i == len(D_data)-1 {
				M_nonExist = true
			}
		}
		if M_nonExist {
			tmp := index.M_UserIncreaseData{Month: M_date, Num: 1}
			M_data = append(M_data, tmp)
		}
	}

	return NewUserNumRet{Success: true, D_data: D_data, M_data: M_data, Msg: "new user num is ok"}
}
