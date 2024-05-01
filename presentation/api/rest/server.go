package rest

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/angelbachev/go-money-api/infrastructure/domain/auth"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

type Server struct {
	listenAddr string
	actions    []APIAction
	auth       auth.JWTAuthService
}

func NewServer(listenAddr string, auth auth.JWTAuthService, actions []APIAction) *Server {
	return &Server{
		listenAddr: listenAddr,
		auth:       auth,
		actions:    actions,
	}
}

func (s Server) Run() {
	fmt.Printf("Starting server on %v\n", s.listenAddr)
	if err := http.ListenAndServe(s.listenAddr, s.router()); err != nil {
		log.Fatalf(err.Error())
	}
}

func (s Server) router() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./views/index.html"))
		tmpl.Execute(w, map[string]any{"ReleasedAt": os.Getenv("RELEASED_AT")})
	})

	// Creating a New Router
	apiRouter := chi.NewRouter()

	// Protected routes
	apiRouter.Group(func(r chi.Router) {
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(s.auth.GetJWTAuth()))

		// Handle valid / invalid tokens. In this example, we use
		// the provided authenticator middleware, but you can write your
		// own very easily, look at the Authenticator method in jwtauth.go
		// and tweak it, its not scary.
		r.Use(jwtauth.Authenticator(s.auth.GetJWTAuth()))

		for _, action := range s.actions {
			if action.IsPublic() {
				continue
			}

			r.MethodFunc(action.Method(), action.Route(), action.Handle)
		}
	})

	// Public routes
	apiRouter.Group(func(r chi.Router) {
		for _, action := range s.actions {
			if !action.IsPublic() {
				continue
			}

			r.MethodFunc(action.Method(), action.Route(), action.Handle)
		}
	})

	// Mounting the new Sub Router on the main router
	r.Mount("/api", apiRouter)

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "files"))
	fmt.Println(filesDir)
	FileServer(r, "/files", filesDir)

	return r
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
