package api

import(
  "net/http"
  "time"
  "encoding/json"
  "gorm.io/gorm"
)

type Api struct{
  db *gorm.DB
}

func NewApi(db *gorm.DB) *Api{
  db.AutoMigrate(&Pomo{})
  return &Api{
    db : db,
  }
}

var running *Pomo = nil

func (api *Api) StartPomo(w http.ResponseWriter, r *http.Request) {
  pomo := &Pomo{
    StartedAt: time.Now(),
    StoppedAt: time.Now(),
    Status: PomoRunning,
  }
  api.db.Create(pomo)
  running = pomo
  json.NewEncoder(w).Encode(pomo)
}

func (api *Api) StopPomo(w http.ResponseWriter, r *http.Request) {
  if running == nil {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte("No running pomo"))
    return
  }

  duration := int(time.Now().Sub(running.StartedAt))/1000/1000/1000

  api.db.Model(running).Updates(Pomo{
    StoppedAt: time.Now(),
    DurationSeconds : duration,
  })
  json.NewEncoder(w).Encode(running)
  running = nil
}
