package router

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/dcruselex/food/internal/rapidapi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

var once sync.Once
var foods []rapidapi.Recipe

func newFoodsResponse(recipes []rapidapi.Recipe) []render.Renderer {
	list := []render.Renderer{}
	for _, recipe := range recipes {
		r := recipe
		list = append(list, &r)
	}
	return list
}

func (e *errResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type errResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func errRender(err error) render.Renderer {
	return &errResponse{
		Err:            err,
		HTTPStatusCode: 500,
		StatusText:     "Error rendering response",
		ErrorText:      err.Error(),
	}
}

func listFoods(w http.ResponseWriter, r *http.Request) {
	once.Do(func() {
		var err error
		foods, err = rapidapi.Extract()
		if err != nil {
			log.Fatal(err)
		}
	})

	option, err := strconv.Atoi(r.URL.Query().Get("option"))
	if err != nil {
		render.Render(w, r, errRender(err))
		return
	}

	filteredFoods := getFilteredFoods(option)

	if err := render.RenderList(w, r, newFoodsResponse(filteredFoods)); err != nil {
		render.Render(w, r, errRender(err))
		return
	}
}

func GetRouter() *chi.Mux {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/foods", func(r chi.Router) {
		r.Get("/", listFoods)
	})

	root := "./static"
	fs := http.FileServer(http.Dir(root))

	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat(root + r.RequestURI); os.IsNotExist(err) {
			http.StripPrefix(r.RequestURI, fs).ServeHTTP(w, r)
		} else {
			fs.ServeHTTP(w, r)
		}
	})

	return r
}
