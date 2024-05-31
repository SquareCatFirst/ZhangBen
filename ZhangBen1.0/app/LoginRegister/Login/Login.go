package Login

import (
	DB "ZhangBen1.0/DB"
	UT "ZhangBen1.0/UserType"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"net/http"
	"strconv"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GetLoginUsertype struct {
	Uid      int    `json:"uid"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetLoginUser(db *pg.DB, username string) (R GetLoginUsertype, er error) {
	var user UT.User
	err := db.Model(&user).Where("username = ?", username).Select()
	if err != nil {
		return GetLoginUsertype{
			Uid:      0,
			Username: "",
			Password: "",
		}, err
	}
	return GetLoginUsertype{
		Uid:      user.Uid,
		Username: username,
		Password: user.Password,
	}, nil

}

func Login(c *gin.Context) {
	var loginReq LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messgae": "无效的请求数据"})
		return
	}
	uslog, err := GetLoginUser(DB.Db, loginReq.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messgae": "查询用户数据失败"})
		fmt.Println("查询用户数据失败")
		return
	}
	//	ok := false

	if uslog.Password == loginReq.Password {
		//	c.JSON(http.StatusOK, gin.H{"messgae": "登录成功"})
		fmt.Println("登录成功")

		c.SetCookie("uid", strconv.Itoa(uslog.Uid), 3600, "/", "", false, false)

		//	c.HTML(200, "index.html", gin.H{})
		//
		c.JSON(http.StatusOK, gin.H{})
		//	c.Redirect(http.StatusMovedPermanently, "/test/index.html")
		//	ok = true

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"messgae": "用户名或密码错误"})
		fmt.Println("用户名或密码错误")
	}

}
