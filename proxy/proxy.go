package proxy

import "fmt"

type ILogin interface {
	Login(username, password string) error
}

type User struct {
}

func (u *User) Login(username, password string) error {
	fmt.Printf("username: %s, password: %s\n", username, password)
	return nil
}

type UserProxy struct {
	ILogin ILogin
}

func (up *UserProxy) Login(username, password string) error {
	fmt.Println("user proxy do log start")
	defer fmt.Println("user proxy do log end")
	if err := up.ILogin.Login(username, password); err != nil {
		return err
	}
	return nil
}

func NewUserProxy(login ILogin) *UserProxy {
	return &UserProxy{
		ILogin: login,
	}
}
