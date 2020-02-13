package dao

import (
	"errors"

	"github.com/gfaraj/brewme.recipes.service/model"
	"github.com/google/uuid"
)

// RecipeDaoMock is a mock recipe repository.
type RecipeDaoMock struct {
}

var mockData = []model.Recipe{
	model.Recipe{
		ID:          "bfc0fa32-f440-4189-bec0-8a457c5cfd3b",
		Name:        "Razorback",
		Description: "This is a smooth dark beer."},
	model.Recipe{
		ID:          "6b530eb4-1219-417f-a5f7-558589d78c43",
		Name:        "Wicked",
		Description: "Delicious amber beer."}}

// NewRecipeDaoMock creates a new mock recipe dao.
func NewRecipeDaoMock() RecipeDao {
	return new(RecipeDaoMock)
}

// GetRecipes gets all recipes.
func (d *RecipeDaoMock) GetRecipes() ([]model.Recipe, error) {
	return mockData, nil
}

// GetRecipe gets the recipe with specified id (if any).
func (d *RecipeDaoMock) GetRecipe(id string) (*model.Recipe, error) {
	index := -1
	for i, v := range mockData {
		if v.ID == id {
			index = i
		}
	}
	if index < 0 {
		return nil, errors.New("recipe does not exist: " + id)
	}
	return &mockData[index], nil
}

// UpsertRecipe adds or updates a receipe.
func (d *RecipeDaoMock) UpsertRecipe(r *model.Recipe) error {
	index := -1
	for i, v := range mockData {
		if v.ID == r.ID {
			index = i
		}
	}
	if index < 0 {
		r.ID = uuid.New().String()
		mockData = append(mockData, *r)
	} else {
		mockData[index] = *r
	}

	return nil
}

// DeleteRecipe deletes the recipe with specified id (if any).
func (d *RecipeDaoMock) DeleteRecipe(id string) error {
	index := -1
	for i, v := range mockData {
		if v.ID == id {
			index = i
		}
	}
	if index < 0 {
		return errors.New("recipe does not exist: " + id)
	}
	mockData = mockData[:index+copy(mockData[index:], mockData[index+1:])]
	return nil
}
