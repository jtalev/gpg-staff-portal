package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jtalev/gpg-staff-portal/api/controllers"
)

func HandleRequests(r *mux.Router) {
	user := &controllers.UserHandler{}
	// timesheet := &controllers.TimesheetHandler{}
	// paycycle := &controllers.PayCycleHandler{}
	// leaveRequest := &controllers.LeaveRequestHandler{}
	
	r.HandleFunc("/", user.ServeHttp).Methods("GET")

	// users endpoints
	r.HandleFunc("/users", user.GetAllUsers).Methods("GET")
	r.HandleFunc("/users/create", user.CreateUser).Methods("POST")
	r.HandleFunc("/users/read/{id}", user.ReadUser).Methods("GET")
	r.HandleFunc("/users/update/{id}", user.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/delete/{id}", user.DeleteUser).Methods("DELETE")

	// // timesheet endpoints
	// r.HandleFunc("/timesheet/create", timesheet.CreateTimesheet).Methods("POST")
	// r.HandleFunc("/timesheet/read/{id}", timesheet.ReadTimesheet).Methods("GET")
	// r.HandleFunc("/timesheet/read-by-employee/{employeeId}", timesheet.ReadTimesheetByEmployee).Methods("GET")
	// r.HandleFunc("/timesheet/update/{id}", timesheet.UpdateTimesheet).Methods("PUT")
	// r.HandleFunc("/timesheet/delete/{id}", timesheet.DeleteTimesheet).Methods("DELETE")

	// // paycycle endpoints
	// r.HandleFunc("/paycycle/create", paycycle.CreatePaycycle).Methods("POST")
	// r.HandleFunc("/paycycle/read/{id}", paycycle.ReadPaycycle).Methods("GET")
	// r.HandleFunc("/paycycle/read/{employeeId}", paycycle.ReadPaycycleByEmployee).Methods("GET")
	// r.HandleFunc("/paycycle/update/{id}", paycycle.UpdatePaycycle).Methods("PUT")
	// r.HandleFunc("/paycycle/delete/{id}", paycycle.DeletePaycycle).Methods("DELETE")

	// // leave request endpoints
	// r.HandleFunc("/leave/create", leaveRequest.CreateLeaveRequest).Methods("POST")
	// r.HandleFunc("leave/read/{id}", leaveRequest.ReadLeaveRequest).Methods("GET")
	// r.HandleFunc("leave/read-by-employee/{employeeId}", leaveRequest.ReadLeaveRequestByEmployee).Methods("GET")
	// r.HandleFunc("leave/update/{id}", leaveRequest.UpdateLeaveRequest).Methods("PUT")
	// r.HandleFunc("leave/delete/{id}", leaveRequest.DeleteLeaveRequest).Methods("DELETE")

	// todo: add config variables for port number, isDevelopment, isProduction etc.
	fmt.Println("Starting server on localhost:5000")
	err := http.ListenAndServe(":5000", r)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func main() {
	r := mux.NewRouter()
	HandleRequests(r)
}
