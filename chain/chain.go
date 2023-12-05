package chain

import "fmt"

type LeaveRequest struct {
	Days int
}

type LeaveRequestHandler interface {
	SetNextHandler(handler LeaveRequestHandler)
	HandleLeaveRequest(req *LeaveRequest)
}

type BaseLeaveRequestHandler struct {
	nextHandler LeaveRequestHandler
}

func (impl *BaseLeaveRequestHandler) SetNextHandler(handler LeaveRequestHandler) {
	impl.nextHandler = handler
}

func (impl *BaseLeaveRequestHandler) HandleLeaveRequest(req *LeaveRequest) {
	panic("pls implement me")
}

type HRLeaveRequestHandler struct {
	*BaseLeaveRequestHandler
}

func NewHRLeaveRequestHandler() *HRLeaveRequestHandler {
	return &HRLeaveRequestHandler{&BaseLeaveRequestHandler{}}
}

func (impl *HRLeaveRequestHandler) HandleLeaveRequest(req *LeaveRequest) {
	if req.Days < 5 {
		fmt.Printf("hr handle this leave request, days is %v\n", req.Days)
		return
	}
	fmt.Printf("hr handle can't do this leave request, days is %v\n", req.Days)
	if impl.nextHandler != nil {
		impl.nextHandler.HandleLeaveRequest(req)
	}
}

type ManagerLeaveRequestHandler struct {
	*BaseLeaveRequestHandler
}

func NewManagerLeaveRequestHandler() *ManagerLeaveRequestHandler {
	return &ManagerLeaveRequestHandler{&BaseLeaveRequestHandler{}}
}

func (impl *ManagerLeaveRequestHandler) HandleLeaveRequest(req *LeaveRequest) {
	if req.Days < 10 {
		fmt.Printf("manager handle this leave request, days is %v\n", req.Days)
		return
	}
	fmt.Printf("manager handle can't do this leave request, days is %v\n", req.Days)
	if impl.nextHandler != nil {
		impl.nextHandler.HandleLeaveRequest(req)
	}
}

type DirectorLeaveRequestHandler struct {
	*BaseLeaveRequestHandler
}

func NewDirectorLeaveRequestHandler() *DirectorLeaveRequestHandler {
	return &DirectorLeaveRequestHandler{&BaseLeaveRequestHandler{}}
}

func (impl *DirectorLeaveRequestHandler) HandleLeaveRequest(req *LeaveRequest) {
	fmt.Printf("director handle this leave request, days is %v\n", req.Days)
}
