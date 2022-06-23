// @Title dir.go
// @Desctiption 操作数据库中 Dir 表相关的函数及相关的数据类型
// @Author 杜沛然 ${DATE} ${TIME}

package database

import (
	"SE/src/Interface/user/desktop"
	"fmt"
	"math/rand"
	"time"
)

// Dir 存储目录信息，对应数据库的 Dir 表
type Dir struct {
	Id         string // 目录 id (键)
	Name       string // 目录名
	Owner      string // 目录所有者(创建人)
	CreateDate string // 目录创建日期
	LastView   string // 目录最后一期被修改的日期
	Subdir     []Dir  // 子目录，递归表示
}

// UserDirRet 查询目录统一的返回结果
type UserDirRet struct {
	Success bool   // 查询是否成功
	Name    string // 用户根目录名
	Data    []Dir  // 用户根目录下的目录， 用 Dir 递归表示
	Msg     string // 若查询失败，填写封装后的错误信息
}

// NewDirInto 新建目录的信息
type NewDirInfo struct {
	FatherDirId string // 新建目录的上一级目录 id
	Name        string // 新建目录名
	Owner       string // 新建目录的所有者
}

// NewDirRes 新建目录的返回结果
type NewDirRes struct {
	Success bool   // 新建目录是否成功
	Msg     string // 如果不成功，填写封装后的错误信息
}

// MoveDirInfo 移动目录所需信息
type MoveDirInfo struct {
	Id       string // 被移动的目录的 id
	MoveTo   string // 将要移动到的目录的 id
	Username string // 操作者的用户名
}

// MoveDirRet 移动目录的返回结果
type MoveDirRet struct {
	Success bool   // 移动目录是否成功
	Msg     string // 如果不成功， 填写封装后的错误信息
}

// DirContentInfo 查询目录下内容的信息
type DirContentInfo struct {
	Id       string // 被查询的目录 id
	Username string // 查询者的用户名
}

// DirContentRes 查询目录内容的返回结果
type DirContentRes struct {
	Success bool                   // 查询是否成功
	Msg     string                 // 如果不成功， 填写封装后的错误信息
	Data    desktop.DirContentData // 递归表示目录下的内容
}

// ImportFileInfo 导入文件所需的信息
type ImportFileInfo struct {
	DirId    string // 导入文件到的文件夹id
	File     []byte // 导入文件的二进制流
	Username string // 导入者的用户名
	FileName string // 新导入的文件名
}

// ImportFileRes 导入文件的返回结果
type ImportFileRes struct {
	Success bool   // 导入是否成功
	Msg     string // 如果不成功，填写封装后的错误信息
	Id      string // 新导入的文件 id
	Name    string // 新导入的文件名
}

// @title UserDir
// @description 用于查询数据库以获得指定用户的文件夹树
// @auth 杜沛然 ${DATE} ${TIME}
// @param  id    string      "要查询的用户id，也是被查询的用户根目录的id"
// @param  root  bool        "查询的目录是否是用户的根目录"
// @return  _    UserDirRet  "包含了用户文件夹树的结构体"
func UserDir(id string, root bool) UserDirRet {

	// 初始化返回的结构体
	var result UserDirRet
	result.Name = id
	result.Success = true

	// 从 Tree 表中查找子目录的 id
	stmt, err := DB.Prepare(`
				select subId
				from Tree
				where dirId = $1 AND root = $2 AND subType = $3
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

	// result 的 Data 字段设置为查询到的子目录 id 数组
	result.Data = append(result.Data, Dir{Id: id, Name: id})
	for rows.Next() {
		var dir Dir
		err = rows.Scan(&dir.Id)
		if err != nil {
			fmt.Println(err)
			return UserDirRet{Success: false, Msg: "database error"}
		}

		// 填写 dir 中的 name, owner, createDate, lastView 字段
		fillinDirInfo(&dir)
		// 填写 dir 中的 subdir
		fillinSubDir(&dir)

		result.Data[0].Subdir = append(result.Data[0].Subdir, dir)
	}

	return result

}

// @title fillinSubDir
// @description 深度优先递归的填写子目录 id 和信息
// @auth 杜沛然 ${DATE} ${TIME}
// @param   dir  *Dir        "要被填写子目录信息的目录指针"
// @result  _    UserDirRet  "用于判断递归终止的结构体，仅用到 Success 字段"
func fillinSubDir(dir *Dir) UserDirRet {
	// 查询 Tree 寻找父目录 id 为 dirId 的字文件夹
	stmt, err := DB.Prepare(`
				select subId
				from Tree
				where dirId = $1 AND subType = $2
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
		// 当 row 为空时递归终止
		var subDir Dir
		row.Scan(&subDir.Id)
		// 递归填写子目录 id
		res := fillinSubDir(&subDir)
		if !res.Success {
			return res
		}
		// 填写子目录 info
		res = fillinDirInfo(&subDir)
		if !res.Success {
			return res
		}

		dir.Subdir = append(dir.Subdir, subDir)
	}
	return UserDirRet{Success: true}

}

