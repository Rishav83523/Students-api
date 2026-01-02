package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/glebarez/sqlite" // pure-Go sqlite driver (no cgo)
	"github.com/mr-raj2001/students-api/internal/config"
	"github.com/mr-raj2001/students-api/internal/types"
)
type Sqlite struct {
	Db *sql.DB       // Db is the database connection 
}


func New (cfg *config.Config) (*Sqlite,error){ 
	db,err := sql.Open("sqlite", cfg.StoragePath)        // opening a connection to sqlite database using the storage path from config (glebarez driver)
     if err != nil {
		return nil, err
	 }

	 _, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
	         id INTEGER PRIMARY KEY AUTOINCREMENT,
	         name TEXT NOT NULL,
	         email TEXT NOT NULL,
			 age INTEGER NOT NULL
	 )`)


	 if err != nil {
		return nil, err
	 }

	 return &Sqlite{Db: db},nil   // returning the pointer to Sqlite struct with Db field initialized and nil error

}


// Implementing the Storage interface's CreateStudent method
func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {

	// Inserting a new student record into the students table
	stmt, err := s.Db.Prepare("INSERT INTO students (name,email,age) VALUES (?,?,?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()   // ensuring the statement is closed after execution

	result, err := stmt.Exec(name, email, age)
	if err != nil {
		return 0, err
	}

     id, err := result.LastInsertId()

	 if err != nil {
		return 0, err
	 }

	return id, nil
}


func (s *Sqlite) GetStudentById(id int64) (*types.Student, error) { 
      stmt,err := s.Db.Prepare ("SELECT * FROM students WHERE id = ? LIMIT 1")

	  if err != nil { 
		return &types.Student{}, err
	  }

	  defer stmt.Close()

	  var student types.Student

	 err = stmt.QueryRow(id).Scan(&student.Id, &student.Name, &student.Email, &student.Age)
	 if err != nil {

		if err == sql.ErrNoRows { 
			return &types.Student{}, fmt.Errorf("student not found")
		}
		return &types.Student{}, fmt.Errorf("query error: %w", err)
	 }
	 return &student, nil
}


func (s *Sqlite) GetStudentList () ([]*types.Student, error) { 
	stmt, err := s.Db.Prepare("SELECT * FROM students")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var students []*types.Student

	for rows.Next() {
		var student types.Student

		err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Age)
		if err != nil {
			return nil, err
		}

		students = append(students, &student)
	}

	return students, nil
}