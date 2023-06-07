package param

// 在进行用户名和密码登陆时，用于接收前端传来的信息
type LoginParam struct {
	Name     string `json:"name"`  //用户名
	Password string `json:"pwd"`   //密码
	Id       string `json:"id"`    //验证码id
	Value    string `json:"value"` //验证码的输入值
}
