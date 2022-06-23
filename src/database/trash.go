package database

import (
	"fmt"
	"time"
)

type DeleteItemInfo struct {
	Username string
	Id       string
	IsDir    bool
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

type TrashInfo struct {
	Username string
}

type TrashData struct {
	Name       string
	Id         string
	Author     string
	DeleteDate string
}

type TrashRes struct {
	Success bool
	Msg     string
	Data    []TrashData
}

type Item struct {
	ItemType string
	Id       string
}

func DeleteItem(info DeleteItemInfo) DeleteItemRes {

	// get item's type doc or dir
	subType := "doc"
	if info.IsDir {
		subType = "dir"
	}

	if subType == "doc" {
		// insert into trash
		stmt, err := DB.Prepare("insert into Trash(itemType , itemId , owner , deleteDate) values (?,?,?,?)")
		if err != nil {
			fmt.Println(err)
			return DeleteItemRes{Success: false, Msg: "database error"}
		}
		_, err = stmt.Exec(subType, info.Id, info.Username, time.Now().Format("2006-01-02 15:04:05"))
		if err != nil {
			fmt.Println(err)
			return DeleteItemRes{Success: false, Msg: "database error"}
		}

		// delete from share
		stmt, err = DB.Prepare("delete from Share where docId = ?")
		if err != nil {
			fmt.Println(err)
			return DeleteItemRes{Success: false, Msg: "database error"}
		}
		_, err = stmt.Exec(info.Id)
		if err != nil {
			fmt.Println(err)
			return DeleteItemRes{Success: false, Msg: "database error"}
		}

	}

	// delete from tree
	stmt, err := DB.Prepare("delete from Tree where subType = ? AND subId = ?")
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

	var item []Item

	for rows.Next() {
		var itemType string
		var itemId string

		err = rows.Scan(&itemType, &itemId)
		if err != nil {
			fmt.Println(err)
			return EmptyTrashRes{Success: false, Msg: "database error"}
		}

		item = append(item, Item{ItemType: itemType, Id: itemId})
	}

	for i := range item {
		if item[i].ItemType == "doc" {
			stmt1, err := DB.Prepare("delete from Doc where docsId = ? AND author = ?")
			if err != nil {
				fmt.Println(err)
				return EmptyTrashRes{Success: false, Msg: "database error"}
			}
			_, err = stmt1.Exec(item[i].Id, info.Username)
			if err != nil {
				fmt.Println(err)
				return EmptyTrashRes{Success: false, Msg: "database error"}
			}

			stmt1, err = DB.Prepare("delete from Trash where itemId = ? AND owner = ?")
			if err != nil {
				fmt.Println(err)
				return EmptyTrashRes{Success: false, Msg: "database error"}
			}
			_, err = stmt1.Exec(item[i].Id, info.Username)
			if err != nil {
				fmt.Println(err)
				return EmptyTrashRes{Success: false, Msg: "database error"}
			}

		}
	}

	return EmptyTrashRes{Success: true}

}

func Trash(info TrashInfo) TrashRes {
	stmt, err := DB.Prepare("select itemId , deleteDate from Trash where itemType == ? AND owner == ?")
	if err != nil {
		fmt.Println(err)
		return TrashRes{Success: false, Msg: "database error"}
	}
	defer stmt.Close()
	rows, err := stmt.Query("doc", info.Username)
	if err != nil {
		fmt.Println(err)
		return TrashRes{Success: false, Msg: "database error"}
	}
	defer rows.Close()

	var result TrashRes
	for rows.Next() {
		var data TrashData
		data.Author = info.Username
		err = rows.Scan(&data.Id, &data.DeleteDate)
		if err != nil {
			fmt.Println(err)
			return TrashRes{Success: false, Msg: "database error"}
		}
		result.Data = append(result.Data, data)
	}
	result.Success = true
	return result
}
