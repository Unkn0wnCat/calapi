package model

type Calendar struct {
	ID          string `json:"id"`
	DbID        uint64
	Name        string `json:"name"`
	Description string `json:"description"`
}
