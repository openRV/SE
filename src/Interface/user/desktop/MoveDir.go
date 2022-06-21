package desktop

// 移动
// /user/movedir
// "POST"

type MoveDirParams struct {
	DirId  string `json:"dirId"`  //
	MoveTo string `json:"moveTo"` // a dir
}

type MoveDirResult struct {
	Success bool `json:"success"`
}

type MoveDir struct {
	Params MoveFileParams
	Result MoveFileResult
}
