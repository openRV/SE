package index

// 用户运行状况
// /admin/userinfo

type UserInfoParams struct {
}

type UserInfoResult struct {
	Success bool `json:"success"`
	Data    struct {
		UserNumbers   int      `json:"userNumbers"`
		OnlineNumbers int      `json:"onlineNumbers"`
		OnlineUsers   []string `json:"onlineUsers"`
		TotalUsers    []string `json:"totalUsers"`
		// 按月新增
		// 按天新增
	} `json:"data"`
}

type UserInfo struct {
	Params UserInfoParams
	Result UserInfoResult
}
