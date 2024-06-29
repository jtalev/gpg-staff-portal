package mockdb_test

import (
	"testing"
	"time"

	"github.com/jtalev/gpg-staff-portal/pkg/mockdb"
)

var tr = mockdb.TimesheetRepo{}

// todo: test GetTimesheetById

func TestGetTimesheetById(t *testing.T) {
	ts, msg := tr.GetTimesheetsByEmployeeId(1234567)

	if msg == "User not assigned to any timesheets" {
		t.Errorf("User not assigned to any timesheets")
	}

	if ts == nil {
		t.Errorf("No timesheets returned")
	}
}

func TestGetTimesheetWithinRange(t *testing.T) {
	tr := mockdb.TimesheetRepo{}
	date := time.Date(2024, time.June, 17, 8, 0, 0, 0, time.Now().Location())
	ts, _ := tr.GetTimesheetWithinRange(1234567, date, date)

	if len(*ts) != 2 {
		t.Errorf("Incorrect number of timesheets returned. want=%v, got=%v", 2, len(*ts))
	}
}
