package requester

type AuthEmailToken struct {
	EmailToken string `json:"atomic_token" bind:"required"`
}
