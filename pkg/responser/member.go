package responser

type GetProfile struct {
	NickName    string `json:"nick_name"`
	MugShotPath string `json:"mug_shot_path"`
}
type Login struct {
	AtomicToken        string `json:"atomic_token"`
	RefreshAtomicToken string `json:"refresh_atomic_token"`
}
