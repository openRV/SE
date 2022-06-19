package edit

// 导出为文件下载
// /user/downloadfile
// "GET"

type DownloadFileParams struct {
	DocsId string `json:"docsId"`
}

type DownloadFileResult struct {
	Success bool `json:"success"`
}

type DownloadFile struct {
	Params DownloadFileParams
	Result DownloadFileResult
}
