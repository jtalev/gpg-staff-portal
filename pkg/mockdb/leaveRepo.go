package mockdb

import (
	"time"

	"github.com/jtalev/gpg-staff-portal/pkg/entities"
)

var leaveRequests = []entities.LeaveRequest{
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

func GetAllLeaveRequests() (l *[]entities.LeaveRequest, msg string) {
	return &leaveRequests, "Successfully fetched leave requests"
}

func GetLeaveRequestsByEmployeeId(empId int) (l *[]entities.LeaveRequest, msg string) {
	list := make([]entities.LeaveRequest, 0)
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

func CreateLeaveRequest(lr entities.LeaveRequest) (l *entities.LeaveRequest, msg string) {
	for _, r := range leaveRequests {
		if r.Id == lr.Id {
			return nil, "This leave request already exists"
		}
	}

	leaveRequests = append(leaveRequests, lr)
	return &lr, "Successfully created leave request"
}

func UpdateLeaveRequest(lr entities.LeaveRequest, id int) (l *entities.LeaveRequest, msg string) {
	for i, r := range leaveRequests {
		if r.Id == id {
			leaveRequests[i] = lr
			return &lr, "successfully updated leave request"
		}
	}

	return nil, "No leave request with given id"
}

func DeleteLeaveRequest(id int) (l *[]entities.LeaveRequest, msg string) {
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