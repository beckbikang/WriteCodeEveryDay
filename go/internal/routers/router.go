package routers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/beckbikang/WriteCodeEveryDay/docs"
	"github.com/beckbikang/WriteCodeEveryDay/pkg/auth"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	if gin.Mode() == gin.DebugMode {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(gin.Recovery())
	}

	// 健康检查路由
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Swagger文档路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API路由组
	api := r.Group("/api/v1")
	{
		// 账号相关路由
		api.POST("/passport", CreatePassport)
		api.PUT("/passport/:id", UpdatePassport)
		api.PUT("/passport/:id/block", BlockPassport)
		api.GET("/passport/exist/:user_name", ExistPassport)
		api.POST("/passport/login", Login)
		api.GET("/passport/verify", func(c *gin.Context) {
			token := c.GetHeader("token")
			if token == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "token is empty"})
				return
			}

			// 验证JWT Token
			claims, err := auth.ParseToken(token)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				return
			}

			// 验证通过，返回成功响应
			c.JSON(http.StatusOK, claims)
		})
	}

	return r
}
