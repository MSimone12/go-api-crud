package user

import (
	"encoding/json"
	"fmt"
	"go-api-crud/db"
	"go-api-crud/structs"
	"net/http"
)

const usersDbName = "users"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := structs.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "body does not satisfies User constraints", 400)
		return
	}
	db := db.GetDB()

	result := db.Create(&user)
	if result.Error != nil {
		http.Error(w, fmt.Sprintf("Error Creating user: %v", result.Error), 400)
		return
	}

	json.NewEncoder(w).Encode(user)
}
