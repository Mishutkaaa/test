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
		id := r.PathValue("id")
		row := db.QueryRow("select user_id, service_name, price, start_date from record where id = $1", id)
		if err := row.Scan(&record.UserID, &record.ServiceName, &record.Price, &record.StartDate); err != nil {
			log.Println("cannot scan row")
		}
		log.Println(record)
	}
}

func GetAll(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var records []model.Record
		rows, err := db.Query("select user_id, service_name, price, start_date from record")
		if err != nil {
			log.Println("cannot get rows", err)
			return
		}
		defer rows.Close()
		for rows.Next() {
			record := model.Record{}
			if err := rows.Scan(&record.UserID, &record.ServiceName, &record.Price, &record.StartDate); err != nil {
				log.Println("cannot scan row")
			}
			records = append(records, record)
		}

		log.Println(records)
	}
}

func CreateRecord(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var record model.Record

		if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
			http.Error(w, "error decode json", http.StatusBadRequest)
			log.Println(err)
			return
		}

		if record.UserID == (pgtype.UUID{}) || record.ServiceName == "" || record.Price == 0 || record.StartDate.IsZero() {
			http.Error(w, "please fill in all required fields", http.StatusBadRequest)
			return
		}

		if _, err := db.Exec("insert into record (user_id, service_name, price, start_date) values($1, $2, $3, $4)", record.UserID, record.ServiceName, record.Price, record.StartDate); err != nil {
			log.Println("cannot create record", record)
			return
		}

	}
}
