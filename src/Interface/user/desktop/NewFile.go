package desktop

// 新建文件
// "POST"

type NewFileParams struct {
	FatherDirId string `json:"fatherDirId"` // 默认根文件夹
	DocsName    string `json:"docsName"`
}

type NewFileResult struct {
	Success bool `json:"success"`
	Data    struct {
		DocsId string `json:"docsId"`
	} `json:"data"`
}

type NewFile struct {
	Params NewFileParams
	Result NewFileResult
}
