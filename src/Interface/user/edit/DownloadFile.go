// @Title DownloadFile.go
// @Description 关于 DownloadFile 功能的 API 的参数、返回结果数据类型
// @Author 杜沛然 ${DATE} ${TIME}

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
