package persistance

import (
	"log"
	"os/exec"
)

func CheckSQLite() {
	cmd := exec.Command("sqlite3", "--version")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal("Failed to create database: %v", err)
	}
	log.Printf("SQLite is installed: %s\n", output)

}

func CreateSQLiteDB() {
	cmd := exec.Command("sqlite3", "example.db", "CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT, age INTEGER);")
	err := cmd.Run()
	if err != nil {
		log.Fatal("Failed to create database: %v", err)
	}
	log.Println("Database created succesfully!")
}
