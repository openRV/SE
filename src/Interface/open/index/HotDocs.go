// @Title HotDocs.go
// @Description 关于 HotDocs 功能的 API 的参数、返回结果数据类型
// @Author 杜沛然 ${DATE} ${TIME}

package index

// /open/hotdocs
// "GET"

type HotdocsParams struct {
}

type HotdocsResult struct {
	Success bool          `json:"success"`
	Total   int           `json:"total"`
	Data    []HotdocsData `json:"data"`
}

type HotdocsData struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	DownloadNum int    `json:"downloadNum"`
}

type Hotdocs struct {
	Params HotdocsParams
	Result HotdocsResult
}