// @title fillinDirInfo
// @description 填写子目录信息
// @auth 杜沛然 ${DATE} ${TIME}
// @param   dir  *Dir        "要被填写子目录信息的目录指针"
// @result  _    UserDirRet  "用于封装数据库的错误并返回"
func fillinDirInfo(dir *Dir) UserDirRet {
	stmt, err := DB.Prepare(`
				select dirName , owner , createDate , lastView
				from Dir
				where dirId = $1
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

// @title NewDir
// @description 新建目录
// @auth 杜沛然 ${DATE} ${TIME}
// @param   info  NewDirInfo  "新建目录有关的信息"
// @result  _     NewDirRes   "新建文件夹的返回信息"
func NewDir(info NewDirInfo) NewDirRes {

	// 新文件夹的 id
	dirId := uniqString()
	isRoot := info.FatherDirId == info.Owner

	// 在 Dir 中插入新文件夹的信息
	stmt, err := DB.Prepare(`
				insert into 
				Dir (dirId , dirName , owner , createDate , lastView) 
				values
				($1 , $2 , $3 , $4 , $5)
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

	// 在 Tree 中插入目录所在的位置
	stmt, err = DB.Prepare(`
				insert into 
				Tree (dirId , root , subType , subId)
				values
				($1 , $2 , $3 , $4)
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

// @title uniqString
// @description 生成全局唯一的键
// @auth 杜沛然 ${DATE} ${TIME}
// @result  _  string  "生成的键，10位，每位 0-9 A-Z"
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

// @title MoveDir
// @description 用于移动文件夹
// @auth 杜沛然 ${DATE} ${TIME}
// @param info MoveDirInfo "移动文件夹所需的信息"
// @result  _  MoveDirRet  "移动文件夹的封装后的返回结果"
func MoveDir(info MoveDirInfo) MoveDirRet {

	// 在 Tree 中将父目录 id 修改为新的被移动到的目录的 id
	stmt, err := DB.Prepare(`
				update 
				Tree
				set dirId = $1 , root = $2
				where subId = $3 AND subType = $4 
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

// @title DirContent
// @description 用于获取文件夹下的内容
// @auth 杜沛然 ${DATE} ${TIME}
// @param info DirContentInfo "获取文件所需的信息"
// @result  _  DirContentRes  "获取文件的封装后的返回结果"
func DirContent(info DirContentInfo) DirContentRes {
	// 在 Tree 表中查询当前目录下的子目录及文件的类型和 id
	stmt, err := DB.Prepare("select subType , subId from Tree where dirId = $1")
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
	var dirArray []Dir
	var docArray []Doc
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
			// 对于子文件夹，查询 Dir 获取目录信息
			dir.Id = subId
			dir.Owner = info.Username
			stmt1, err := DB.Prepare("select dirName , createDate , lastView from Dir where dirId = $1 AND owner = $2")
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
			dirArray = append(dirArray, dir)
		} else {
			// 对于子文件，查询 Doc 获取文件信息
			doc.DocsId = subId
			doc.Author = info.Username
			stmt1, err := DB.Prepare("select docsName , createDate , lastUpdate from Doc where docsId = $1 AND author = $2")
			if err != nil {
				fmt.Println(err)
				return DirContentRes{Success: false, Msg: "database error"}
			}
			defer stmt1.Close()
			err = stmt1.QueryRow(doc.DocsId, doc.Author).Scan(&doc.DocsName, &doc.CreateDate, &doc.LastUpdate)
			if err != nil {
				fmt.Println(err)
				return DirContentRes{Success: false, Msg: "database error"}
			}
			docArray = append(docArray, doc)
		}

	}
	var data desktop.DirContentData

	for i := range dirArray {
		data.Dir = append(data.Dir, desktop.DirListItem{
			DirId:      dirArray[i].Id,
			DirName:    dirArray[i].Name,
			Owner:      dirArray[i].Owner,
			CreateDate: dirArray[i].CreateDate,
			LastView:   dirArray[i].LastView,
		})
	}

	for i := range docArray {
		data.Docs = append(data.Docs, desktop.DocsListItem{
			DocsId:     docArray[i].DocsId,
			DocsName:   docArray[i].DocsName,
			Author:     docArray[i].Author,
			CreateDate: docArray[i].CreateDate,
			LastView:   docArray[i].LastUpdate,
		})
	}

	result.Data = data
	result.Success = true

	return result

}

// @title ImportFile
// @description 导入文件
// @auth 杜沛然 ${DATE} ${TIME}
// @param info ImportFileInfo "导入文件所需的信息"
// @result  _  ImportFileInfo  "导入文件的封装后的返回结果"
func ImportFile(info ImportFileInfo) ImportFileRes {

	// 在 Doc 中插入新导入的文件
	stmt, err := DB.Prepare(`insert into Doc 
							(docsId , docsName , dodcsFile , author , createDate , lastUpdate , DocsType , viewCounts , open)
							values ($1,$2,$3,$4,$5,$6,$7,$8)
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

	// 在 Tree 中插入新文件在目录中的位置
	stmt, err = DB.Prepare(`
				insert into Tree(dirId , root , subType , subId) 
				values ($1,$2,$3,$4)
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
