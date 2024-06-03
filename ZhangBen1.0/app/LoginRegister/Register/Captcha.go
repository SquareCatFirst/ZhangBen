package Register

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"net/http"
)

/*
	var digitDriver = base64Captcha.DriverDigit{
		Height: 50 ,  // 生成图片高度
		Width: 150,  // 生成图片宽度
		Length: 5,  // 验证码长度
		MaxSkew: 1,  // 文字的倾斜度 越大倾斜越狠 ， 越不容易看懂
		DotCount: 1,  // 背景的点数，越大，字体越模糊
	}
*/
type CaptchaReply struct {
	CaptchaID string `json:"captchaid"`
	ImageData string `json:"imagedata"`
}

var (
	CapchaStore base64Captcha.Store
	Drivers     *base64Captcha.DriverString
)

func InitCatcha() {
	CapchaStore = base64Captcha.DefaultMemStore
	Drivers = base64Captcha.NewDriverString(
		50, 110, 6, 1, 4,
		"123456789abcdefghjkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ", // 包含数字和字母的字符集
		nil, nil, nil)
}

func Captcha(c *gin.Context) {
	InitCatcha()

	driver := Drivers
	captcha := base64Captcha.NewCaptcha(driver, CapchaStore)
	id, b64s, answer, err := captcha.Generate()
	fmt.Println("####################", answer, " ", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate captcha"})
		return
	}
	c.JSON(http.StatusOK,
		CaptchaReply{
			CaptchaID: id,
			ImageData: b64s,
		})
}

type VerifyRequest struct {
	CaptchaID     string `json:"captchaid"`
	CaptchaAnswer string `json:"captchaanswer"`
}

func Verify(c *gin.Context) {
	in := new(VerifyRequest)
	if err := c.ShouldBindJSON(in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	captchaId := in.CaptchaID
	captchaSolution := in.CaptchaAnswer

	if captchaId == "" || captchaSolution == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "captcha_id and captcha_solution are required"})
		return
	}

	driver := base64Captcha.NewDriverString(
		80, 240, 6, 1, 4,
		"123456789abcdefghjkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ", // 包含数字和字母的字符集
		nil, nil, nil)
	captcha := base64Captcha.NewCaptcha(driver, CapchaStore)

	if captcha.Verify(captchaId, captchaSolution, true) {
		c.JSON(http.StatusOK, gin.H{"message": "验证码正确"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "验证码错误"})
	}
}
