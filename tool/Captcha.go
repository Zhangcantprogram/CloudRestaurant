package tool

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

type CaptchaResult struct {
	Id          string `json:"id"`
	Base64Blob  string `json:"base_64_blob"`
	VerifyValue string `json:"code"`
}

// 生成图形化验证码
func GenerateCaptcha(ctx *gin.Context) {
	//图形验证码的默认配置
	parameters := base64Captcha.ConfigCharacter{
		Height:             30,
		Width:              60,
		Mode:               3,
		ComplexOfNoiseText: 0,
		ComplexOfNoiseDot:  0,
		IsUseSimpleFont:    true,
		IsShowHollowLine:   false,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         4,
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 254,
		},
	}
	captchaId, captchaInterfaceInstance := base64Captcha.GenerateCaptcha("", parameters)
	base64blob := base64Captcha.CaptchaWriteToBase64Encoding(captchaInterfaceInstance)

	captchaResult := CaptchaResult{Id: captchaId, Base64Blob: base64blob}

	//设置json相应
	Success(ctx, map[string]interface{}{
		"captcha_result": captchaResult,
	})

}

// 验证验证码是否正确
func VerifyCaptcha(id string, value string) bool {
	verifyResult := base64Captcha.VerifyCaptcha(id, value)
	return verifyResult
}
