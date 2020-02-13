// +build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/gfaraj/brewme.recipes.service/dao"
	"github.com/gfaraj/brewme.recipes.service/web"
)

func InitializeRecipeDao() dao.RecipeDao {
	panic(wire.Build(dao.NewRecipeDaoMock))
}

func InitializeRecipeController() *web.RecipeController {
	panic(wire.Build(web.NewRecipeController, InitializeRecipeDao))
}
