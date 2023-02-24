package model

import (
	"github.com/Unkn0wnCat/calapi/internal/db_model"
	"time"
)

type Event struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Location    *Location `json:"location"`
	Description string    `json:"description"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
	DbCalendar  *db_model.Calendar
}
