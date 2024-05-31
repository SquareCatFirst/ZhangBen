package LoginRegister

import (
	"ZhangBen1.0/DB"
	UT "ZhangBen1.0/UserType"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegUser struct {
	Username        string `json:"username"`
	Nickname        string `json:"nickname"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmpassword"`
}
type uid struct {
	Id  int `json:"id"`
	Cnt int `json:"cnt"`
}

func Register(c *gin.Context) {
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

	if u.Password != u.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"messgae": "两次输入的密码不一致"})
		return
	} else {

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
