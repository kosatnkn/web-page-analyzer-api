package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kosatnkn/web-page-analyzer-api/app/container"
	"github.com/kosatnkn/web-page-analyzer-api/transport/http/controllers"
	"github.com/kosatnkn/web-page-analyzer-api/transport/http/middleware"
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
	reportController := controllers.NewReportController(ctr)

	// bind controller functions to routes
	// api info
	r.HandleFunc("/", apiController.GetInfo).Methods(http.MethodGet, http.MethodOptions)
	// report
	r.HandleFunc("/reports", reportController.Get).Methods(http.MethodGet, http.MethodOptions)

	return r
}
