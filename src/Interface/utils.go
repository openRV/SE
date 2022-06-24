// @Title utils.go
// @Description 关于统一错误返回、抽象接口的数据类型
// @Author 杜沛然 ${DATE} ${TIME}

package Interface

// 统一错误返回
type ErrorRes struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}

// 抽象接口 ， 未使用
type Request[T interface{}, K interface{} | ErrorRes] struct {
	Params T    `json:"params"`
	Result K    `json:"result"`
	Method int8 `json:"method"`
}
