package postgres

import (
    "database/sql"
)

type Storage interface{
}


type Postgresql struct{
    db *sql.DB
}

func ConnectDB() *sql.DB {
    connStr := "user=postgres dbname=todo password=mysecretpass sslmode=disable"
    
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    }
    
    err = db.Ping()
    if err != nil {
        panic(err)
    }
    
    return db
}


func (s *Postgresql) CreateTable() error {
    query := `CREATE TABLE IF NOT EXISTS(
        ID SERIAL PRIMARY KEY,
        CONTENT TEXT,
        
        )`
    _, err := s.db.Exec(query)
    return err 
}


func (s *Postgresql) CreateTodo()  {
    
}



