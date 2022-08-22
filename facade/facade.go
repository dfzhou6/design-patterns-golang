package facade

import (
	"errors"
	"fmt"
)

type User struct {
	Username string
	Password string
}

type ILogin interface {
	Login(username, password, captcha string) error
}

type IGetOneUser interface {
	GetOne(username string) (User, error)
}

type IVerifyCaptcha interface {
	Verify(captcha string) error
}

type LoginFacade struct {
	IGetOneUser    IGetOneUser
	IVerifyCaptcha IVerifyCaptcha
}

func GetLoginFacade() LoginFacade {
	return LoginFacade{
		IGetOneUser:    &UserService{},
		IVerifyCaptcha: &VerifyCaptchaService{},
	}
}

func (l *LoginFacade) Login(username, password, captcha string) error {
	var err error

	if err = l.IVerifyCaptcha.Verify(captcha); err != nil {
		return err
	}
	user, err := l.IGetOneUser.GetOne(username)
	if err != nil {
		return err
	}
	if user.Password != password {
		return errors.New("username or password incorrect")
	}

	return nil
}

type VerifyCaptchaService struct {
}

func (*VerifyCaptchaService) Verify(captcha string) error {
	fmt.Println("VerifyCaptchaService Verify", captcha)
	return nil
}

type UserService struct {
}

func (*UserService) GetOne(username string) (User, error) {
	fmt.Println("UserService GetOne", username)
	return User{Username: username, Password: "123456"}, nil
}
