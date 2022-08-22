package adapter

import "testing"

const (
	username = "felix"
	password = "123456"
	captcha  = "abcefg"
)

func TestNewLoginAdapter(t *testing.T) {
	user := &User{}
	adapter := NewLoginAdapter(user)
	if err := adapter.NewLogin(username, password, captcha); err != nil {
		t.Errorf("NewLogin failed, err: %v", err)
	}
}
