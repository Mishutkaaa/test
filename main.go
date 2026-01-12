package main

import (
	"net/http"
	"test/database"
	"test/handles"
)

func main() {
	db := database.Conn()
	defer db.Close()

	http.HandleFunc("/get/{id}", handles.GetByID(db))
	http.HandleFunc("/get", handles.GetAll(db))
	http.HandleFunc("/create", handles.CreateRecord(db))
	http.HandleFunc("/update", handles.DeleteRecord(db))

	http.ListenAndServe(":8080", nil)
}
