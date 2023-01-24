package main

import (
	"net/http"
	"time"

	"github.com/ettoma/star/auth"
	"github.com/ettoma/star/database"
	"github.com/ettoma/star/handles"
	"github.com/ettoma/star/utils"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const PORT = ":8000"

func main() {
	database.DbInit()

	r := mux.NewRouter().StrictSlash(false)

	srv := &http.Server{
		Addr:         PORT,
		Handler:      r,
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
	}

	r.Use(utils.Cors)
	r.Use(utils.LoggerMiddleware)
	r.Use(utils.ContentTypeMiddleware)

	authR := r.PathPrefix("/kudos").Subrouter()
	authR.Use(utils.TokenValidationMiddleware)
	authR.HandleFunc("/users&r={receiver}", handles.GetKudosPerUser).Methods("GET")

	// authRouter := r.PathPrefix("/kudos/users").Subrouter()
	// authRouter.Use(utils.LoggerMiddleware)
	// authRouter.Use(utils.Cors)
	// authRouter.HandleFunc("&r={receiver}", handles.GetKudosPerUser).Methods("GET")

	r.HandleFunc("/", handles.Home).Methods("GET")

	r.HandleFunc("/users", handles.GetAllUsers).Methods("GET", "OPTIONS")
	r.HandleFunc("/users/user", handles.GetUserById).Methods("GET")
	r.HandleFunc("/users", handles.DeleteUser).Methods("DELETE", "OPTIONS")
	//TODO: implement Patch user handle
	// r.HandleFunc("/users", handles.PatchUser).Methods("PATCH")

	r.HandleFunc("/kudos", handles.GetAllKudos).Methods("GET")
	r.HandleFunc("/kudos", handles.AddKudos).Methods("POST")
	r.HandleFunc("/kudos", handles.DeleteKudos).Methods("DELETE")

	r.HandleFunc("/auth-issuer", auth.GenerateJWT).Methods("POST")
	r.HandleFunc("/auth", auth.ValidateJWT).Methods("POST")

	r.HandleFunc("/register", handles.RegisterUser).Methods("POST")
	r.HandleFunc("/login", handles.Login).Methods("POST")

	//! TODO: token validation should only be triggered for user-specific endpoints
	// r.HandleFunc("/kudos/users&r={receiver}", handles.GetKudosPerUser).Methods("GET")

	utils.HandleFatal(srv.ListenAndServe())

}
