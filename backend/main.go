package main

import (
	"net/http"
	"time"

	"github.com/ettoma/star/auth"
	"github.com/ettoma/star/database"
	"github.com/ettoma/star/handles"
	mdw "github.com/ettoma/star/middlewares"
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

	r.Use(mdw.Cors)
	r.Use(mdw.LoggerMiddleware)
	r.Use(mdw.ContentTypeMiddleware)

	//? These are the routes that require authentication
	//?----------------------------------------------------------------
	authR := r.PathPrefix("/kudos").Subrouter()
	authR.Use(mdw.TokenValidationMiddleware)
	authR.HandleFunc("", handles.GetAllKudos).Methods("GET")
	authR.HandleFunc("", handles.AddKudos).Methods("POST", "OPTIONS")
	authR.HandleFunc("", handles.DeleteKudos).Methods("DELETE", "OPTIONS")
	authR.HandleFunc("/{receiver}", handles.GetKudosPerUser).Methods("GET", "OPTIONS")

	tokensR := r.PathPrefix("/auth-refresh").Subrouter()
	tokensR.Use(mdw.TokenValidationMiddleware)
	tokensR.HandleFunc("", auth.RefreshJWT).Methods("POST", "OPTIONS")
	//?----------------------------------------------------------------

	//? These routes are public and don't require token validation
	//?----------------------------------------------------------------
	r.HandleFunc("/", handles.Home).Methods("GET")

	r.HandleFunc("/users", handles.GetAllUsers).Methods("GET", "OPTIONS")
	r.HandleFunc("/users/user", handles.GetUserById).Methods("GET")
	r.HandleFunc("/users", handles.DeleteUser).Methods("DELETE", "OPTIONS")

	r.HandleFunc("/auth-issuer", auth.GenerateJWT).Methods("POST")
	r.HandleFunc("/auth", auth.ValidateJWT).Methods("POST")

	r.HandleFunc("/register", handles.RegisterUser).Methods("POST")
	r.HandleFunc("/login", handles.Login).Methods("POST")

	//TODO: implement Patch user handle
	// r.HandleFunc("/users", handles.PatchUser).Methods("PATCH")
	//?----------------------------------------------------------------

	utils.HandleFatal(srv.ListenAndServe())

}
