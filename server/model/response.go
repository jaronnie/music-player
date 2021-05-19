package model

//响应
type Error struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

type Success struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

//简单工厂模式
func NewError() *Error {
	return &Error{
		Code: 404,
		Message: "not found resource",
	}
}

func NewUserHasExistedErr(username string) *Error {

	return &Error{
		Code: 2001,
		Message: "the user named " + username + " has existed...",
	}
}

func NewLoginErr() *Error {
	return &Error{
		Code: 2002,
		Message: "The password is incorrect",
	}
}

func NewLoginSuccess(token string) *Success {
	return &Success{
		Code: 200,
		Message: "login successfully",
		Data: token,
	}
}
