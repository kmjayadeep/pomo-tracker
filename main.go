package main

import(
  "log"
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/kmjayadeep/pomo-tracker/internal/api"
  "gorm.io/gorm"
  "gorm.io/driver/mysql"
)

func main() {

  dsn := "user:password@tcp(127.0.0.1:3306)/db"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

  if err != nil {
    log.Fatal(err, "unable to connect to db")
  }

  router := mux.NewRouter()

  a := api.NewApi(db)

  router.HandleFunc("/pomo/start", a.StartPomo).Methods(http.MethodGet)
  router.HandleFunc("/pomo/stop", a.StopPomo).Methods(http.MethodGet)

  fmt.Println("Starting server on port 8080")
  log.Fatal(http.ListenAndServe(":8080", router))
}
