package database

import (
	"fmt"
	"time"
)

type DocSearchInfo struct {
	Type    string // Title | Author
	Content string
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
}

type RenameRet struct {
	Success bool
	Msg     string
}

// returned string -> error msg, nil for success
func OpenSearch(search DocSearchInfo) ([]Doc, string) {

	var result []Doc

	if search.Type == "Author" {

		stmt := `select 
				docsId , docsName , author , createDate , lastUpdate , docsType , viewCounts 
				from 
				Doc where 
				author like ?
				AND open = ?
				`
		rows, err := DB.Query(stmt, fmt.Sprintf("%%%s%%", search.Content), true)
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

	} else {

		stmt := `select 
				docsId , docsName , author , createDate , lastUpdate , docsType , viewCounts 
				from Doc 
				where 
				docsName like ?
				AND open = ?
				`
		rows, err := DB.Query(stmt, fmt.Sprintf("%%%s%%", search.Content), true)
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

}

func HotOpenDocs() ([]Doc, string) {

	var result []Doc

	stmt := `select 
				docsId , docsName , author , createDate , lastUpdate , docsType , viewCounts 
				from Doc where open = ?
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
				(? , ? , ? , ? , ? , ? , ?)
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
				(? , ? , ? , ?)
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
	stmt, err := DB.Prepare("select docFile From Doc")
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
				set open = ?
				where docsId = ? AND author = ?
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
				set dirId = ? , root = ?
				where subId = ? AND subType = ? 
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
	stmt, err := DB.Prepare(`
				update 
				Doc
				set docsName = ?
				where docsId = ? AND author = ? 
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

	return RenameRet{Success: true}
}
