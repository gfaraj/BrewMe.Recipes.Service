package main

import (
	"fmt"
	"net/http"

	"github.com/gfaraj/brewme.recipes.service/web"

	"github.com/urfave/negroni"
)

const (
	port = 3100
)

func main() {
	router := web.NewRouter()

	controller := InitializeRecipeController()
	controller.SetupRoutes(router)

	n := negroni.Classic()
	n.UseHandler(router)

	fmt.Println(fmt.Sprintf("Listening on port %d", port))

	http.ListenAndServe(fmt.Sprintf(":%d", port), n)
}
