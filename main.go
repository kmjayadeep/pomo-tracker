package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kmjayadeep/pomo-tracker/internal/api"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {

	dsn := "user:password@tcp(127.0.0.1:3306)/db?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err, "unable to connect to db")
	}

	router := mux.NewRouter()

	a := api.NewApi(db)

	router.HandleFunc("/pomo/start", a.StartPomo).Methods(http.MethodGet)
	router.HandleFunc("/pomo/stop", a.StopPomo).Methods(http.MethodGet)
	router.HandleFunc("/pomo/stats/today", a.TodayStats).Methods(http.MethodGet)

	fmt.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
