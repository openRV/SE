// @Title DownloadLog.go
// @Description 关于 DownloadLog 功能的 API 的参数、返回结果数据类型
// @Author 杜沛然 ${DATE} ${TIME}

package desktop

// 下载记录
// /user/downloadlog
// "GET"

type DownloadLogParams struct {
	CurPage  int `json:"curPage"`
	PageSize int `json:"pageSize"`
}

type DownloadLogResult struct {
	Success bool `json:"success"`
	Total   int  `json:"total"`
	Data    []struct {
		ListItem     DocsListItem `json:"item"` // TODO 在docsTtem的基础上加上一个下载日期
		DownloadData string       `json:"date"` // TODO
	} `json:"data"`
}

type DownloadLog struct {
	Params DownloadLogParams
	Result DownloadLogResult
}
