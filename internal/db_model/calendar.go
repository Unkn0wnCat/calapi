package db_model

import "time"

//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen

type Calendar struct {
	Id          uint64
	Name        string
	Description string

	DateCreated time.Time `objectbox:"date"`
}
