package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"student-enrollment/models"

	"github.com/gorilla/mux"
)

// GetStudents handles the request to get all students
func GetStudents(w http.ResponseWriter, r *http.Request) {
	students, err := models.GetAllStudents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

// GetStudent handles the request to get a specific student
func GetStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	student, err := models.GetStudentByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

// CreateStudent handles the request to create a new student
func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate required fields
	if student.FirstName == "" || student.LastName == "" || student.Email == "" {
		http.Error(w, "First name, last name, and email are required", http.StatusBadRequest)
		return
	}

	id, err := models.CreateStudent(student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "Student created successfully",
		"id":      id,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// UpdateStudent handles the request to update a student
func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	var student models.Student
	err = json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	student.ID = id
	err = models.UpdateStudent(student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"message": "Student updated successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DeleteStudent handles the request to delete a student
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	err = models.DeleteStudent(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"message": "Student deleted successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
