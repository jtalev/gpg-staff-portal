package api

import (
	"github.com/gorilla/mux"
	"github.com/jtalev/gpg-staff-portal/api/handler"
)

func SetupRoutes(r *mux.Router) {

	// web
	r.HandleFunc("/", handler.LoginHandler)
	r.HandleFunc("/home", handler.HomeHandler)
	r.HandleFunc("/clock-on", handler.ClockOnHandler)
	r.HandleFunc("/leave", handler.LeaveHandler)

	// user
	r.HandleFunc("/user", handler.GetUsersHandler).Methods("GET")
	r.HandleFunc("/user/{id}", handler.GetUserByIdHandler).Methods("GET")
	r.HandleFunc("/user/{employeeId}", handler.GetUserByEmployeeIdHandler).Methods("GET")
	r.HandleFunc("/user/create", handler.CreateUserHandler).Methods("POST")
	r.HandleFunc("/user/update/{id}", handler.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/user/delete/{id}", handler.DeleteUserHandler).Methods("DELETE")

	// timesheet

	// paycycle

	// leave request
}
