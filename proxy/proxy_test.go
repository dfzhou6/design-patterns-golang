package proxy

import "testing"

const (
	username = "felix"
	password = "123456"
)

func TestNewUserProxy(t *testing.T) {
	login := &User{}
	up := NewUserProxy(login)
	if err := up.Login(username, password); err != nil {
		t.Errorf("login failed, err: %v", err)
	}
}
