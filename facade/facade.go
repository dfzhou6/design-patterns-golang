package facade

import "fmt"

type UserLogin interface {
	Login(username string, password string)
}

type UserRegister interface {
	Register(username string, password string, email string)
}

type UserLoginImpl struct {
}

func (impl *UserLoginImpl) Login(username string, password string) {
	fmt.Printf("user login username %v password %v\n", username, password)
}

type UserRegisterImpl struct {
}

func (impl *UserRegisterImpl) Register(username string, password string, email string) {
	fmt.Printf("user register username %v password %v email %v\n", username, password, email)
}

type UserFacade struct {
	login    UserLogin
	register UserRegister
}

func NewUserFacade() *UserFacade {
	return &UserFacade{
		login:    &UserLoginImpl{},
		register: &UserRegisterImpl{},
	}
}

func (impl *UserFacade) Login(username string, password string) {
	impl.login.Login(username, password)
}

func (impl *UserFacade) Register(username string, password string, email string) {
	impl.register.Register(username, password, email)
}
