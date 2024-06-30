package api

import (
	"github.com/gorilla/mux"
	"github.com/jtalev/gpg-staff-portal/api/handler"
	"github.com/jtalev/gpg-staff-portal/api/handler/endpoints"
	"github.com/jtalev/gpg-staff-portal/pkg/mockdb"
)

func SetupRoutes(r *mux.Router, db mockdb.Db) {
	u := endpoints.UserHandler{Db: db}
	t := endpoints.TimesheetHandler{Db: db}

	// web
	r.HandleFunc("/", handler.LoginHandler)
	r.HandleFunc("/home", handler.HomeHandler)
	r.HandleFunc("/clock-on", handler.ClockOnHandler)
	r.HandleFunc("/leave", handler.LeaveHandler)

	// user
	r.HandleFunc("/user", u.GetUsersHandler).Methods("GET")
	r.HandleFunc("/user/{id}", u.GetUserByIdHandler).Methods("GET")
	r.HandleFunc("/user/{employeeId}", u.GetUserByEmployeeIdHandler).Methods("GET")
	r.HandleFunc("/user/create", u.CreateUserHandler).Methods("POST")
	r.HandleFunc("/user/update/{id}", u.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/user/delete/{id}", u.DeleteUserHandler).Methods("DELETE")

	// timesheet
	r.HandleFunc("/timesheet/create", t.CreateTimesheetHandler).Methods("POST")

	// paycycle

	// leave request
}
