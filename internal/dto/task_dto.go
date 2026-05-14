package dto

type TaskResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

func ToTaskResponse(id int, name string, done bool) TaskResponse {
	return TaskResponse{
		ID:   id,
		Name: name,
		Done: done,
	}
}
