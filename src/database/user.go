package database

import (
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

	stmt, err := DB.Prepare("select exists (select * from User where userName = ? limit 1)")
	if err != nil {
		fmt.Println(err)
		return RegisterRet{
			Success: false,
			Msg:     "database error",
		}
	}
	defer stmt.Close()

	var num int
	err = stmt.QueryRow(user.Username).Scan(&num)
	if err != nil {
		fmt.Println(err)
		return RegisterRet{
			Success: false,
			Msg:     "database error",
		}
	}

	if num != 0 {
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
		return DeleteRet{
			Success: false,
			Msg:     "delete process error",
		}
	}
	return DeleteRet{
		Success: true,
	}
}
