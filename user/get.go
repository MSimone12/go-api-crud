package user

import (
	"encoding/json"
	"go-api-crud/db"
	"go-api-crud/structs"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	user := structs.User{}

	db := db.GetDB()

	params := mux.Vars(r)

	result := db.First(&user, params["id"])

	if result.Error != nil {
		http.Error(w, "Error when fetching user", 400)
		return
	}

	json.NewEncoder(w).Encode(user)
}
