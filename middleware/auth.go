package middleware

import (
	"Project/model"
	"Project/vo"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// CurrentUser 获取登录用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		uid := session.Get("user_id") // user_id 一一对应用户  一一对应  mysql数据
		if uid != nil {
			user, err := model.GetUser(uid)
			if err == nil {
				c.Set("user", &user)
			}
		}
		c.Next()
	}
}

// AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if _, ok := user.(*model.TMember); ok {
				c.Next()
				return
			}
		}

		c.JSON(200, vo.ResponseMeta{Code: vo.LoginRequired})
		c.Abort()
	}
}
