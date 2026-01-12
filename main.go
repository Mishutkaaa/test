package main

import (
	"test/database"
)

func main() {
	db := database.Conn()
	defer db.Close()

}
