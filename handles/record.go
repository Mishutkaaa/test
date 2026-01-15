package handles

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"test/model"
)

func GetByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var record model.Record

		if r.Method != "GET" {
			http.Error(w, "method not allowed ", http.StatusMethodNotAllowed)
			return
		}

		id := r.PathValue("id")
		row := db.QueryRow("select user_id, service_name, price, to_char(start_date, 'MM-YYYY') AS start_date from record where id = $1", id)
		if err := row.Scan(&record.UserID, &record.ServiceName, &record.Price, &record.StartDate); err != nil {
			http.Error(w, "cannot scan row", http.StatusInternalServerError)
			log.Println(err)
		}
		if record == (model.Record{}) {
			http.Error(w, "record not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(record)
	}
}

func GetAll(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var records []model.Record

		if r.Method != "GET" {
			http.Error(w, "method not allowed ", http.StatusMethodNotAllowed)
			return
		}

		rows, err := db.Query("select user_id, service_name, price, to_char(start_date, 'MM-YYYY') AS start_date from record")
		if err != nil {
			http.Error(w, "cannot get rows", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		defer rows.Close()
		for rows.Next() {
			record := model.Record{}
			if err := rows.Scan(&record.UserID, &record.ServiceName, &record.Price, &record.StartDate); err != nil {
				http.Error(w, "cannot scan rows", http.StatusInternalServerError)
				log.Println(err)
			}
			records = append(records, record)
		}

		json.NewEncoder(w).Encode(records)
	}
}

func CreateRecord(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var record model.Record

		if r.Method != "POST" {
			http.Error(w, "method not allowed ", http.StatusMethodNotAllowed)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
			http.Error(w, "error decode json", http.StatusBadRequest)
			log.Println(err)
			return
		}

		if record.UserID == "" || record.ServiceName == "" || record.Price == 0 || record.StartDate == "" {
			http.Error(w, "please fill in all required fields", http.StatusBadRequest)
			return
		}

		record.StartDate = "01-" + record.StartDate

		if _, err := db.Exec("insert into record (user_id, service_name, price,  start_date) values($1, $2, $3, $4)", record.UserID, record.ServiceName, record.Price, record.StartDate); err != nil {
			http.Error(w, "cannot create record", http.StatusInternalServerError)
			log.Println(err)
			return
		}

	}
}

func DeleteRecord(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "DELETE" {
			http.Error(w, "method not allowed ", http.StatusMethodNotAllowed)
			return
		}

		id := r.PathValue("id")

		if _, err := db.Exec("delete from record where id = $1", id); err != nil {
			http.Error(w, "cannot delete record", http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}
}

func GetSumm(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "method not allowed ", http.StatusMethodNotAllowed)
			return
		}

		dateTo := r.URL.Query().Get("dateTo")
		dateFrom := r.URL.Query().Get("dateFrom")
		userID := r.URL.Query().Get("user")
		serviceName := r.URL.Query().Get("name")

		var record model.Record
		var query = fmt.Sprintf("select sum(price) as price from record where to_char(start_date, 'MM-YYYY') >= '%s' and to_char(start_date, 'MM-YYYY') <= '%s' ", dateTo, dateFrom)

		if dateTo == "" || dateFrom == "" {
			http.Error(w, "pleale fill dateTo and dateFrom", http.StatusBadRequest)
			return
		}

		if userID != "" {
			query += fmt.Sprintf("and user_id = '%s' ", userID)
		}

		if serviceName != "" {
			query += fmt.Sprintf("and service_name = '%s' ", serviceName)
		}

		row := db.QueryRow(query)
		if err := row.Scan(&record.Price); err != nil {
			log.Println("cannot scan row", err)
		}
		if record == (model.Record{}) {
			http.Error(w, "record not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(&record.Price)
	}
}
