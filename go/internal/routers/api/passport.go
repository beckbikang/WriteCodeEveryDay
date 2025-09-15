package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/beckbikang/WriteCodeEveryDay/internal/service"
	"github.com/beckbikang/WriteCodeEveryDay/pkg/app"
	"github.com/beckbikang/WriteCodeEveryDay/pkg/auth"
)

// @Summary 创建账号
// @Produce  json
// @Param user_name body string true "用户名" minlength(2) maxlength(50)
// @Param password body string true "密码" minlength(6) maxlength(50)
// @Param email body string true "邮箱" maxlength(100)
// @Param user_info body string false "用户信息" maxlength(1024)
// @Param head_img body string false "头像" maxlength(1024)
// @Success 200 {string} string "成功"
// @Failure 400 {object} app.ErrorResponse "请求错误"
// @Failure 500 {object} app.ErrorResponse "内部错误"
// @Router /api/v1/passport [post]
func CreatePassport(c *gin.Context) {
	param := service.CreatePassportRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		response.ToErrorResponse(errs)
		return
	}

	svc := service.New(c.Request.Context(), global.DBEngine)
	err := svc.CreatePassport(&param)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}

	response.ToResponse(gin.H{})
}

// @Summary 更新账号
// @Produce  json
// @Param id path int true "账号ID"
// @Param user_name body string false "用户名" minlength(2) maxlength(50)
// @Param email body string false "邮箱" maxlength(100)
// @Param user_info body string false "用户信息" maxlength(1024)
// @Param head_img body string false "头像" maxlength(1024)
// @Success 200 {string} string "成功"
// @Failure 400 {object} app.ErrorResponse "请求错误"
// @Failure 500 {object} app.ErrorResponse "内部错误"
// @Router /api/v1/passport/{id} [put]
func UpdatePassport(c *gin.Context) {
	param := service.UpdatePassportRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		response.ToErrorResponse(errs)
		return
	}

	svc := service.New(c.Request.Context(), global.DBEngine)
	err := svc.UpdatePassport(&param)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}

	response.ToResponse(gin.H{})
}

// @Summary 拉黑账号
// @Produce  json
// @Param id path int true "账号ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} app.ErrorResponse "请求错误"
// @Failure 500 {object} app.ErrorResponse "内部错误"
// @Router /api/v1/passport/{id}/block [put]
func BlockPassport(c *gin.Context) {
	param := service.BlockPassportRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		response.ToErrorResponse(errs)
		return
	}

	svc := service.New(c.Request.Context(), global.DBEngine)
	err := svc.BlockPassport(&param)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}

	response.ToResponse(gin.H{})
}

// @Summary 判断账号是否存在
// @Produce  json
// @Param user_name path string true "用户名"
// @Success 200 {object} gin.H "成功"
// @Failure 500 {object} app.ErrorResponse "内部错误"
// @Router /api/v1/passport/exist/{user_name} [get]
func ExistPassport(c *gin.Context) {
	userName := c.Param("user_name")
	response := app.NewResponse(c)

	svc := service.New(c.Request.Context(), global.DBEngine)
	exist, err := svc.ExistPassport(userName)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}

	response.ToResponse(gin.H{"exist": exist})
}

// @Summary 登录账号
// @Produce json
// @Param user_name body string true "用户名"
// @Param password body string true "密码"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/passport/login [post]
func Login(c *gin.Context) {
	// 参数绑定和验证
	param := &service.LoginRequest{}
	if valid, errs := app.BindAndValid(c, param); !valid {
		app.ResponseError(c, errs)
		return
	}

	// 调用服务层方法
	svc := service.New(c.Request.Context(), nil)
	success, err := svc.Login(param)
	if err != nil {
		app.ResponseError(c, err)
		return
	}

	// 生成JWT Token
	if success {
		token, err := svc.GenerateToken(param.UserName)
		if err != nil {
			app.ResponseError(c, err)
			return
		}
		app.ResponseSuccess(c, map[string]string{"token": token})
	} else {
		app.ResponseSuccess(c, success)
	}
}

// @Summary 验证JWT Token
// @Produce json
// @Param token header string true "JWT Token"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/passport/verify [get]
func Verify(c *gin.Context) {
	token := c.GetHeader("token")
	if token == "" {
		app.ResponseError(c, errors.New("token is empty"))
		return
	}

	// 验证JWT Token
	claims, err := auth.ParseToken(token)
	if err != nil {
		app.ResponseError(c, err)
		return
	}

	// 验证通过，返回成功响应
	app.ResponseSuccess(c, claims)
}
