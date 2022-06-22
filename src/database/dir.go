package database

import (
	"SE/src/Interface/user/desktop"
	"fmt"
	"math/rand"
	"time"
)

type Dir struct {
	Id         string
	Name       string
	Owner      string
	CreateDate string
	LastView   string
	Subdir     []Dir
}

type UserDirRet struct {
	Success bool
	Name    string
	Data    []Dir
	Msg     string
}

type NewDirInfo struct {
	FatherDirId string
	Name        string
	Owner       string
}

type NewDirRes struct {
	Success bool
	Msg     string
}

type MoveDirInfo struct {
	Id       string
	MoveTo   string
	Username string
}

type MoveDirRet struct {
	Success bool
	Msg     string
}

type DirContentInfo struct {
	Id       string
	Username string
}

type DirContentRes struct {
	Success bool
	Msg     string
	Data    []desktop.DirContentData
}

type ImportFileInfo struct {
	DirId    string
	File     []byte
	Username string
	FileName string
}

type ImportFileRes struct {
	Success bool
	Msg     string
	Id      string
	Name    string
}

func UserDir(id string, root bool) UserDirRet {

	var result UserDirRet

	// fill in dir name

	//stmt, err := DB.Prepare(`
	//			select dirName
	//			from Dir
	//			where dirId = ?
	//		`)
	//if err != nil {
	//	fmt.Println(err)
	//	return UserDirRet{Success: false, Msg: "database error"}
	//}
	//defer stmt.Close()
	//err = stmt.QueryRow(id).Scan(&result.Name)
	//if err != nil {
	//	fmt.Println(err)
	//	return UserDirRet{Success: false, Msg: "database error"}
	//}

	result.Name = id

	// find sub dir Id from table Tree
	stmt, err := DB.Prepare(`
				select subId
				from Tree
				where dirId = ? AND root = ? AND subType = ?
			`)
	if err != nil {
		fmt.Println(err)
		return UserDirRet{Success: false, Msg: "database error"}
	}
	rows, err := stmt.Query(id, root, "dir")
	if err != nil {
		fmt.Println(err)
		return UserDirRet{Success: false, Msg: "database error"}
	}
	defer rows.Close()

	result.Data = append(result.Data, Dir{Id: id, Name: id})
	for rows.Next() {
		var dir Dir
		err = rows.Scan(&dir.Id)
		if err != nil {
			fmt.Println(err)
			return UserDirRet{Success: false, Msg: "database error"}
		}

		// fill in blank name, owner, createDate, lastView of result
		fillinDirInfo(&dir)
		// fill in subdir
		fillinSubDir(&dir)

		result.Data[0].Subdir = append(result.Data[0].Subdir, dir)
	}

	result.Success = true
	return result

}

func fillinSubDir(dir *Dir) UserDirRet {
	stmt, err := DB.Prepare(`
				select subId
				from Tree
				where dirId = ? AND subType = ?
			`)
	if err != nil {
		fmt.Println(err)
		return UserDirRet{Success: false, Msg: "database error"}
	}
	row, err := stmt.Query(dir.Id, "dir")
	if err != nil {
		fmt.Println(err)
		return UserDirRet{Success: false, Msg: "database error"}
	}

	for row.Next() {
		var subDir Dir
		row.Scan(&subDir.Id)
		res := fillinSubDir(&subDir)
		if !res.Success {
			return res
		}
		res = fillinDirInfo(&subDir)
		if !res.Success {
			return res
		}

		dir.Subdir = append(dir.Subdir, subDir)
	}
	return UserDirRet{Success: true}

}

func fillinDirInfo(dir *Dir) UserDirRet {
	stmt, err := DB.Prepare(`
				select dirName , owner , createDate , lastView
				from Dir
				where dirId = ?
			`)
	if err != nil {
		fmt.Println(err)
		return UserDirRet{Success: false, Msg: "database error"}
	}
	err = stmt.QueryRow(dir.Id).Scan(&dir.Name, &dir.Owner, &dir.CreateDate, &dir.LastView)
	if err != nil {
		fmt.Println(err)
		return UserDirRet{Success: false, Msg: "database error"}
	}
	return UserDirRet{Success: true}
}

