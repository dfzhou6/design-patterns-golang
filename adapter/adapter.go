package adapter

import "fmt"

type ILogin interface {
	Login(username, password string) error
}

type INewLogin interface {
	NewLogin(username, password, captcha string) error
}

type User struct {
}

func (*User) Login(username, password string) error {
	fmt.Printf("username: %s, password: %s\n", username, password)
	return nil
}

type LoginAdapter struct {
	ILogin ILogin
}

func (a *LoginAdapter) NewLogin(username, password, captcha string) error {
	fmt.Println("LoginAdapter start")
	defer fmt.Println("LoginAdapter end")
	fmt.Println("check captcha:", captcha)
	if err := a.ILogin.Login(username, password); err != nil {
		return err
	}
	return nil
}

func NewLoginAdapter(login ILogin) *LoginAdapter {
	return &LoginAdapter{
		ILogin: login,
	}
}
