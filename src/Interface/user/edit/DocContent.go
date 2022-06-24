// @Title DocContent.go
// @Description 关于 DocContent 功能的 API 的参数、返回结果数据类型
// @Author 杜沛然 ${DATE} ${TIME}

package edit

// /user/readcontent
// "GET"

type DocContentParams struct {
	DocId string `json:"docsId"`
}

type DocData struct {
	DocContent string `json:"docsContent"`
}

type DocContentResult struct {
	Success bool    `json:"success"`
	Data    DocData `json:"data"`
}

type DocContent struct {
	Params DocContentParams
	Result DocContentResult
}
