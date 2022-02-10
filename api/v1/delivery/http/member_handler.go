package http

import (
	"github.com/gin-gonic/gin"
	"github.com/junminhong/member-center-service/domain"
	"github.com/junminhong/member-center-service/pkg/requester"
	"github.com/junminhong/member-center-service/pkg/responser"
	"net/http"
	"time"
)

type MemberHandler struct {
	MemberUseCase    domain.MemberUseCase
	MemberRepository domain.MemberRepository
}

func NewMemberHandler(router *gin.Engine, memberUseCase domain.MemberUseCase, memberRepository domain.MemberRepository) {
	handler := &MemberHandler{
		MemberUseCase:    memberUseCase,
		MemberRepository: memberRepository,
	}
	router.POST("/api/v1/member", handler.Register)
	router.POST("/api/v1/member/login", handler.Login)
	router.POST("/api/v1/member/resend-email", handler.ResendEmail)
	router.POST("/api/v1/member/reset-password", handler.ResetPassword)
	router.POST("/api/v1/member/forget-password", handler.ForgetPassword)
	router.GET("api/v1/member/profile", handler.GetProfile)
	router.POST("api/v1/member/profile", handler.EditProfile)
	router.POST("api/v1/member/upload-mug-shot", handler.UploadMugShot)
}

// Register
// @Summary 註冊會員帳號
// @Description
// @Tags member
// @version 1.0
// @Accept application/json
// @produce application/json
// @param data body requester.Register true "請求資料"
// @Success 1004 {object} responser.Response "帳戶註冊成功"
// @failure 1000 {object} responser.Response "request格式錯誤"
// @failure 1002 {object} responser.Response "信箱已經存在"
// @failure 1003 {object} responser.Response "帳戶註冊失敗"
// @Router /member [post]
func (memberHandler *MemberHandler) Register(c *gin.Context) {
	request := requester.Register{}
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, responser.Response{
			ResultCode: responser.ReqBindErr.Code(),
			Message:    responser.ReqBindErr.Message(),
			Data:       "",
			TimeStamp:  time.Now(),
		})
		return
	}
	response := memberHandler.MemberUseCase.Register(request)
	c.JSON(http.StatusOK, response)
}

// Login
// @Summary 登入會員帳號
// @Description
// @Tags member
// @version 1.0
// @Accept application/json
// @produce application/json
// @param data body requester.Login true "請求資料"
// @Success 1007 {object} responser.Response "登入成功"
// @failure 1000 {object} responser.Response "request格式錯誤"
// @failure 1005 {object} responser.Response "信箱不存在"
// @failure 1006 {object} responser.Response "密碼輸入錯誤"
// @failure 1012 {object} responser.Response "該會員信箱未認證"
// @Router /member/login [post]
func (memberHandler *MemberHandler) Login(c *gin.Context) {
	request := requester.Login{}
	// 因為BindJOSN沒有return任何東西，所以要傳pointer進去給他修改，如果你單純傳值是不會影響原本的request
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, responser.Response{
			ResultCode: responser.ReqBindErr.Code(),
			Message:    responser.ReqBindErr.Message(),
			Data:       "",
			TimeStamp:  time.Now(),
		})
		return
	}
	response := memberHandler.MemberUseCase.Login(request)
	c.JSON(http.StatusOK, response)
}

// ResendEmail
// @Summary 重新發送驗證信
// @Description
// @Tags member
// @version 1.0
// @Accept application/json
// @produce application/json
// @param data body requester.ResendEmail true "請求資料"
// @Success 1008 {object} responser.Response "驗證信已重新發送"
// @failure 1000 {object} responser.Response "request格式錯誤"
// @failure 1005 {object} responser.Response "信箱不存在"
// @Router /member/resend-email [post]
func (memberHandler *MemberHandler) ResendEmail(c *gin.Context) {
	request := requester.ResendEmail{}
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, responser.Response{
			ResultCode: responser.ReqBindErr.Code(),
			Message:    responser.ReqBindErr.Message(),
			Data:       "",
			TimeStamp:  time.Now(),
		})
		return
	}
	response := memberHandler.MemberUseCase.ResendEmail(request)
	c.JSON(http.StatusOK, response)
}

// ResetPassword
// @Summary 重置密碼
// @Description
// @Tags member
// @version 1.0
// @Accept application/json
// @produce application/json
// @Param Authorization header string true "Atomic Token" default(Bearer <請在這邊輸入Atomic Token>)
// @param data body requester.ResetPassword true "請求資料"
// @Success 1009 {object} responser.Response "重置密碼成功"
// @failure 1000 {object} responser.Response "request格式錯誤"
// @failure 1011 {object} responser.Response "你沒有權限發起該請求"
// @failure 1010 {object} responser.Response "Atomic Token已經過期"
// @failure 1006 {object} responser.Response "密碼輸入錯誤"
// @Router /member/reset-password [post]
func (memberHandler *MemberHandler) ResetPassword(c *gin.Context) {
	request := requester.ResetPassword{}
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, responser.Response{
			ResultCode: responser.ReqBindErr.Code(),
			Message:    responser.ReqBindErr.Message(),
			Data:       "",
			TimeStamp:  time.Now(),
		})
		return
	}
	atomicToken := requester.GetAtomicToken(c)
	if atomicToken == "" {
		c.JSON(http.StatusOK, responser.Response{
			ResultCode: responser.NotAtomicTokenErr.Code(),
			Message:    responser.NotAtomicTokenErr.Message(),
			Data:       "",
			TimeStamp:  time.Now(),
		})
		return
	}
	response := memberHandler.MemberUseCase.ResetPassword(request, atomicToken)
	c.JSON(http.StatusOK, response)
}

