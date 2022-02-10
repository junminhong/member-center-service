package http

import (
	"github.com/gin-gonic/gin"
	"github.com/junminhong/member-center-service/domain"
	"github.com/junminhong/member-center-service/pkg/requester"
	"github.com/junminhong/member-center-service/pkg/responser"
	"net/http"
	"time"
)

type authHandler struct {
	AuthUseCase domain.AuthUseCase
	AuthRepo    domain.AuthRepository
}

func NewAuthRepo(router *gin.Engine, authUseCase domain.AuthUseCase, authRepo domain.AuthRepository) {
	handler := &authHandler{authUseCase, authRepo}
	router.POST("/api/v1/auth/email", handler.AuthEmailToken)
	router.POST("/api/v1/auth/refresh-atomic-token", handler.RefreshAtomicToken)
}

// AuthEmailToken
// @Summary 驗證信箱的Token
// @Description
// @Tags auth
// @version 1.0
// @Accept application/json
// @produce application/json
// @Param email_token query string true "Email Token"
// @Success 1014 {object} responser.Response "信箱驗證成功"
// @failure 1000 {object} responser.Response "request格式錯誤"
// @failure 1013 {object} responser.Response "Email Token已經過期"
// @failure 1005 {object} responser.Response "信箱不存在"
// @Router /auth/email [post]
func (authHandler *authHandler) AuthEmailToken(c *gin.Context) {
	emailToken := c.Query("email_token")
	if emailToken == "" {
		c.JSON(http.StatusOK, responser.Response{
			ResultCode: responser.ReqBindErr.Code(),
			Message:    responser.ReqBindErr.Message(),
			Data:       "",
			TimeStamp:  time.Now(),
		})
		return
	}
	request := requester.AuthEmailToken{EmailToken: emailToken}
	response := authHandler.AuthUseCase.AuthEmailToken(request)
	c.JSON(http.StatusOK, response)
}

// RefreshAtomicToken
// @Summary 重新取得Atomic Token
// @Description
// @Tags auth
// @version 1.0
// @Accept application/json
// @produce application/json
// @Param Authorization header string true "Refresh Atomic Token" default(Bearer <請在這邊輸入Refresh Atomic Token>)
// @Success 1021 {object} responser.Response "已成功重新取得Atomic Token"
// @failure 1000 {object} responser.Response "request格式錯誤"
// @failure 1010 {object} responser.Response "Refresh Atomic Token已經過期"
// @failure 1011 {object} responser.Response "你沒有權限發起該請求"
// @Router /auth/refresh-atomic-token [post]
func (authHandler *authHandler) RefreshAtomicToken(c *gin.Context) {
	refreshAtomicToken := requester.GetAtomicToken(c)
	if refreshAtomicToken == "" {
		c.JSON(http.StatusOK, responser.Response{
			ResultCode: responser.ReqBindErr.Code(),
			Message:    responser.ReqBindErr.Message(),
			Data:       "",
			TimeStamp:  time.Now(),
		})
		return
	}
	response := authHandler.AuthUseCase.RefreshAtomicToken(refreshAtomicToken)
	c.JSON(http.StatusOK, response)
}
