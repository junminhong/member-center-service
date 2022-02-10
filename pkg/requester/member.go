package requester

type Register struct {
	Email    string `json:"email" bind:"required"`
	Password string `json:"password" bind:"required"`
	NickName string `json:"nick_name" bind:"required"`
}

type Login struct {
	Email    string `json:"email" bind:"required;email"`
	Password string `json:"password" bind:"required"`
}

type ResendEmail struct {
	Email string `json:"email" bind:"required;email"`
}

type ResetPassword struct {
	OldPassword string `json:"old_password" bind:"required"`
	NewPassword string `json:"new_password" bind:"required"`
}

type ForgetPassword struct {
	Email string `json:"email" bind:"required"`
}

type EditProfile struct {
	NickName string `json:"nick_name"`
}
