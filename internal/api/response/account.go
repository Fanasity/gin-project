package response

type Account struct {
	ID       int    `json:"id" example:"1"`
	Nickname string `json:"nickname" example:"fan"`
	Email    string `json:"email" example:"liufanwh@126.com"`
}

type AccountAuth struct {
	Account
	Token string `json:"token" example:"asdeERASXAq324AQW"`
}
