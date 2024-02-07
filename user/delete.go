package user

import (
	"encoding/json"
	"go-api-crud/db"
	"go-api-crud/structs"
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	db := db.GetDB()

	user := structs.User{}

	response := db.First(&user, params["id"])

	if response.Error != nil || user == (structs.User{}) {
		http.Error(w, "Error getting user of id "+params["id"], 400)
		return
	}

	tx := db.Delete(&user)

	tx.Commit()

	json.NewEncoder(w).Encode(map[string]string{})
}
