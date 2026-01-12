package handles

import (
	"database/sql"
	"log"
	"net/http"
	"test/model"
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
