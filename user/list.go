package user

import (
	"encoding/json"
	"go-api-crud/db"
	"go-api-crud/structs"
	"net/http"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	db := db.GetDB()
	var users []structs.User
	result := db.Find(&users)
	if result.Error != nil {
		http.Error(w, "Failed to get users", 400)
		return
	}

	usersMap := map[string][]structs.User{
		"users": users,
	}

	json.NewEncoder(w).Encode(usersMap)
}