// ForgetPassword
// @Summary 忘記密碼
// @Description
// @Tags member
// @version 1.0
// @Accept application/json
// @produce application/json
// @param data body requester.ForgetPassword true "請求資料"
// @Success 1008 {object} responser.Response "已寄送忘記密碼信"
// @failure 1000 {object} responser.Response "request格式錯誤"
// @failure 1005 {object} responser.Response "信箱不存在"
// @Router /member/forget-password [post]
func (memberHandler *MemberHandler) ForgetPassword(c *gin.Context) {
	request := requester.ForgetPassword{}
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, responser.Response{
			ResultCode: responser.ReqBindErr.Code(),
			Message:    responser.ReqBindErr.Message(),
			Data:       "",
			TimeStamp:  time.Now(),
		})
		return
	}
	response := memberHandler.MemberUseCase.ForgetPassword(request)
	c.JSON(http.StatusOK, response)
}

// GetProfile
// @Summary 取得會員Profile
// @Description
// @Tags member
// @version 1.0
// @Accept application/json
// @produce application/json
// @Param Authorization header string true "Atomic Token" default(Bearer <請在這邊輸入Atomic Token>)
// @Success 1015 {object} responser.Response "成功取得會員Profile資料"
// @failure 1011 {object} responser.Response "你沒有權限發起該請求"
// @failure 1005 {object} responser.Response "信箱不存在"
// @Router /member/profile [get]
func (memberHandler *MemberHandler) GetProfile(c *gin.Context) {
	atomicToken := requester.GetAtomicToken(c)
	if atomicToken == "" {
		c.JSON(http.StatusOK, responser.Response{
			ResultCode: responser.NotAtomicTokenErr.Code(),
			Message:    responser.NotAtomicTokenErr.Message(),
			Data:       "",
			TimeStamp:  time.Now(),
		})
		return
	}
	response := memberHandler.MemberUseCase.GetProfile(atomicToken)
	c.JSON(http.StatusOK, response)
}

// EditProfile
// @Summary 編輯會員Profile
// @Description
// @Tags member
// @version 1.0
// @Accept application/json
// @produce application/json
// @Param Authorization header string true "Atomic Token" default(Bearer <請在這邊輸入Atomic Token>)
// @param data body requester.EditProfile true "請求資料"
// @Success 1016 {object} responser.Response "成功編輯會員Profile資料"
// @failure 1000 {object} responser.Response "request格式錯誤"
// @failure 1011 {object} responser.Response "你沒有權限發起該請求"
// @failure 1005 {object} responser.Response "信箱不存在"
// @Router /member/profile [post]
func (memberHandler *MemberHandler) EditProfile(c *gin.Context) {
	request := requester.EditProfile{}
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, responser.Response{
			ResultCode: responser.ReqBindErr.Code(),
			Message:    responser.ReqBindErr.Message(),
			Data:       "",
			TimeStamp:  time.Now(),
		})
		return
	}
	atomicToken := requester.GetAtomicToken(c)
	if atomicToken == "" {
		c.JSON(http.StatusOK, responser.Response{
			ResultCode: responser.NotAtomicTokenErr.Code(),
			Message:    responser.NotAtomicTokenErr.Message(),
			Data:       "",
			TimeStamp:  time.Now(),
		})
		return
	}
	response := memberHandler.MemberUseCase.EditProfile(atomicToken, request)
	c.JSON(http.StatusOK, response)
}

// UploadMugShot
// @Summary 上傳大頭照
// @Description
// @Tags member
// @version 1.0
// @Accept application/json
// @produce multipart/form-data
// @Param Authorization header string true "Atomic Token" default(Bearer <請在這邊輸入Atomic Token>)
// @Param mug_shot formData file true "大頭照"
// @Success 1020 {object} responser.Response "大頭照上傳成功"
// @failure 1000 {object} responser.Response "request格式錯誤"
// @failure 1011 {object} responser.Response "你沒有權限發起該請求"
// @failure 1005 {object} responser.Response "信箱不存在"
// @failure 1017 {object} responser.Response "上傳檔案超過10MB"
// @failure 1018 {object} responser.Response "僅接受圖片檔案格式"
// @failure 1018 {object} responser.Response "大頭照上傳失敗"
// @Router /member/upload-mug-shot [post]
func (memberHandler *MemberHandler) UploadMugShot(c *gin.Context) {
	file, uploadFile, err := c.Request.FormFile("mug_shot")
	if err != nil {
		c.JSON(http.StatusOK, responser.Response{
			ResultCode: responser.ReqBindErr.Code(),
			Message:    responser.ReqBindErr.Message(),
			Data:       "",
			TimeStamp:  time.Now(),
		})
		return
	}
	atomicToken := requester.GetAtomicToken(c)
	if atomicToken == "" {
		c.JSON(http.StatusOK, responser.Response{
			ResultCode: responser.NotAtomicTokenErr.Code(),
			Message:    responser.NotAtomicTokenErr.Message(),
			Data:       "",
			TimeStamp:  time.Now(),
		})
		return
	}
	response := memberHandler.MemberUseCase.UploadMugShot(atomicToken, file, uploadFile)
	c.JSON(http.StatusOK, response)
}
