package routers

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/xiaoriri-team/gin-skeleton/internal/middleware"
	"github.com/xiaoriri-team/gin-skeleton/internal/routers/api"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.HandleMethodNotAllowed = true
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 跨域配置
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")
	r.Use(cors.New(corsConfig))

	// 获取version
	r.GET("/", api.Version)

	// 用户登录
	r.POST("/auth/login", api.Login)

	// 用户注册
	r.POST("/auth/register", api.Register)

	// 获取验证码
	r.GET("/captcha", api.GetCaptcha)

	// 无鉴权路由组
	noAuthApi := r.Group("/")
	{
		// 获取用户基本信息
		noAuthApi.GET("/user/profile", api.GetUserProfile)
	}

	// 鉴权路由组
	authApi := r.Group("/").Use(middleware.JWT())
	privApi := r.Group("/").Use(middleware.JWT()).Use(middleware.Priv())
	{

		// 获取当前用户信息
		authApi.GET("/user/info", api.GetUserInfo)

		// 修改密码
		authApi.POST("/user/password", api.ChangeUserPassword)

		// 修改昵称
		authApi.POST("/user/nickname", api.ChangeNickname)

		// 修改头像
		authApi.POST("/user/avatar", api.ChangeAvatar)

		// 上传资源
		privApi.POST("/attachment", api.UploadAttachment)

	}
	// 默认404
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "Not Found",
		})
	})
	// 默认405
	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"code": 405,
			"msg":  "Method Not Allowed",
		})
	})
	return r
}
