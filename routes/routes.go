package routes

import (
	"student-enrollment/controllers"

	"github.com/gorilla/mux"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Student routes
	router.HandleFunc("/api/students", controllers.GetStudents).Methods("GET")
	router.HandleFunc("/api/students/{id}", controllers.GetStudent).Methods("GET")
	router.HandleFunc("/api/students", controllers.CreateStudent).Methods("POST")
	router.HandleFunc("/api/students/{id}", controllers.UpdateStudent).Methods("PUT")
	router.HandleFunc("/api/students/{id}", controllers.DeleteStudent).Methods("DELETE")

	return router
}
