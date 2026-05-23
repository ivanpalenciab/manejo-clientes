package repository

import "my-api/internal/database"

func createUser() {
	database.ConnectDB()

}
