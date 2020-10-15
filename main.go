package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"zagorceva-klet/handlers"
)

func main()  {
	//db, err := dbInit()

	router := mux.NewRouter()
	router.HandleFunc("/phrase", handlers.PhraseGet).Methods("GET")
	router.HandleFunc("/phrase", handlers.PhrasePost).Methods("POST")

	fmt.Println("Dobro mi došel prijatelj, vu staru zagorsku klet.")
	http.ListenAndServe(":8891", router)
}