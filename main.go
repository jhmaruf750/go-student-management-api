package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Student struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Department string `json:"department"`
}

var students []Student

const fileName = "students.json"

func main() {

	loadFromFile()

	http.HandleFunc("/students", studentsHandler)

	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

// ================= HANDLER =================

func studentsHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		getStudents(w, r)

	case "POST":
		createStudent(w, r)

	case "PUT":
		updateStudent(w, r)

	case "DELETE":
		deleteStudent(w, r)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// ================= GET =================

func getStudents(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	idParam := r.URL.Query().Get("id")

	if idParam == "" {
		json.NewEncoder(w).Encode(students)
		return
	}

	id, _ := strconv.Atoi(idParam)

	for _, student := range students {
		if student.ID == id {
			json.NewEncoder(w).Encode(student)
			return
		}
	}

	http.Error(w, "Student not found", http.StatusNotFound)
}

// ================= POST =================

func createStudent(w http.ResponseWriter, r *http.Request) {

	var newStudent Student

	err := json.NewDecoder(r.Body).Decode(&newStudent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	students = append(students, newStudent)

	saveToFile()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newStudent)
}

// ================= PUT =================

func updateStudent(w http.ResponseWriter, r *http.Request) {

	idParam := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idParam)

	var updatedStudent Student
	json.NewDecoder(r.Body).Decode(&updatedStudent)

	for i, student := range students {

		if student.ID == id {

			students[i].Name = updatedStudent.Name
			students[i].Age = updatedStudent.Age
			students[i].Department = updatedStudent.Department

			saveToFile()

			json.NewEncoder(w).Encode(students[i])
			return
		}
	}

	http.Error(w, "Student not found", http.StatusNotFound)
}

// ================= DELETE =================

func deleteStudent(w http.ResponseWriter, r *http.Request) {

	idParam := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idParam)

	for i, student := range students {

		if student.ID == id {

			students = append(students[:i], students[i+1:]...)
			saveToFile()

			fmt.Fprintln(w, "Student deleted")
			return
		}
	}

	http.Error(w, "Student not found", http.StatusNotFound)
}

// ================= FILE HANDLING =================

func loadFromFile() {

	file, err := os.ReadFile(fileName)
	if err != nil {
		students = []Student{}
		return
	}

	json.Unmarshal(file, &students)
}

func saveToFile() {

	data, _ := json.MarshalIndent(students, "", "  ")
	os.WriteFile(fileName, data, 0644)
}
