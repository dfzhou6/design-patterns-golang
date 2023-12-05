package adapter

import "testing"

const (
	username = "felix"
	password = "123456"
	captcha  = "abcefg"
)

func TestAdapter(t *testing.T) {
	var iLogin ILogin
	var cLogin CaptchaILogin
	iLogin = NewOriginalLogin()
	cLogin = NewAliYunCaptchaLogin(iLogin)
	cLogin.LoginWithCaptcha(username, password, captcha)
}
