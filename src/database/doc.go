package database

import (
	"fmt"
	"time"
)

type OpenDocSearchInfo struct {
	Title  string
	Author string
}

type SetVizInfoRet struct {
	Success bool
	Msg     string
}

type SetVisInfo struct {
	Id       string
	Username string
	Vis      string
}

type Doc struct {
	DocsId     string
	DocsName   string
	DocsFile   []byte
	Author     string
	CreateDate string
	LastUpdate string
	DocsType   string
	ViewCounts int
}

type NewDocInfo struct {
	FatherDirId string
	DocName     string
	Username    string
}

type NewDocRes struct {
	Success bool
	Msg     string
	Id      string
}
type MoveDocInfo struct {
	Id       string
	MoveTo   string
	Username string
}

type MoveDocRet struct {
	Success bool
	Msg     string
}

type RenameInfo struct {
	Id       string
	Newname  string
	Username string
	IsDir    bool
}

type RenameRet struct {
	Success bool
	Msg     string
}

type LastViewInfo struct {
	Username string
}

type LastViewRes struct {
	Success bool
	Msg     string
	Data    []Doc
}

type UserSearchInfo struct {
	SearchContent string
	SearchType    string // Title Author
	Username      string
}

type UserSearchRes struct {
	Success bool
	Msg     string
	Data    []Doc
}

type DocsContentInfo struct {
	Username string
	Id       string
}

type DocsContentRes struct {
	Success bool
	Msg     string
	Data    string
}

type WriteDocsInfo struct {
	Id       string
	Username string
	Content  string
}

type WriteDocsRes struct {
	Success bool
	Msg     string
}

// returned string -> error msg, nil for success
func OpenSearch(search OpenDocSearchInfo) ([]Doc, string) {

	var result []Doc
	var flag int
	// Author
	if search.Author == "" && search.Title != "" {
		flag = 0
	}
	// Title
	if search.Author != "" && search.Title == "" {
		flag = 1
	}
	// Author & Title
	if search.Author != "" && search.Title != "" {
		flag = 2
	}
	switch flag {
	case 1:
		{
			stmt := `select 
					docsId , docsName , author , createDate , lastUpdate , docsType , viewCounts 
					from 
					Doc where 
					author like $1
					AND open = $2
					`
			rows, err := DB.Query(stmt, fmt.Sprintf("%%%s%%", search.Author), true)
			if err != nil {
				fmt.Println(err)
				return nil, "database error"
			}
			defer rows.Close()
			for rows.Next() {
				var doc Doc
				rows.Scan(&doc.DocsId, &doc.DocsName, &doc.Author, &doc.CreateDate, &doc.LastUpdate, &doc.DocsType, &doc.ViewCounts)
				result = append(result, doc)
			}

			// return result, ""
		}
	case 0:
		{
			stmt := `select 
					docsId , docsName , author , createDate , lastUpdate , docsType , viewCounts 
					from Doc 
					where 
					docsName like $1
					AND open = $2
					`
			rows, err := DB.Query(stmt, fmt.Sprintf("%%%s%%", search.Title), true)
			if err != nil {
				fmt.Println(err)
				return nil, "database error"
			}
			defer rows.Close()
			for rows.Next() {
				var doc Doc
				rows.Scan(&doc.DocsId, &doc.DocsName, &doc.Author, &doc.CreateDate, &doc.LastUpdate, &doc.DocsType, &doc.ViewCounts)
				result = append(result, doc)
			}

			// return result, ""
		}
	case 2:
		{
			stmt := `select 
					docsId , docsName , author , createDate , lastUpdate , docsType , viewCounts 
					from Doc 
					where 
					docsName like $1
					AND author like $2
					AND open = $3
					`
			rows, err := DB.Query(stmt, fmt.Sprintf("%%%s%%", search.Title), fmt.Sprintf("%%%s%%", search.Author), true)
			if err != nil {
				fmt.Println(err)
				return nil, "database error"
			}
			defer rows.Close()
			for rows.Next() {
				var doc Doc
				rows.Scan(&doc.DocsId, &doc.DocsName, &doc.Author, &doc.CreateDate, &doc.LastUpdate, &doc.DocsType, &doc.ViewCounts)
				result = append(result, doc)
			}
		}

	}
	return result, ""
}

