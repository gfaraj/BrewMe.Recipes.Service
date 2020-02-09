package dao

import (
	"../model"
)

// RecipeDaoMock is a mock recipe repository.
type RecipeDaoMock struct {
}

// GetRecipes gets all recipes.
func (d *RecipeDaoMock) GetRecipes() ([]model.Recipe, error) {
	return nil, nil
}

// GetRecipe gets the recipe with specified id (if any).
func (d *RecipeDaoMock) GetRecipe(id string) (*model.Recipe, error) {
	return nil, nil
}

// UpsertRecipe adds or updates a receipe.
func (d *RecipeDaoMock) UpsertRecipe(r *model.Recipe) error {
	return nil
}

// DeleteRecipe deletes the recipe with specified id (if any).
func (d *RecipeDaoMock) DeleteRecipe(id string) error {
	return nil
}
