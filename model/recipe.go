package model

// Recipe is the model used to store and return from the service.
type Recipe struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
