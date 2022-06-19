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
