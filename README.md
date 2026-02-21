# 🎓 Go Student Management API

A file-based RESTful API built using **Go (Golang)** that performs full CRUD operations for managing student data.

This project is built as part of backend learning to understand HTTP servers, JSON handling, file persistence, and REST architecture using pure Go.

---

## 🚀 Features

* ✅ Get all students
* ✅ Get single student by ID
* ✅ Create new student
* ✅ Update existing student
* ✅ Delete student
* ✅ File-based data persistence (JSON storage)
* ✅ RESTful API design
* ✅ Proper HTTP status handling

---

## 🛠 Tech Stack

* **Language:** Go (Golang)
* **Server:** net/http
* **Data Storage:** JSON file (students.json)
* **Architecture:** Basic REST API structure

---

## 📁 Project Structure

```
go-student-management-api/
│── main.go
│── students.json
│── go.mod
│── README.md
```

---

## ▶️ How to Run

### 1️⃣ Clone the repository

```bash
git clone https://github.com/jhmaruf750/go-student-management-api.git
cd go-student-management-api
```

### 2️⃣ Run the server

```bash
go run main.go
```

Server will start at:

```
http://localhost:8080
```

---

## 📌 API Endpoints

### 🔹 Get All Students

```
GET /students
```

---

### 🔹 Get Single Student

```
GET /students?id=1
```

---

### 🔹 Create Student

```
POST /students
```

**Body (JSON):**

```json
{
  "id": 1,
  "name": "Rahim",
  "age": 22,
  "department": "CSE"
}
```

---

### 🔹 Update Student

```
PUT /students?id=1
```

**Body (JSON):**

```json
{
  "name": "Updated Name",
  "age": 25,
  "department": "EEE"
}
```

---

### 🔹 Delete Student

```
DELETE /students?id=1
```

---

## 🧠 Concepts Practiced

* Struct and JSON tags
* HTTP methods (GET, POST, PUT, DELETE)
* Query parameters
* JSON encoding and decoding
* File read/write in Go
* Slice manipulation
* Basic REST API design
* Error handling

---

## 📚 Learning Purpose

This project was built to strengthen backend fundamentals using pure Go without any external frameworks.

---

## 🔮 Future Improvements

* Add proper routing (Chi or Gorilla Mux)
* Add PostgreSQL database support
* Add authentication (JWT)
* Implement Clean Architecture
* Add Docker support

---

## 👨‍💻 Author

**Maruf**
GitHub: https://github.com/jhmaruf750

---

⭐ If you found this project useful, feel free to star the repository.
