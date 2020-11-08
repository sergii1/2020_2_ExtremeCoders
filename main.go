package main

import (
	"CleanArch/app/Delivery"
	"CleanArch/app/Repository/Postgres"
	"CleanArch/app/UseCase"
	"fmt"
	"github.com/rs/cors"
	"net/http"
	"time"
)



func main() {
	var db=Postgres.DataBase{ User: "postgres", Password: "1538", DataBaseName: "maila"}
	db.Init()
	var uc=UseCase.UseCase{Db: db}
	var yaFood=Delivery.Delivery{Uc: uc}
	mux := http.NewServeMux()
	mux.HandleFunc("/signup", yaFood.Signup)
	mux.HandleFunc("/signin", yaFood.SignIn)
	mux.HandleFunc("/profile", yaFood.Profile)
	mux.HandleFunc("/logout", yaFood.Logout)
	mux.HandleFunc("/getAvatar", yaFood.GetAvatar)
	mux.HandleFunc("/sendMessage", yaFood.SendLetter)
	mux.HandleFunc("/getRecvLetters", yaFood.GetRecvLetters)
	mux.HandleFunc("/getSendLetters", yaFood.GetSendLetters)
	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "http://127.0.0.1:3000", "http://95.163.209.195:3000"},
		AllowedHeaders: []string{"Version", "Authorization", "Content-Type"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowCredentials: true,
	}).Handler(mux)
	server := http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("starting server at :8080")
	server.ListenAndServe()
}
