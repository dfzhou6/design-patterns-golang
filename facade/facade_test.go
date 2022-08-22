package facade

import "testing"

func TestLoginFacade(t *testing.T) {
	var err error

	f := GetLoginFacade()
	if err = f.Login("felix", "123456", "abcdef"); err != nil {
		t.Error(err)
	}
	if err = f.Login("felix", "654321", "abcdef"); err == nil {
		t.Error("should not be success")
	}
	t.Log(err)
}
