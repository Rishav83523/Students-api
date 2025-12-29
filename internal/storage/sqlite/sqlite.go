package sqlite

import (
	"database/sql"

	_ "github.com/glebarez/sqlite" // pure-Go sqlite driver (no cgo)
	"github.com/mr-raj2001/students-api/internal/config"
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

