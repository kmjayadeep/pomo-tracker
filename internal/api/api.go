package api

import (
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type Api struct {
	db *gorm.DB
}

func NewApi(db *gorm.DB) *Api {
	db.AutoMigrate(&Pomo{})
	return &Api{
		db: db,
	}
}

var running *Pomo = nil

func (api *Api) StartPomo(w http.ResponseWriter, r *http.Request) {
	pomo := &Pomo{
		StartedAt: time.Now(),
		StoppedAt: time.Now(),
		Status:    PomoRunning,
	}
	api.db.Create(pomo)
	running = pomo
	json.NewEncoder(w).Encode(pomo)
}

func (api *Api) TodayStats(w http.ResponseWriter, r *http.Request) {
	l := []Pomo{}
	api.db.Where("DATE(stopped_at) = CURDATE()").Find(&l)

	dur := 0
	for _, p := range l {
		dur = dur + p.DurationSeconds
	}

	res := struct {
		TotalPomos           int `json:"totalPomos"`
		TotalDurationSeconds int `json:"totalDurationSeconds"`
	}{
		TotalPomos:           len(l),
		TotalDurationSeconds: dur,
	}

	json.NewEncoder(w).Encode(res)
}

func (api *Api) StopPomo(w http.ResponseWriter, r *http.Request) {
	if running == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No running pomo"))
		return
	}

	duration := int(time.Now().Sub(running.StartedAt)) / 1000 / 1000 / 1000

	api.db.Model(running).Updates(Pomo{
		Status:          PomoStopped,
		StoppedAt:       time.Now(),
		DurationSeconds: duration,
	})
	json.NewEncoder(w).Encode(running)
	running = nil
}
