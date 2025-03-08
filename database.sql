CREATE DATABASE IF NOT EXISTS student_enrollment;

USE student_enrollment;

CREATE TABLE IF NOT EXISTS students (
    id INT AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    dob DATE,
    major VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Sample data
INSERT INTO students (first_name, last_name, email, dob, major) VALUES
('John', 'Doe', 'john.doe@example.com', '2000-01-15', 'Computer Science'),
('Jane', 'Smith', 'jane.smith@example.com', '2001-05-20', 'Mathematics'),
('Bob', 'Johnson', 'bob.johnson@example.com', '1999-11-10', 'Physics');