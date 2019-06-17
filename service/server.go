package service

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	repo := newInMemoryRepository()

	initRoutes(mx, formatter, repo)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render, repo metricRepository) {
	mx.HandleFunc("/metrics", getMetricsHandler(formatter, repo)).Methods("GET")
	mx.HandleFunc("/metrics/{appname}", getAppMetricHandler(formatter, repo)).Methods("GET")
}

func testHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"This is a test"})
	}
}
