package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func (r *CreateTodoRequest) Validate() error {
	if r.Title == "" && r.Description == "" {
		return fmt.Errorf("request body is malformed or empty")
	}
	if r.Title == "" {
		return errParamIsRequired("title", "string")
	}
	if r.Description == "" {
		return errParamIsRequired("description", "string")
	}
	return nil
}

type UpdateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func (r *UpdateTodoRequest) Validate() error {
	if r.Title != "" || r.Description != "" {
		return nil
	}
	return fmt.Errorf("at least one of the following fields must be present: title, description")
}
