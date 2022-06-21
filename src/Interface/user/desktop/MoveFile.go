package desktop

// 移动
// /user/movefile
// "POST"

type MoveFileParams struct {
	DocsId string `json:"docsId"` //
	MoveTo string `json:"moveTo"` // a dir
}

type MoveFileResult struct {
	Success bool `json:"success"`
}

type MoveFile struct {
	Params MoveFileParams
	Result MoveFileResult
}
