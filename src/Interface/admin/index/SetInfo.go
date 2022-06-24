// @Title SetInfo.go
// @Description 关于 SelfInfo 功能的 API 的参数、返回结果数据类型
// @Author 杜沛然 ${DATE} ${TIME}
package index

// 设置自己的信息
// /admin/setinfo
// "POST"

type SetInfoParams struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}

type SetInfoResult struct {
	Success bool `json:"success"`
}
