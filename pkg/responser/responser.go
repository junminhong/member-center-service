package responser

import (
	"time"
)

var (
	Ok                    = add(0, "ok")
	ErrRequest            = add(400, "")
	ErrNotFind            = add(404, "")
	ErrForbidden          = add(403, "")
	ErrNoPermission       = add(405, "")
	ErrServer             = add(500, "")
	ReqBindErr            = add(1000, "請依照API文件進行請求")
	EmailExistsErr        = add(1002, "信箱已存在")
	MemberRegisterErr     = add(1003, "帳戶註冊失敗")
	MemberRegisterOk      = add(1004, "帳戶註冊成功")
	MemberNotFoundErr     = add(1005, "信箱不存在")
	PasswordNotMatchErr   = add(1006, "密碼輸入錯誤")
	LoginOk               = add(1007, "登入成功")
	ResendEmailOk         = add(1008, "已重新寄送驗證信")
	ResetPasswordOk       = add(1009, "重置密碼成功")
	AtomicTokenExpiredErr = add(1010, "Atomic Token已經過期")
	NotAtomicTokenErr     = add(1011, "你沒有權限發起該請求")
	EmailNotAuthErr       = add(1012, "該會員帳戶信箱未認證")
	EmailTokenExpiredErr  = add(1013, "Email token已經過期")
	EmailAuthOk           = add(1014, "Email驗證完成")
	GetProfileOk          = add(1015, "取得會員資料成功")
	EditProfileOk         = add(1016, "修改會員資料成功")
	UploadFileTooBigErr   = add(1017, "上傳檔案容量超過限制")
	UploadFileNotImgErr   = add(1018, "僅接受上傳檔案格式為圖片")
	UploadFileErr         = add(1019, "檔案上傳失敗")
	UploadFileOk          = add(1020, "檔案上傳成功")
	RefreshAtomicTokenOk  = add(1021, "重新取得Atomic Token成功")
)

func New(code int, msg string) ResponseFlag {
	if code < 1000 {
		panic("error code must be greater than 1000")
	}
	return add(code, msg)
}

func add(code int, msg string) ResponseFlag {
	return ResponseFlag{
		code: code, message: msg,
	}
}

func (responseFlag *ResponseFlag) Error() string {
	return responseFlag.message
}

func (responseFlag ResponseFlag) Message() string {
	return responseFlag.message
}

func (responseFlag ResponseFlag) Reload(message string) ResponseFlag {
	responseFlag.message = message
	return responseFlag
}

func (responseFlag ResponseFlag) Code() int {
	return responseFlag.code
}

type Response struct {
	ResultCode int         `json:"result_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	TimeStamp  time.Time   `json:"time_stamp"`
}

type ResponseFlag struct {
	code    int
	message string
}

type ResponseFunc interface {
	Error() string
	Code() int
	Message() string
	Reload(string) ResponseFlag
}
