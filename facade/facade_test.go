package facade

import "testing"

const (
	username = "felix"
	password = "123456"
	email    = "zhoudf6@gmail.com"
)

func TestUserFacade(t *testing.T) {
	f := NewUserFacade()
	f.Login(username, password)
	f.Register(username, password, email)
}
