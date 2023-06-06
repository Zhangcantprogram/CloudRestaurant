package param

// 手机号+验证码 登录时的传参
type SmsLoginParam struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}
