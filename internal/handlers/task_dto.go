package handlers

type CreateTaskRequest struct {
	Name string `json:"name"`
	Done bool   `json:"done"`
}
