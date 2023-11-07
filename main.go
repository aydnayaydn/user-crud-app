package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var db *sql.DB

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func main() {
	// MSSQL veritabanı bağlantısını başlat
	var err error
	db, err = InitDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := NewSQLUserRepository(db)

	r := http.NewServeMux()
	r.HandleFunc("/create", createUserHandler(repo))
	r.HandleFunc("/read", readUserHandler(repo))
	r.HandleFunc("/update", updateUserHandler(repo))
	r.HandleFunc("/delete", deleteUserHandler(repo))

	fmt.Println("Server listening on port 4400...")
	http.ListenAndServe(":4400", r)
}
