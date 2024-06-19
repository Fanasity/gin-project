package request

type UserEmail struct {
	Email string `example:"liufanwh@126.com"`
}

type UserLogin struct {
	Email    string `example:"liufanwh@126.com"`
	Password string `example:"4dfg54df4g8"`
}

type UserRegister struct {
	Email    string `example:"liufanwh@126.com"`
	Nickname string `example:"liufanwh"`
	Password string `example:"4dfg54df4g8"`
}

type UserReset struct {
	Email     string `example:"liufanwh@126.com"`
	Password  string `example:"4dfg54df4g8"`
	ResetCode string `example:"THGSDe"`
}
