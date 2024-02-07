package api

import (
	"go-api-crud/user"
	"net/http"

	"github.com/gorilla/mux"
)

func JsonMiddleware(r *mux.Router) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, req)
		})
	}
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(JsonMiddleware(router))
	router.HandleFunc("/user", user.CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", user.GetUser).Methods("GET")
	router.HandleFunc("/users", user.ListUsers).Methods("GET")
	router.HandleFunc("/user/{id}", user.DeleteUser).Methods("DELETE")
	return router
}
