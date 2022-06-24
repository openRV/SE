// @Title EmptyTrash.go
// @Description 关于 EmptyTrash 功能的 API 的参数、返回结果数据类型
// @Author 杜沛然 ${DATE} ${TIME}

package desktop

// 清空回收站
// /user/emptytrash
// "DELETE"

type EmptyTrashParams struct {
}

type EmptyTrashResult struct {
	Success bool `json:"success"`
}

type EmptyTrash struct {
	Params EmptyTrashParams
	Result EmptyTrashResult
}
