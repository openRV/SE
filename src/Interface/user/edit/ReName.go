package edit

// 更改名称
// /user/rename
// "POST"

type ReNameParams struct {
	DocsId  string `json:"docsId"`
	NewName string `json:"newName"`
	IsDir   bool   `json:"isDir"`
}

type ReNameResult struct {
	Success bool `json:"success"`
}

type ReName struct {
	Params ReNameParams
	Result ReNameResult
}
