package main

import(
  "log"
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/kmjayadeep/pomo-tracker/internal/api"
)

func main() {
  router := mux.NewRouter()

  a := api.Api{}

  router.HandleFunc("/pomo/start", a.StartPomo).Methods(http.MethodGet)
  router.HandleFunc("/pomo/stop", a.StopPomo).Methods(http.MethodGet)

  fmt.Println("Starting server on port 8080")
  log.Fatal(http.ListenAndServe(":8080", router))
}
