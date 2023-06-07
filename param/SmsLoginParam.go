package param

// 手机号+验证码 登录时的传参，从前端接收参数的结构体
type SmsLoginParam struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}
