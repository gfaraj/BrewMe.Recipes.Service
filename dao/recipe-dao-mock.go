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
		ID:          "12324245242",
		Name:        "Razorback",
		Description: "This is a smooth dark beer."},
	model.Recipe{
		ID:          "57474456456",
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
	return &mockData[0], nil
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
		return errors.New("update not implemented yet")
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
		return errors.New("recipe does not exist")
	}
	mockData = mockData[:index+copy(mockData[index:], mockData[index+1:])]
	return nil
}
