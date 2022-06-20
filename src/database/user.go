package database

import "fmt"

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

func SearchUser(user User) UserSearchRet {

	stmt, err := DB.Prepare("select password , registDate , role , avarat from User where userName = ?")
	if err != nil {
		fmt.Println(err)
		return UserSearchRet{
			Exist: false,
		}
	}
	defer stmt.Close()

	var ret = UserSearchRet{Exist: true, Password: true}
	var password string
	err = stmt.QueryRow(user.Username).Scan(&password, ret.RegistDate, ret.Role, ret.Avatar)
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