func NewDir(info NewDirInfo) NewDirRes {

	// insert into Dir
	dirId := uniqString()

	isRoot := info.FatherDirId == info.Owner

	stmt, err := DB.Prepare(`
				insert into 
				Dir (dirId , dirName , owner , createDate , lastView) 
				values
				(? , ? , ? , ? , ?)
			`)
	if err != nil {
		fmt.Println(err)
		return NewDirRes{Success: false, Msg: "database error"}
	}
	defer stmt.Close()

	_, err = stmt.Exec(dirId, info.Name, info.Owner, time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"))

	if err != nil {
		fmt.Println(err)
		return NewDirRes{Success: false, Msg: "database error"}
	}

	// insert into Tree
	stmt, err = DB.Prepare(`
				insert into 
				Tree (dirId , root , subType , subId)
				values
				(? , ? , ? , ?)
			`)
	if err != nil {
		fmt.Println(err)
		return NewDirRes{Success: false, Msg: "database error"}
	}

	_, err = stmt.Exec(info.FatherDirId, isRoot, "dir", dirId)
	if err != nil {
		fmt.Println(err)
		return NewDirRes{Success: false, Msg: "database error"}
	}

	return NewDirRes{Success: true}

}

func uniqString() string {
	rand.Seed(time.Now().UnixNano())
	var uniqId string
	for i := 0; i < 10; i++ {
		rand1 := rand.Int63n(2)
		var res int64
		if rand1 == 0 {
			res = 48 + rand.Int63n(10)
		} else {
			res = 97 + rand.Int63n(26)
		}
		character := fmt.Sprintf("%c", res)
		uniqId += character
	}
	return uniqId
}

func MoveDir(info MoveDirInfo) MoveDirRet {
	stmt, err := DB.Prepare(`
				update 
				Tree
				set dirId = ? , root = ?
				where subId = ? AND subType = ? 
			`)
	if err != nil {
		fmt.Println(err)
		return MoveDirRet{Success: false, Msg: "database error"}
	}
	defer stmt.Close()

	_, err = stmt.Exec(info.MoveTo, info.MoveTo == info.Username, info.Id, "dir")

	if err != nil {
		fmt.Println(err)
		return MoveDirRet{Success: false, Msg: "database error"}
	}

	return MoveDirRet{Success: true}

}

func DirContent(info DirContentInfo) DirContentRes {
	stmt, err := DB.Prepare("select subType , subId from Tree where dirId = ?")
	if err != nil {
		fmt.Println(err)
		return DirContentRes{Success: false, Msg: "database error"}
	}
	defer stmt.Close()

	rows, err := stmt.Query(info.Id)
	if err != nil {
		fmt.Println(err)
		return DirContentRes{Success: false, Msg: "database error"}
	}

	var result DirContentRes
	for rows.Next() {
		var dir Dir
		var doc Doc
		var subType string
		var subId string
		err = rows.Scan(&subType, &subId)
		if err != nil {
			fmt.Println(err)
			return DirContentRes{Success: false, Msg: "database error"}
		}

		if subType == "dir" {
			dir.Id = subId
			dir.Owner = info.Username
			stmt1, err := DB.Prepare("select dirName , createDate , lastView from Dir where dirId = ? AND owner = ?")
			if err != nil {
				fmt.Println(err)
				return DirContentRes{Success: false, Msg: "database error"}
			}
			defer stmt1.Close()
			err = stmt1.QueryRow(dir.Id, dir.Owner).Scan(&dir.Name, &dir.CreateDate, &dir.LastView)
			if err != nil {
				fmt.Println(err)
				return DirContentRes{Success: false, Msg: "database error"}
			}
		} else {
			doc.DocsId = subId
			doc.Author = info.Username
			stmt1, err := DB.Prepare("select docsName , docsType , createDate , lastUpdate from Dir where docsId = ? AND author = ?")
			if err != nil {
				fmt.Println(err)
				return DirContentRes{Success: false, Msg: "database error"}
			}
			defer stmt1.Close()
			err = stmt1.QueryRow(doc.DocsId, doc.Author).Scan(&doc.DocsName, &doc.DocsType, &doc.CreateDate, &doc.LastUpdate)
			if err != nil {
				fmt.Println(err)
				return DirContentRes{Success: false, Msg: "database error"}
			}
		}

		var data desktop.DirContentData
		data.Dir.CreateDate = dir.CreateDate
		data.Dir.DirId = dir.Id
		data.Dir.DirName = dir.Name
		data.Dir.LastView = dir.LastView
		data.Dir.Owner = dir.Owner

		data.Docs.Author = doc.Author
		data.Docs.CreateDate = doc.CreateDate
		data.Docs.DocsId = doc.DocsId
		data.Docs.DocsName = doc.DocsName
		data.Docs.DocsType = doc.DocsType
		data.Docs.LastView = doc.LastUpdate

		result.Data = append(result.Data, data)
	}
	return result

}

func ImportFile(info ImportFileInfo) ImportFileRes {
	stmt, err := DB.Prepare(`insert into Doc 
							(docsId , docsName , dodcsFile , author , createDate , lastUpdate , DocsType , viewCounts , open)
							values (?,?,?,?,?,?,?,?)
							`)
	if err != nil {
		fmt.Println(err)
		return ImportFileRes{Success: false, Msg: "database error"}
	}
	defer stmt.Close()

	id := uniqString()
	_, err = stmt.Exec(id, info.FileName, info.File, info.Username, time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"), "", 0, false)
	if err != nil {
		fmt.Println(err)
		return ImportFileRes{Success: false, Msg: "database error"}
	}

	stmt, err = DB.Prepare(`
				insert into Tree(dirId , root , subType , subId) 
				values (?,?,?,?)
							`)
	if err != nil {
		fmt.Println(err)
		return ImportFileRes{Success: false, Msg: "database error"}
	}
	_, err = stmt.Exec(info.DirId, info.DirId == info.Username, "doc", id)
	if err != nil {
		fmt.Println(err)
		return ImportFileRes{Success: false, Msg: "database error"}
	}

	return ImportFileRes{Success: true, Id: id, Name: info.FileName}
}
