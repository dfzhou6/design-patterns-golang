package proxy

import "testing"

const (
	username = "felix"
	password = "123456"
)

func TestProxy(t *testing.T) {
	userSrv := NewUserService()
	userSrv.Login(username, password)
	userSrvPx := NewUserServiceProxy(userSrv)
	userSrvPx.Login(username, password)
}
