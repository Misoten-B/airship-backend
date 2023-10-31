package main

import "github.com/Misoten-B/airship-backend/internal/database"

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}

	if err = db.AutoMigrate(&database.User{}); err != nil {
		panic(err)
	}
}
