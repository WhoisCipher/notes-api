package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/lib/pq"
)

func main()  {
    dbHost := os.Getenv("DB_HOST")
    dbUser := os.Getenv("DB_USER")
    dbPort := os.Getenv("DB_PORT")
    dbPass := os.Getenv("DB_PASS")
    dbName := os.Getenv("DB_NAME")

    connStr := fmt.Sprintf("host=%s user=%s port=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPort, dbPass, dbName)

    var err error
    db, err := sql.Open("postgres", connStr)

    if err != nil {
        log.Fatal("Connection to database Failed", err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatal("Ping to database Failed", err)
    }

    fmt.Println("Connected to database")
}
