package database

import (
	"SE/src/Interface/admin/index"
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

type DataRet struct {
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

	err = stmt.QueryRow(userName).Scan(nil)
	if err != nil {
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

func GetSelfInfo(userName string) DataRet {
	stmt, err := DB.Prepare("select * from User where userName = ?")
	if err != nil {
		fmt.Println(err)
		return DataRet{
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
	return DataRet{
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
			Msg:     "update process err",
		}
	}
	defer stmt.Close()

	_ = stmt.QueryRow(params.UserName, params.Password, params.Avatar, oldUserName)
	return UpdateRet{
		Success: true,
	}
}
