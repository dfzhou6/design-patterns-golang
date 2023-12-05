package adapter

import "fmt"

type ILogin interface {
	Login(username string, password string)
}

type OriginalLogin struct {
}

func NewOriginalLogin() *OriginalLogin {
	return &OriginalLogin{}
}

func (impl *OriginalLogin) Login(username string, password string) {
	fmt.Printf("OriginalLogin username %v password %v\n", username, password)
}

type CaptchaILogin interface {
	LoginWithCaptcha(username string, password string, captcha string)
}

type AliYunCaptchaLogin struct {
	i ILogin
}

func NewAliYunCaptchaLogin(iLogin ILogin) *AliYunCaptchaLogin {
	return &AliYunCaptchaLogin{
		i: iLogin,
	}
}

func (impl *AliYunCaptchaLogin) LoginWithCaptcha(username string, password string, captcha string) {
	if len(captcha) > 0 {
		fmt.Printf("check captcha %v\n", captcha)
	}
	impl.i.Login(username, password)
}
