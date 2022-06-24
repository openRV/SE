// @Title DeleteItem.go
// @Description 关于 DeleteItem 功能的 API 的参数、返回结果数据类型
// @Author 杜沛然 ${DATE} ${TIME}

package desktop

// 删除 相当于移动到回收站
// /user/deleleitem
// "DELETE"

type DeleteItemParams struct {
	DocsId string `json:"docsId"`
	IsDir  string `json:"isDir"`
}

type DeleteItemResult struct {
	Success bool `json:"success"`
}

type DeleteItem struct {
	Params DeleteItemParams
	Result DeleteItemResult
}
