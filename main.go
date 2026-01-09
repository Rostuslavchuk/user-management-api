package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	gde "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var db *sql.DB

func main() {
	err := gde.Load()
	if err != nil {
		log.Fatal(err)
	}

	user := os.Getenv(strings.ToUpper("user"))
	password := os.Getenv(strings.ToUpper("password"))
	dbname := os.Getenv(strings.ToUpper("dbname"))
	port, _ := strconv.Atoi(os.Getenv(strings.ToUpper("port")))

	connStr := fmt.Sprintf("user='%s' password='%s' dbname='%s' port=%d sslmode=disable",
		user, password, dbname, port)

	var dbErr error
	db, dbErr = sql.Open("postgres", connStr)
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users_test(
			id INT GENERATED ALWAYS AS IDENTITY,
			name VARCHAR(55) NOT NULL,
			age SMALLINT NOT NULL check (age >= 18) default 18
	)`)
	if err != nil {
		log.Fatal(err)
	}

	router := http.NewServeMux()
	router.HandleFunc("GET /users", getUsers)
	router.HandleFunc("GET /users/{id}", getByID)
	router.HandleFunc("DELETE /users/{id}", deleteUser)
	router.HandleFunc("POST /users", createUser)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getByID(r http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	var u User
	err := db.QueryRow("SELECT * FROM users_test WHERE id = $1", id).Scan(&u.ID, &u.Name, &u.Age)
	if err != nil {
		log.Fatal(err)
		return
	}
	r.Header().Set("Content-Type", "application/json")
	r.WriteHeader(http.StatusOK)
	json.NewEncoder(r).Encode(u)
}

func deleteUser(r http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	var u User
	err := db.QueryRow("DELET FROM users_test WHERE id = $1 RETURNING id, name, age", id).Scan(&u.ID, &u.Name, &u.Age)
	if err != nil {
		log.Fatal(err)
		return
	}
	r.Header().Set("Content-Type", "application/json")
	r.WriteHeader(http.StatusOK)
	json.NewEncoder(r).Encode(u)
}

func createUser(r http.ResponseWriter, req *http.Request) {
	var u User
	if err := json.NewDecoder(req.Body).Decode(&u); err != nil {
		log.Fatal(err, http.StatusBadRequest)
		return
	}
	sqlStatement := `INSERT INTO users_test (name, age) VALUES ($1, $2) RETURNING id`
	err := db.QueryRow(sqlStatement, u.Name, u.Age).Scan(&u.ID)
	if err != nil {
		log.Fatal(err)
		return
	}
	r.Header().Set("Content-Type", "application/json")
	r.WriteHeader(http.StatusCreated)
	json.NewEncoder(r).Encode(u)
}

func getUsers(r http.ResponseWriter, req *http.Request) {
	rows, err := db.Query("SELECT * FROM users_test")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			log.Fatal(err)
			return
		}
		users = append(users, u)
	}

	r.Header().Set("Content-Type", "application/json")
	r.WriteHeader(http.StatusAccepted)
	json.NewEncoder(r).Encode(users)
}
