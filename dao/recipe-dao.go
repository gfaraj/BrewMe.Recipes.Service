package dao

import (
	"github.com/gfaraj/brewme.recipes.service/model"
)

// RecipeDao represents an abstract data repository for recipes.
type RecipeDao interface {
	// GetRecipes gets all recipes.
	GetRecipes() ([]model.Recipe, error)
	// GetRecipe gets the recipe with the specified id (if any).
	GetRecipe(id string) (*model.Recipe, error)
	// UpsertRecipe adds or updates a recipe.
	UpsertRecipe(r *model.Recipe) error
	// DeleteRecipe delets the recipe with the specified id (if any).
	DeleteRecipe(id string) error
}
