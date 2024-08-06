package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"githubcom/kosatnkn/web-page-analyzer-api/app/container"
	"githubcom/kosatnkn/web-page-analyzer-api/transport/http/controllers"
	"githubcom/kosatnkn/web-page-analyzer-api/transport/http/middleware"
)

// Init initializes the router.
func Init(ctr *container.Container) *mux.Router {
	// create new router
	r := mux.NewRouter()

	// add middleware to the router
	//
	// NOTE: middleware will execute in the order they are added to the router
	// add metrics middleware first
	r.Use(middleware.NewMetricsMiddleware().Middleware)
	// enable CORS
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(middleware.NewCORSMiddleware().Middleware)
	// other middleware
	r.Use(middleware.NewRequestAlterMiddleware(ctr).Middleware)
	r.Use(middleware.NewRequestCheckerMiddleware(ctr).Middleware)

	// initialize controllers
	apiController := controllers.NewAPIController(ctr)
	sampleController := controllers.NewSampleController(ctr)

	// bind controller functions to routes
	// api info
	r.HandleFunc("/", apiController.GetInfo).Methods(http.MethodGet)
	// sample
	r.HandleFunc("/samples", sampleController.Get).Methods(http.MethodGet)

	return r
}
