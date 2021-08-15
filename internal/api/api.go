package api

import(
  "net/http"
  "time"
  "encoding/json"
)

type Api struct{

}

var running *Pomo = nil

func (api *Api) StartPomo(w http.ResponseWriter, r *http.Request) {
  pomo := &Pomo{
    StartedAt: time.Now(),
    StoppedAt: time.Now(),
    Status: PomoRunning,
  }
  running = pomo
  json.NewEncoder(w).Encode(pomo)
}

func (api *Api) StopPomo(w http.ResponseWriter, r *http.Request) {
  if running == nil {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte("No running pomo"))
    return
  }

  running.StoppedAt = time.Now()
  running.DurationSeconds = int(running.StoppedAt.Sub(running.StartedAt))/1000/1000/1000
  json.NewEncoder(w).Encode(running)
  running = nil
}
