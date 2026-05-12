package models

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `josn:"done"`
}
