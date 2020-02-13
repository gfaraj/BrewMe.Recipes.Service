package web

import (
	"encoding/json"
	"net/http"

	"github.com/gfaraj/brewme.recipes.service/dao"
	"github.com/gfaraj/brewme.recipes.service/model"
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
	router.HandleFunc("/api/recipes", c.GetRecipes).Methods("GET")
	router.HandleFunc("/api/recipes/{id}", c.GetRecipe).Methods("GET")
	router.HandleFunc("/api/recipes", c.CreateRecipe).Methods("POST")
	router.HandleFunc("/api/recipes", c.UpdateRecipe).Methods("PUT")
	router.HandleFunc("/api/recipes/{id}", c.DeleteRecipe).Methods("DELETE")
}

// GetRecipes handles the /recipes GET request.
func (c *RecipeController) GetRecipes(w http.ResponseWriter, r *http.Request) {
	recipes, err := c.recipeDao.GetRecipes()
	if err != nil {
		// TODO: log error
		sendJSON(w, http.StatusInternalServerError, "Error while retrieving recipes")
		return
	}

	sendJSONOk(w, recipes)
}

// GetRecipe handles the /recipes/{recipeId} GET request.
func (c *RecipeController) GetRecipe(w http.ResponseWriter, r *http.Request) {
	recipeID := r.URL.Query().Get("id")
	recipe, err := c.recipeDao.GetRecipe(recipeID)
	if err != nil {
		// TODO: log error
		sendJSON(w, http.StatusInternalServerError, "Error while retrieving recipe")
		return
	}

	sendJSONOk(w, recipe)
}

// CreateRecipe handles the /recipes POST request.
func (c *RecipeController) CreateRecipe(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	recipe := &model.Recipe{}
	json.NewDecoder(r.Body).Decode(recipe)

	c.recipeDao.UpsertRecipe(recipe)
	sendJSONOk(w, recipe)
}

// UpdateRecipe handles the /recipes PUT request.
func (c *RecipeController) UpdateRecipe(w http.ResponseWriter, r *http.Request) {

}

// DeleteRecipe handles the /recipe/{recipeId} DELETE request.
func (c *RecipeController) DeleteRecipe(w http.ResponseWriter, r *http.Request) {

}

func sendJSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)

	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			panic(err)
		}
	}
}

func sendJSONOk(w http.ResponseWriter, data interface{}) {
	sendJSON(w, http.StatusOK, data)
}
