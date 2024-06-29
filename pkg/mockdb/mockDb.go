package mockdb

type Db struct {
	UserRepo 		UserRepo
	LeaveRepo		LeaveRepo
	TimesheetRepo	TimesheetRepo
}

func Connect() Db {
	db := Db{
		UserRepo: 		UserRepo{},
		LeaveRepo: 		LeaveRepo{},
		TimesheetRepo: 	TimesheetRepo{},
	}

	return db
}