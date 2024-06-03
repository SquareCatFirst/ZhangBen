package Register

import (
	"ZhangBen1.0/DB"
	UT "ZhangBen1.0/UserType"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"net/http"
)

type RegUser struct {
	Username        string `json:"username"`
	Nickname        string `json:"nickname"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmpassword"`
	CaptchaID       string `json:"captchaid"`
	CaptchaAnswer   string `json:"captchaanswer"`
}
type uid struct {
	Id  int `json:"id"`
	Cnt int `json:"cnt"`
}

func Register(c *gin.Context) {
	const (
		emailRegexPattern    = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
		passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&+-/])[A-Za-z\d$@$!%*#?&+-/]{8,}$`
	)
	var u RegUser
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messgae": "无效的请求数据"})
		return
	}
	var finduser UT.User
	if err := DB.Db.Model(&finduser).Where("username = ?", u.Username).Select(); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"messgae": "用户名重复"})
		return
	}

	{
		emailexp := regexp.MustCompile(emailRegexPattern, regexp.None)
		isemail, err := emailexp.MatchString(u.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"messgae": "系统错误"})
			return
		}
		if !isemail {
			c.JSON(http.StatusBadRequest, gin.H{"messgae": "邮箱格式错误"})
			return
		}

		pwdexp := regexp.MustCompile(passwordRegexPattern, regexp.None)
		ispwd, err := pwdexp.MatchString(u.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"messgae": "系统错误"})
			return
		}
		if !ispwd {
			c.JSON(http.StatusBadRequest, gin.H{"messgae": "密码格式错误(密码必须不少于八位，必须要包含一个数字、一个字母、一个特殊字符($@$!%*#?&+-/))"})
			return
		}
	}

	if u.Password != u.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"messgae": "两次输入的密码不一致"})
		return
	}

	{
		captchaId := u.CaptchaID
		captchaSolution := u.CaptchaAnswer

		if captchaId == "" || captchaSolution == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "captcha_id and captcha_solution are required"})
			return
		}

		driver := Drivers
		captcha := base64Captcha.NewCaptcha(driver, CapchaStore)

		if captcha.Verify(captchaId, captchaSolution, true) {
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "验证码错误"})
			return
		}
	}

	{

		var t uid
		err := DB.Db.Model(&t).Select()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"messgae": "分配uid异常"})
			return
		}
		user := UT.User{
			Uid:      t.Cnt,
			Username: u.Username,
			Nickname: u.Nickname,
			Email:    u.Email,
			Password: u.Password,
			Phone:    u.Phone,
		}
		_, err = DB.Db.Model(&user).Insert()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"messgae": "注册失败"})
		} else {
			upd := uid{
				Id:  1,
				Cnt: t.Cnt + 1,
			}
			_, err := DB.Db.Model(&upd).WherePK().Update()
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"messgae": "数据更新失败"})
			}
			c.JSON(http.StatusOK, gin.H{"messgae": "注册成功"})

		}
	}

}
