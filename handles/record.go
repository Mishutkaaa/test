package handles

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"test/model"

	"github.com/jackc/pgx/pgtype"
)

func GetByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var record model.Record

		if r.Method != "GET" {
			http.Error(w, "method not allowed ", http.StatusMethodNotAllowed)
		}

		id := r.PathValue("id")
		row := db.QueryRow("select user_id, service_name, price, to_char(start_date, 'MM-YYYY') AS start_date from record where id = $1", id)
		if err := row.Scan(&record.UserID, &record.ServiceName, &record.Price, &record.StartDate); err != nil {
			log.Println("cannot scan row", err)
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
		}

		rows, err := db.Query("select user_id, service_name, price, to_char(start_date, 'MM-YYYY') AS start_date from record")
		if err != nil {
			log.Println("cannot get rows", err)
			return
		}
		defer rows.Close()
		for rows.Next() {
			record := model.Record{}
			if err := rows.Scan(&record.UserID, &record.ServiceName, &record.Price, &record.StartDate); err != nil {
				log.Println("cannot scan row", err)
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
		}

		if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
			http.Error(w, "error decode json", http.StatusBadRequest)
			log.Println(err)
			return
		}

		if record.UserID == (pgtype.UUID{}) || record.ServiceName == "" || record.Price == 0 || record.StartDate == "" {
			http.Error(w, "please fill in all required fields", http.StatusBadRequest)
			return
		}

		if _, err := db.Exec("insert into record (user_id, service_name, price, start_date) values($1, $2, $3, $4)", record.UserID, record.ServiceName, record.Price, record.StartDate); err != nil {
			log.Println("cannot create record", record)
			return
		}

	}
}

func DeleteRecord(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "DELETE" {
			http.Error(w, "method not allowed ", http.StatusMethodNotAllowed)
		}

		id := r.PathValue("id")

		if _, err := db.Exec("delete from record where id = $1", id); err != nil {
			log.Println("cannot delete record", err)
			return
		}
	}
}
