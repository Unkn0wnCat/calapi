package db_model

import "time"

//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen

type Event struct {
	Id          uint64
	Title       string
	Description string
	Calendar    *Calendar `objectbox:"link"`

	LocationLat  float64
	LocationLon  float64
	LocationName string
	LocationAddr string

	Start       time.Time `objectbox:"date,index"`
	End         time.Time `objectbox:"date,index"`
	DateCreated time.Time `objectbox:"date"`
}
