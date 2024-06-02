package Register

import "github.com/mojocn/base64Captcha"

var digitDriver = base64Captcha.DriverDigit{
	Height:   50,  // 生成图片高度
	Width:    150, // 生成图片宽度
	Length:   5,   // 验证码长度
	MaxSkew:  1,   // 文字的倾斜度 越大倾斜越狠 ， 越不容易看懂
	DotCount: 1,   // 背景的点数，越大，字体越模糊
}

var store = base64Captcha.DefaultMemStore

/*
func CaptchaFenerate() (base64Captcha.Captcha, error) {
	var ret base64Captcha.Captcha

	c := base64Captcha.NewCaptcha(&digitDriver, store) //生成验证码
	id, b64s, _, err := c.Generate()
	if err != nil {
		return ret, err
	}
	ret.Verify(id, "1145", true)

}
*/
