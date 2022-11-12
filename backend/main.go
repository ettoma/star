package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ettoma/star/database"
	"github.com/ettoma/star/handles"
	"github.com/ettoma/star/models"
	"github.com/ettoma/star/utils"
	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

const PORT = ":8000"

var users = []models.User{
	{
		Name:     "Borat",
		Username: "mrborat",
	},
	{
		Name:     "Joaquin Phoenix",
		Username: "totallyJoker",
	},
}

func getUserData(username string) (models.User, error) {
	for _, user := range users {
		if user.Username == username {
			return user, nil
		}
	}
	errMsg := fmt.Sprintf("user: %s not found", username)
	return models.User{}, errors.New(errMsg)

}

func submitReview(sender, receiver, content string) {
	review := models.Review{
		Content:   content,
		Timestamp: time.Now().Unix(),
	}

	r, err := getUserData(receiver)
	utils.HandleFatal(err)

	s, err := getUserData(sender)
	utils.HandleFatal(err)

	review.Receiver = r
	review.Sender = s

	fmt.Println(review)

}

func addUser(name, username string) error {
	newUser := models.User{
		Name:     name,
		Username: username,
	}

	for _, u := range users {
		if u.Username == newUser.Username {
			errMsg := fmt.Sprintf("Username: %s already exists", newUser.Username)
			return errors.New(errMsg)
		}
	}

	users = append(users, newUser)
	return nil
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n Url: %s \n Method: %s \n Content-length: %d \n", r.URL, r.Method, r.ContentLength)
		next.ServeHTTP(w, r)
	})

}

func main() {
	database.OpenDb()

	r := mux.NewRouter()
	srv := &http.Server{
		Addr:         PORT,
		Handler:      r,
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
	}

	r.HandleFunc("/", handles.Home).Methods("GET")
	r.HandleFunc("/users", handles.GetAllUsers).Methods("GET")
	r.HandleFunc("/users", handles.AddUser).Methods("POST")
	r.HandleFunc("/users", handles.DeleteUserByUsername).Methods("DELETE") // TODO: merge deleteById and deleteByUsername

	r.Use(loggingMiddleware)

	log.Fatal(srv.ListenAndServe())

}