func HotOpenDocs() ([]Doc, string) {

	var result []Doc

	stmt := `select 
				docsId , docsName , author , createDate , lastUpdate , docsType , viewCounts 
				from Doc where open = $1
				order by viewCounts desc
				`
	rows, err := DB.Query(stmt, true)
	if err != nil {
		fmt.Println(err)
		return nil, "database error"
	}
	defer rows.Close()

	for rows.Next() {
		var doc Doc
		rows.Scan(&doc.DocsId, &doc.DocsName, &doc.Author, &doc.CreateDate, &doc.LastUpdate, &doc.DocsType, &doc.ViewCounts)
		result = append(result, doc)
	}

	return result, ""

}

func NewDoc(info NewDocInfo) NewDocRes {

	// insert into Dir
	docId := uniqString()
	isRoot := info.FatherDirId == info.Username

	stmt, err := DB.Prepare(`
				insert into
				Doc (docsId ,docsName , author , createDate , lastUpdate , viewCounts , open)
				values
				($1 , $2 , $3 , $4 , $5 , $6 , $7)
			`)
	if err != nil {
		fmt.Println(err)
		return NewDocRes{Success: false, Msg: "database error"}
	}
	defer stmt.Close()

	_, err = stmt.Exec(docId, info.DocName, info.Username, time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"), 0, false)

	if err != nil {
		fmt.Println(err)
		return NewDocRes{Success: false, Msg: "database error"}
	}

	// insert into Tree
	stmt, err = DB.Prepare(`
				insert into
				Tree (dirId , root , subType , subId)
				values
				($1 , $2 , $3 , $4)
			`)
	if err != nil {
		fmt.Println(err)
		return NewDocRes{Success: false, Msg: "database error"}
	}

	_, err = stmt.Exec(info.FatherDirId, isRoot, "doc", docId)
	if err != nil {
		fmt.Println(err)
		return NewDocRes{Success: false, Msg: "database error"}
	}

	return NewDocRes{Success: true, Id: docId}
}

func GetAllDocSize() int {
	var ret int
	stmt, err := DB.Prepare("select docsFile From Doc")
	if err != nil {
		fmt.Println(err)
		return -1
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		fmt.Println(err)
		return -1
	}

	for rows.Next() {
		var file []byte
		rows.Scan(&file)
		ret += len(file)
	}

	return ret
}

func SetVisibility(info SetVisInfo) SetVizInfoRet {

	stmt, err := DB.Prepare(`
				update 
				Doc
				set open = $1
				where docsId = $2 AND author = $3
			`)
	if err != nil {
		fmt.Println(err)
		return SetVizInfoRet{Success: false, Msg: "database error"}
	}
	defer stmt.Close()

	_, err = stmt.Exec(info.Vis == "public", info.Id, info.Username)

	if err != nil {
		fmt.Println(err)
		return SetVizInfoRet{Success: false, Msg: "database error"}
	}

	return SetVizInfoRet{Success: true}

}

func MoveDoc(info MoveDocInfo) MoveDocRet {
	stmt, err := DB.Prepare(`
				update 
				Tree
				set dirId = $1 , root = $2
				where subId = $3 AND subType = $4
			`)
	if err != nil {
		fmt.Println(err)
		return MoveDocRet{Success: false, Msg: "database error"}
	}
	defer stmt.Close()

	_, err = stmt.Exec(info.MoveTo, info.MoveTo == info.Username, info.Id, "doc")

	if err != nil {
		fmt.Println(err)
		return MoveDocRet{Success: false, Msg: "database error"}
	}

	return MoveDocRet{Success: true}

}

func Rename(info RenameInfo) RenameRet {
	if !info.IsDir {
		// rename file
		stmt, err := DB.Prepare(`
				update 
				Doc
				set docsName = $1
				where docsId = $2 AND author = $3 
			`)
		if err != nil {
			fmt.Println(err)
			return RenameRet{Success: false, Msg: "database error"}
		}
		defer stmt.Close()

		_, err = stmt.Exec(info.Newname, info.Id, info.Username)

		if err != nil {
			fmt.Println(err)
			return RenameRet{Success: false, Msg: "database error"}
		}
	} else {
		// rename dir
		stmt, err := DB.Prepare(`
				update 
				Dir
				set dirName = $1
				where dirId = $2 AND owner = $3 
			`)
		if err != nil {
			fmt.Println(err)
			return RenameRet{Success: false, Msg: "database error"}
		}
		defer stmt.Close()

		_, err = stmt.Exec(info.Newname, info.Id, info.Username)

		if err != nil {
			fmt.Println(err)
			return RenameRet{Success: false, Msg: "database error"}
		}
	}

	return RenameRet{Success: true}
}

func LastView(info LastViewInfo) LastViewRes {
	stmt, err := DB.Prepare(`
				select docsId , docsName , createDate , viewCounts , lastUpdate
				from Doc 
				where author = $1
				order by lastUpdate desc
			`)
	if err != nil {
		fmt.Println(err)
		return LastViewRes{Success: false, Msg: "database error"}
	}
	defer stmt.Close()

	rows, err := stmt.Query(info.Username)

	if err != nil {
		fmt.Println(err)
		return LastViewRes{Success: false, Msg: "database error"}
	}
	var result LastViewRes
	for rows.Next() {
		var doc Doc
		err = rows.Scan(&doc.DocsId, &doc.DocsName, &doc.CreateDate, &doc.ViewCounts, &doc.LastUpdate)
		if err != nil {
			fmt.Println(err)
			return LastViewRes{Success: false, Msg: "database error"}

		}
		doc.Author = info.Username
		result.Data = append(result.Data, doc)
	}

	result.Success = true
	return result

}

func UserSearch(info UserSearchInfo) UserSearchRes {

	var result UserSearchRes

	if info.SearchType == "Author" {

		stmt, err := DB.Prepare(`
					select docsId , docsName , createDate , lastUpdate , viewCounts
					from Doc
					where author = $1 AND author like $2
				`)
		if err != nil {
			fmt.Println(err)
			return UserSearchRes{Success: false, Msg: "database error"}
		}

		rows, err := stmt.Query(info.Username, fmt.Sprintf("%%%s%%", info.SearchContent))
		if err != nil {
			fmt.Println(err)
			return UserSearchRes{Success: false, Msg: "database error"}
		}

		for rows.Next() {
			var doc Doc
			doc.Author = info.Username
			err = rows.Scan(&doc.DocsId, &doc.DocsName, &doc.CreateDate, &doc.LastUpdate, &doc.ViewCounts)
			if err != nil {
				fmt.Println(err)
				return UserSearchRes{Success: false, Msg: "database error"}
			}
			result.Data = append(result.Data, doc)
		}
	} else {
		stmt, err := DB.Prepare(`
					select docsId , docsName , createDate , lastUpdate , viewCounts
					from Doc
					where author = $1 AND docsName like $2
				`)
		if err != nil {
			fmt.Println(err)
			return UserSearchRes{Success: false, Msg: "database error"}
		}

		rows, err := stmt.Query(info.Username, fmt.Sprintf("%%%s%%", info.SearchContent))
		if err != nil {
			fmt.Println(err)
			return UserSearchRes{Success: false, Msg: "database error"}
		}
		for rows.Next() {
			var doc Doc
			doc.Author = info.Username
			err = rows.Scan(&doc.DocsId, &doc.DocsName, &doc.CreateDate, &doc.LastUpdate, &doc.ViewCounts)
			if err != nil {
				fmt.Println(err)
				return UserSearchRes{Success: false, Msg: "database error"}
			}
			result.Data = append(result.Data, doc)
		}
	}

	result.Success = true
	return result

}

func DocsContent(info DocsContentInfo) DocsContentRes {
	stmt, err := DB.Prepare("select docsFile from Doc where docsId = $1 AND author = $2")
	if err != nil {
		fmt.Println(err)
		return DocsContentRes{Success: false, Msg: "database error"}
	}
	defer stmt.Close()

	var data []byte
	err = stmt.QueryRow(info.Id, info.Username).Scan(&data)
	if err != nil {
		fmt.Println(err)
		return DocsContentRes{Success: false, Msg: "database error"}
	}

	return DocsContentRes{Success: true, Data: string(data)}
}

func WriteDocs(info WriteDocsInfo) WriteDocsRes {
	stmt, err := DB.Prepare("update Doc set docsFile = $1 , lastUpdate = $2 where docsId = $3 AND author = $4")
	if err != nil {
		fmt.Println(err)
		return WriteDocsRes{Success: false, Msg: "database error"}
	}

	_, err = stmt.Exec([]byte(info.Content), time.Now().Format("2006-01-12 12:13:14"), info.Id, info.Username)
	if err != nil {
		fmt.Println(err)
		return WriteDocsRes{Success: false, Msg: "database error"}
	}

	return WriteDocsRes{Success: true}
}

type DocsNameRet struct {
	Success  bool
	Msg      string
	DocsName string
}

func GetDocsName(docsId string) DocsNameRet {
	row := DB.QueryRow("select docsName from Doc where docsId=$1", docsId)

	var docsName string
	err := row.Scan(&docsName)
	if err != nil {
		fmt.Println(err)
		return DocsNameRet{Success: false, Msg: "database error"}
	}

	return DocsNameRet{Success: true, DocsName: docsName}
}
