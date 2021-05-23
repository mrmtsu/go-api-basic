package app

import (
	"fmt"
	"go-api-basic/domain"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func sanityCheck() {
	godotenv.Load("envfiles/.env")

	if os.Getenv("DB_ADDRESS") == "" ||
		os.Getenv("DB_PORT") == "" {
		log.Fatal("Environment variable not defined...")
	}
}

func Start() {
	sanityCheck()
	router := mux.NewRouter()

	dbConnect()

	// define routes
	router.HandleFunc("/", Hello)
	router.HandleFunc("/api/users", GetAllUsers).Methods("GET")
	router.HandleFunc("/api/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/api/users", CreateUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/api/users/{id}", DeleteUser).Methods("DELETE")
	cors := cors.Default().Handler(router)

	// starting server
	http.ListenAndServe(":8000", cors)
}

var DB *gorm.DB

func dbConnect() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database_name := os.Getenv("DB_DATABASE_NAME")

	dns := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database_name + "?charset=utf8mb4"
	database, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}

	DB = database

	database.AutoMigrate(&domain.User{})
}
