package proxy

import "fmt"

type LoginService interface {
	Login(username, password string)
}

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (impl *UserService) Login(username, password string) {
	fmt.Printf("user service login user, username: %v, password: %v\n", username, password)
}

type UserServiceProxy struct {
	loginSrv LoginService
}

func NewUserServiceProxy(loginSrv LoginService) *UserServiceProxy {
	return &UserServiceProxy{
		loginSrv: loginSrv,
	}
}

func (impl *UserServiceProxy) checkAccess(username string) {
	fmt.Printf("user service proxy check access %v\n", username)
}

func (impl *UserServiceProxy) Login(username, password string) {
	// check access
	impl.checkAccess(username)

	// do login
	impl.loginSrv.Login(username, password)
}
