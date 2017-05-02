package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/baptistelambert/go-api-boilerplate/app/router"
	"github.com/baptistelambert/go-api-boilerplate/config"
	"github.com/julienschmidt/httprouter"
)

// App contain our config and our router
type App struct {
	Router *httprouter.Router
	Config *config.Config
}

// Init the app
func (app *App) Init() {
	app.Router = router.GetRouter()
	app.Config = config.LoadConfig()
}

// Start our application
func (app *App) Start() {
	addr := app.Config.Host + ":" + app.Config.Port

	fmt.Println("API available at " + addr)
	log.Fatal(http.ListenAndServe(addr, app.Router))
}
