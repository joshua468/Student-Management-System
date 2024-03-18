
package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/mattn/go-sqlite3"
)

type Student struct {
    ID   int
    Name string
    Age  int
    Grade int
}

func main() {
    
    db, err := sql.Open("sqlite3", "./student.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        age INTEGER,
        grade INTEGER
    )`)
    if err != nil {
        log.Fatal(err)
    }

    menu := []string{
        "1. Add Student",
        "2. View Students",
        "3. Update Student",
        "4. Delete Student",
        "5. Exit",
    }

    for {
        fmt.Println("Student Management System")
        for _, option := range menu {
            fmt.Println(option)
        }
        fmt.Print("Enter your choice: ")

        var choice int
        _, err := fmt.Scanln(&choice)
        if err != nil {
            fmt.Println("Invalid input. Please enter a number.")
            continue
        }

        switch choice {
        case 1:
            addStudent(db)
        case 2:
            viewStudents(db)
        case 3:
            updateStudent(db)
        case 4:
            deleteStudent(db)
        case 5:
            fmt.Println("Exiting...")
            os.Exit(0)
        default:
            fmt.Println("Invalid choice. Please enter a number between 1 and 5.")
        }
    }
}

func addStudent(db *sql.DB) {
    var name string
    var age, grade int
    fmt.Print("Enter student name: ")
    fmt.Scanln(&name)
    fmt.Print("Enter student age: ")
    fmt.Scanln(&age)
    fmt.Print("Enter student grade: ")
    fmt.Scanln(&grade)

    _, err := db.Exec("INSERT INTO students (name, age, grade) VALUES (?, ?, ?)", name, age, grade)
    if err != nil {
        fmt.Println("Error adding student:", err)
        return
    }
    fmt.Println("Student added successfully.")
}

func viewStudents(db *sql.DB) {
    rows, err := db.Query("SELECT id, name, age, grade FROM students")
    if err != nil {
        fmt.Println("Error fetching students:", err)
        return
    }
    defer rows.Close()

    fmt.Println("ID\tName\tAge\tGrade")
    for rows.Next() {
        var s Student
        err := rows.Scan(&s.ID, &s.Name, &s.Age, &s.Grade)
        if err != nil {
            fmt.Println("Error reading student:", err)
            return
        }
        fmt.Printf("%d\t%s\t%d\t%d\n", s.ID, s.Name, s.Age, s.Grade)
    }
}


func updateStudent(db *sql.DB) {
    var id, age, grade int
    var name string
    fmt.Print("Enter student ID to update: ")
    fmt.Scanln(&id)
    fmt.Print("Enter updated student name: ")
    fmt.Scanln(&name)
    fmt.Print("Enter updated student age: ")
    fmt.Scanln(&age)
    fmt.Print("Enter updated student grade: ")
    fmt.Scanln(&grade)

    _, err := db.Exec("UPDATE students SET name=?, age=?, grade=? WHERE id=?", name, age, grade, id)
    if err != nil {
        fmt.Println("Error updating student:", err)
        return
    }
    fmt.Println("Student updated successfully.")
}

func deleteStudent(db *sql.DB) {
    var id int
    fmt.Print("Enter student ID to delete: ")
    fmt.Scanln(&id)

    _, err := db.Exec("DELETE FROM students WHERE id=?", id)
    if err != nil {
        fmt.Println("Error deleting student:", err)
        return
    }
    fmt.Println("Student deleted successfully.")
}
