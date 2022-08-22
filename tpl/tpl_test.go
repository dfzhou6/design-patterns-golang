package tpl

import "testing"

const number = "13160677729"

func TestNewHuawei(t *testing.T) {
	h := NewHuawei()
	if err := h.CallNumber(number); err != nil {
		t.Errorf("CallNumber failed, err: %v", err)
	}
}

func TestNewSamsung(t *testing.T) {
	s := NewSamsung()
	if err := s.CallNumber(number); err != nil {
		t.Errorf("CallNumber failed, err: %v", err)
	}
}
