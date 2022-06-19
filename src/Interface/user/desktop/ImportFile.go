package desktop

// 导入文档
// /user/importfile
// "POST"

type ImportFileParams struct {
	//FormData FormData `json:"FormData"` // TODO 文件以formData上传
}

type ImportFileResult struct {
	Success bool `json:"success"`
	Data    struct {
		DocsId   string `json:"docsId"`
		DocsName string `json:"docsName"`
	} `json:"data"`
}

type ImportFile struct {
	Params ImportFileParams
	Result ImportFileResult
}
