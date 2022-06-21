package database

import "fmt"

type DocSearchInfo struct {
	Type    string // Title | Author
	Content string
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
