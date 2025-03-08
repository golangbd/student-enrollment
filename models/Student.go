package models

import (
	"time"

	"student-enrollment/config"
)

type Student struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	DOB       time.Time `json:"dob"`
	Major     string    `json:"major"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetAllStudents retrieves all students from the database
func GetAllStudents() ([]Student, error) {
	var students []Student

	db, err := config.GetDB()
	if err != nil {
		return nil, err
	}
	
	rows, err := db.Query("SELECT id, first_name, last_name, email, dob, major, created_at, updated_at FROM students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var student Student
		var dobStr, createdAtStr, updatedAtStr string // Temporary strings to hold date values
		
		err := rows.Scan(&student.ID, &student.FirstName, &student.LastName, &student.Email, 
			&dobStr, &student.Major, &createdAtStr, &updatedAtStr)
		if err != nil {
			return nil, err
		}
		
		// Parse the date string into a time.Time if it's not empty
		if dobStr != "" {
			dob, err := time.Parse("2006-01-02", dobStr)
			if err != nil {
				return nil, err
			}
			student.DOB = dob
		}
		
		// Parse the timestamp strings
		if createdAtStr != "" {
			createdAt, err := time.Parse("2006-01-02 15:04:05", createdAtStr)
			if err != nil {
				return nil, err
			}
			student.CreatedAt = createdAt
		}
		
		if updatedAtStr != "" {
			updatedAt, err := time.Parse("2006-01-02 15:04:05", updatedAtStr)
			if err != nil {
				return nil, err
			}
			student.UpdatedAt = updatedAt
		}
		
		students = append(students, student)
	}

	return students, nil
}

// GetStudentByID retrieves a student by ID
func GetStudentByID(id int) (Student, error) {
	var student Student
	var dobStr, createdAtStr, updatedAtStr string // Temporary strings to hold date values
	
	db, err := config.GetDB()
	if err != nil {
		return Student{}, err
	}

	err = db.QueryRow("SELECT id, first_name, last_name, email, dob, major, created_at, updated_at FROM students WHERE id = ?", id).
		Scan(&student.ID, &student.FirstName, &student.LastName, &student.Email, 
			&dobStr, &student.Major, &createdAtStr, &updatedAtStr)
	if err != nil {
		return Student{}, err
	}
	
	// Parse the date string into a time.Time if it's not empty
	if dobStr != "" {
		dob, err := time.Parse("2006-01-02", dobStr)
		if err != nil {
			return Student{}, err
		}
		student.DOB = dob
	}
	
	// Parse the timestamp strings
	if createdAtStr != "" {
		createdAt, err := time.Parse("2006-01-02 15:04:05", createdAtStr)
		if err != nil {
			return Student{}, err
		}
		student.CreatedAt = createdAt
	}
	
	if updatedAtStr != "" {
		updatedAt, err := time.Parse("2006-01-02 15:04:05", updatedAtStr)
		if err != nil {
			return Student{}, err
		}
		student.UpdatedAt = updatedAt
	}

	return student, nil
}

// CreateStudent adds a new student to the database
func CreateStudent(student Student) (int64, error) {
	db, err := config.GetDB()
	if err != nil {
		return 0, err
	}

	result, err := db.Exec("INSERT INTO students (first_name, last_name, email, dob, major, created_at, updated_at) VALUES (?, ?, ?, ?, ?, NOW(), NOW())",
		student.FirstName, student.LastName, student.Email, student.DOB, student.Major)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// UpdateStudent updates an existing student
func UpdateStudent(student Student) error {
	db, err := config.GetDB()
	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE students SET first_name = ?, last_name = ?, email = ?, dob = ?, major = ?, updated_at = NOW() WHERE id = ?",
		student.FirstName, student.LastName, student.Email, student.DOB, student.Major, student.ID)
	return err
}

// DeleteStudent removes a student from the database
func DeleteStudent(id int) error {
	db, err := config.GetDB()
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM students WHERE id = ?", id)
	return err
}
