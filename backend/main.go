package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"

	. "marc/todoapp/backend/types"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var schema = `
CREATE TABLE users (
    first_name text,
    last_name text,
    email text
)
`

var mkey string

func init() {
	fmt.Println(runtime.NumCPU())
	ky := &mkey
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	*ky = os.Getenv("MONGODB")
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/users/create", users).Methods("POST")
	r.HandleFunc("/users/get", users).Methods("GET")
	r.HandleFunc("/api/get", api).Methods("GET")
	//	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/build")))
	log.Fatal(http.ListenAndServe(":5000", r))
}
func api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := sqlx.Connect("postgres", "host=localhost  dbname=postgres password=shitter sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	switch r.URL.String() {
	case "/api/get":
		fmt.Println("is this running?")
		rows, err := db.Queryx("SELECT * FROM todos")
		if err != nil {
			fmt.Println(err)
		}
		var todos []Todo
		for rows.Next() {
			todo := Todo{}
			err := rows.StructScan(&todo)
			if err != nil {
				fmt.Println(err)
			}
			todos = append(todos, todo)
		}
		rz, _ := json.Marshal(todos)
		w.Write(rz)
	}
}
func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("users running")
	db, err := sqlx.Connect("postgres", "host=localhost  dbname=postgres password=shitter sslmode=disable")
	if err != nil {
		fmt.Println("error..")
		log.Fatalln(err)
	}
	defer db.Close()
	switch r.URL.String() {
	case "/users/get":
		var users []User
		rows, err := db.Queryx("SELECT * FROM users")
		if err != nil {
			fmt.Println(err)
		}
		for rows.Next() {
			user := User{}
			err := rows.StructScan(&user)
			if err != nil {
				fmt.Println(err)
			}
			users = append(users, user)
		}
		rs, _ := json.Marshal(users)
		w.Write(rs)
	case "/users/create":
		user := &ReceivedUser{}
		ch := make(chan int)
		go func() {
			var checker EmailCheck
			json.NewDecoder(r.Body).Decode(&user)
			defer r.Body.Close()
			err := db.Get(&checker, `SELECT email, COUNT(email) FROM users WHERE email=$1 GROUP BY email `, user.Email)
			if err != nil {
				fmt.Println(err)
			}
			rz, err := strconv.Atoi(checker.Count)
			ch <- rz
		}()
		rq := <-ch
		if rq > 0 {
			invalid := Error{Error: "Email already in use"}
			rz, _ := json.Marshal(invalid)
			w.Write(rz)
			return
		}
		fmt.Println("it checks out captain")
	}
}
