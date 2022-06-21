package database

import (
	"fmt"
	"time"
)

type DeleteItemInfo struct {
	Username string
	Id       string
}

type DeleteItemRes struct {
	Success bool
	Msg     string
}

type EmptyTrashInfo struct {
	Username string
}

type EmptyTrashRes struct {
	Success bool
	Msg     string
}

func DeleteItem(info DeleteItemInfo) DeleteItemRes {

	// get item's type doc or dir
	stmt, err := DB.Prepare("select subType from Tree where subId = ?")
	if err != nil {
		fmt.Println(err)
		return DeleteItemRes{Success: false, Msg: "database error"}
	}
	defer stmt.Close()

	var subType string
	stmt.QueryRow(info.Id).Scan(&subType)

	// insert into trash
	stmt, err = DB.Prepare("insert into Trash(itemType , itemId , owner , deleteDate) values (?,?,?,?)")
	if err != nil {
		fmt.Println(err)
		return DeleteItemRes{Success: false, Msg: "database error"}
	}
	_, err = stmt.Exec(subType, info.Id, info.Username, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		fmt.Println(err)
		return DeleteItemRes{Success: false, Msg: "database error"}
	}

	// delete from tree
	stmt, err = DB.Prepare("delete from Tree where subType = ? AND subId = ?")
	if err != nil {
		fmt.Println(err)
		return DeleteItemRes{Success: false, Msg: "database error"}
	}
	_, err = stmt.Exec(subType, info.Id)
	if err != nil {
		fmt.Println(err)
		return DeleteItemRes{Success: false, Msg: "database error"}
	}

	return DeleteItemRes{Success: true}
}

func EmptyTrash(info EmptyTrashInfo) EmptyTrashRes {

	stmt, err := DB.Prepare("select itemType , itemId from Trash where owner = ?")
	if err != nil {
		fmt.Println(err)
		return EmptyTrashRes{Success: false, Msg: "database error"}
	}
	rows, err := stmt.Query(info.Username)
	if err != nil {
		fmt.Println(err)
		return EmptyTrashRes{Success: false, Msg: "database error"}
	}
	for rows.Next() {
		var itemType string
		var itemId string

		err = rows.Scan(&itemType, &itemId)
		if err != nil {
			fmt.Println(err)
			return EmptyTrashRes{Success: false, Msg: "database error"}
		}
		if itemType == "doc" {
			stmt, err = DB.Prepare("delete from Doc where docsId = ? AND author = ?")
			if err != nil {
				fmt.Println(err)
				return EmptyTrashRes{Success: false, Msg: "database error"}
			}
			_, err = stmt.Exec(itemId, info.Username)
			if err != nil {
				fmt.Println(err)
				return EmptyTrashRes{Success: false, Msg: "database error"}
			}

		} else {
			stmt, err = DB.Prepare("delete from Dir where dirId = ? AND author = ?")
			if err != nil {
				fmt.Println(err)
				return EmptyTrashRes{Success: false, Msg: "database error"}
			}
			_, err = stmt.Exec(itemId, info.Username)
			if err != nil {
				fmt.Println(err)
				return EmptyTrashRes{Success: false, Msg: "database error"}
			}

		}

	}

	return EmptyTrashRes{Success: true}

}
