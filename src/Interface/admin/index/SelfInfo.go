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
