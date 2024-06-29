package mockdb

import (
	"time"

	"github.com/jtalev/gpg-staff-portal/pkg/types"
)

type LeaveRepo struct {}

var leaveRequests = []types.LeaveRequest{
	{
		Id:         0,
		EmployeeId: 000000,
		Type:       "Annual",
		StartDate:  time.Date(2024, time.July, 24, 0, 0, 0, 0, time.Now().Location()),
		EndDate:	time.Date(2024, time.August, 10, 0, 0, 0, 0, time.Now().Location()),
		TotalDays:  18,
		IsApproved: true,
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	},
	{
		Id:         1,
		EmployeeId: 000001,
		Type:       "Annual",
		StartDate:  time.Date(2024, time.July, 24, 0, 0, 0, 0, time.Now().Location()),
		EndDate:	time.Date(2024, time.August, 10, 0, 0, 0, 0, time.Now().Location()),
		TotalDays:  18,
		IsApproved: false,
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	},
}

func (l *LeaveRepo) GetAllLeaveRequests() (lr *[]types.LeaveRequest, msg string) {
	return &leaveRequests, "Successfully fetched leave requests"
}

func (l *LeaveRepo) GetLeaveRequestsByEmployeeId(empId int) (lr *[]types.LeaveRequest, msg string) {
	list := make([]types.LeaveRequest, 0)
	for _, lr := range leaveRequests {
		if lr.EmployeeId == empId {
			list = append(list, lr)
		}
	}
	if len(list) == 0 {
		return nil, "No leave requests for this user"
	}

	return &list, "Successfully fetched leave requests"
}

func (l *LeaveRepo) CreateLeaveRequest(leaveRequest types.LeaveRequest) (lr *types.LeaveRequest, msg string) {
	for _, r := range leaveRequests {
		if r.Id == leaveRequest.Id {
			return nil, "This leave request already exists"
		}
	}

	leaveRequests = append(leaveRequests, leaveRequest)
	return &leaveRequest, "Successfully created leave request"
}

func (l *LeaveRepo) UpdateLeaveRequest(leaveRequest types.LeaveRequest, id int) (lr *types.LeaveRequest, msg string) {
	for i, r := range leaveRequests {
		if r.Id == id {
			leaveRequests[i] = leaveRequest
			return &leaveRequest, "successfully updated leave request"
		}
	}

	return nil, "No leave request with given id"
}

func (l *LeaveRepo) DeleteLeaveRequest(id int) (lr *[]types.LeaveRequest, msg string) {
	for i, r := range leaveRequests {
		if r.Id == id {
			before := leaveRequests[:i]
			after := leaveRequests[i+1:]
			leaveRequests = append(before, after...)
			return &leaveRequests, "Successfully deleted leave request"
		}
	}
	return nil, "No leave request with given id"
}