package web

import (
	"net/http"

	"../dao"
	"github.com/gorilla/mux"
)

// RecipeController is the web endpoint controller for recipes.
type RecipeController struct {
	recipeDao dao.RecipeDao
}

// NewRecipeController creates a new recipe controller.
func NewRecipeController(d dao.RecipeDao) *RecipeController {
	controller := RecipeController{recipeDao: d}
	return &controller
}

// SetupRoutes sets up routes for the controller.
func (c *RecipeController) SetupRoutes(router *mux.Router) {
	router.HandleFunc("/recipes", c.GetRecipes).Methods("GET")
	router.HandleFunc("/recipes/{recipeId}", c.GetRecipe).Methods("GET")
	router.HandleFunc("/recipes", c.AddRecipe).Methods("POST")
	router.HandleFunc("/recipes", c.UpdateRecipe).Methods("PUT")
	router.HandleFunc("/recipes/{recipeId}", c.DeleteRecipe).Methods("DELETE")
}

// GetRecipes handles the /recipes GET request.
func (c *RecipeController) GetRecipes(w http.ResponseWriter, r *http.Request) {

}

// GetRecipe handles the /recipes/{recipeId} GET request.
func (c *RecipeController) GetRecipe(w http.ResponseWriter, r *http.Request) {

}

// AddRecipe handles the /recipes POST request.
func (c *RecipeController) AddRecipe(w http.ResponseWriter, r *http.Request) {

}

// UpdateRecipe handles the /recipes PUT request.
func (c *RecipeController) UpdateRecipe(w http.ResponseWriter, r *http.Request) {

}

// DeleteRecipe handles the /recipe/{recipeId} DELETE request.
func (c *RecipeController) DeleteRecipe(w http.ResponseWriter, r *http.Request) {

}
