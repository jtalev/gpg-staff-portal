package mockdb

import (
	"time"

	"github.com/jtalev/gpg-staff-portal/internal/entities"
)

var lrList = []entities.LeaveRequest{
	{
		Id:         0,
		EmployeeId: 000000,
		Type:       "Annual",
		StartDate:  time.Date(2024, time.July, 24, 0, 0, 0, 0, nil),
		EndDate:	time.Date(2024, time.August, 10, 0, 0, 0, 0, nil),
		TotalDays:  18,
		IsApproved: true,
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	},
	{
		Id:         1,
		EmployeeId: 000001,
		Type:       "Annual",
		StartDate:  time.Date(2024, time.July, 24, 0, 0, 0, 0, nil),
		EndDate:	time.Date(2024, time.August, 10, 0, 0, 0, 0, nil),
		TotalDays:  18,
		IsApproved: false,
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	},
}

func GetAllLeaveRequests() (l *[]entities.LeaveRequest, msg string) {
	
	return &lrList, "Successfully fetched leave requests"
}

func GetLeaveRequestsByEmployeeId(empId int) (l *[]entities.LeaveRequest, msg string) {

	list := make([]entities.LeaveRequest, 0)
	for _, lr := range lrList {
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

	for _, r := range lrList {
		if r.Id == lr.Id {
			return nil, "This leave request already exists"
		}
	}

	lrList = append(lrList, lr)
	return &lr, "Successfully created leave request"
}

func UpdateLeaveRequest(lr entities.LeaveRequest, id int) (l *entities.LeaveRequest, msg string) {

	for i, r := range lrList {
		if r.Id == id {
			lrList[i] = lr
			return &lr, "successfully updated leave request"
		}
	}

	return nil, "No leave request with given id"
}

func DeleteLeaveRequest(id int) (l *[]entities.LeaveRequest, msg string) {

	for i, r := range lrList {
		if r.Id == id {
			before := lrList[:i]
			after := lrList[i+1:]
			lrList = append(before, after...)
			return &lrList, "Successfully deleted leave request"
		}
	}
	return nil, "No leave request with given id"
}