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
	http.HandleFunc("/delete/{id}", handles.DeleteRecord(db))
	http.HandleFunc("/sum", handles.GetSumm(db))

	http.ListenAndServe(":8080", nil)
}
