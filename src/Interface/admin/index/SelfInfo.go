// @Title SelfInfo.go
// @Description 关于 SelfInfo 功能的 API 的参数、返回结果数据类型
// @Author 杜沛然 ${DATE} ${TIME}

package index

// 查询自己的信息
// /admin/selfinfo
// "GET"

type SelfInfoParams struct {
}

type SelfInfoResult struct {
	Success bool `json:"success"`
	//Total   int  `json:"total"`
	Data struct {
		UserName string `json:"userName"`
		Password string `json:"password"`
		Avatar   string `json:"avatar"`
	}
}

type SelfInfo struct {
	Params SelfInfoParams
	Result SelfInfoResult
}
