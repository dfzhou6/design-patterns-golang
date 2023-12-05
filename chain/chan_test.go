package chain

import "testing"

func TestChain(t *testing.T) {
	hr := NewHRLeaveRequestHandler()
	manager := NewManagerLeaveRequestHandler()
	director := NewDirectorLeaveRequestHandler()

	hr.SetNextHandler(manager)
	manager.SetNextHandler(director)

	req := &LeaveRequest{Days: 3}
	hr.HandleLeaveRequest(req)

	req.Days = 7
	hr.HandleLeaveRequest(req)

	req.Days = 12
	hr.HandleLeaveRequest(req)
}
