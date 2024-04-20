package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/angelbachev/go-money-api/models"
	"github.com/angelbachev/go-money-api/storage"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

type Server struct {
	listenAddr string
	store      storage.Store
}

func NewServer(listenAddr string, store storage.Store) *Server {
	return &Server{listenAddr: listenAddr, store: store}
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
		w.Write([]byte("Hello World!"))
	})

	// Creating a New Router
	apiRouter := chi.NewRouter()

	// Protected routes
	apiRouter.Group(func(r chi.Router) {
		setTokenAuth()
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(tokenAuth))

		// Handle valid / invalid tokens. In this example, we use
		// the provided authenticator middleware, but you can write your
		// own very easily, look at the Authenticator method in jwtauth.go
		// and tweak it, its not scary.
		r.Use(jwtauth.Authenticator(tokenAuth))

		r.Post("/budgets", s.handleCreateBudget)

		r.Post("/budgets/{budgetID}/categories", s.handleCreateCategory)
		r.Get("/budgets/{budgetID}/categories", s.handleListCategories)

		r.Post("/budgets/{budgetID}/expenses", s.handleCreateExpense)
		r.Get("/budgets/{budgetID}/expenses", s.handleListExpenses)
	})

	// Public routes
	apiRouter.Group(func(r chi.Router) {
		r.Post("/register", s.handleRegisterUser)
		r.Post("/login", s.handleLoginUser)
	})

	// Mounting the new Sub Router on the main router
	r.Mount("/api", apiRouter)

	return r
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func getBody(r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(&data)
}

func getAuthUserID(r *http.Request) int64 {
	_, claims, _ := jwtauth.FromContext(r.Context())

	return int64(claims["userId"].(float64))
}

func (s Server) handleRegisterUser(w http.ResponseWriter, r *http.Request) {
	// TODO: validate email and password
	var req RegisterUserRequest
	if err := getBody(r, &req); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	existingUser, _ := s.store.GetUserByEmail(req.Email)
	if existingUser != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": "Unable to register user"})
		return
	}

	user, err := models.NewUser(req.Email, req.Password)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
	}

	s.store.CreateUser(user)
	writeJSON(w, http.StatusCreated, *user)
}

func (s Server) handleLoginUser(w http.ResponseWriter, r *http.Request) {
	// TODO: validate email and password
	var req LoginUserRequest
	if err := getBody(r, &req); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := s.store.GetUserByEmail(req.Email)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": "Unable to login user"})
		return
	}

	if !user.ValidPassword(req.Password) {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": "Unable to login user"})
		return
	}

	fmt.Println(createJWT(user.ID))
	writeJSON(w, http.StatusCreated, map[string]string{"token": createJWT(user.ID)})
}

func (s Server) handleCreateBudget(w http.ResponseWriter, r *http.Request) {
	userID := getAuthUserID(r)

	var req CreateBudgetRequest
	if err := getBody(r, &req); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	budget := models.NewBudget(userID, req.Name, req.Description)

	s.store.CreateBudget(budget)
	writeJSON(w, http.StatusCreated, budget)
}

func (s Server) handleCreateCategory(w http.ResponseWriter, r *http.Request) {
	userID := getAuthUserID(r)

	var req CreateCategoryRequest
	if err := getBody(r, &req); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	budgetID, err := strconv.ParseInt(chi.URLParam(r, "budgetID"), 10, 0)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	category := models.NewCategory(userID, budgetID, req.ParentID, req.Name, req.Description)

	s.store.CreateCategory(category)
	writeJSON(w, http.StatusCreated, category)
}

func (s Server) handleListCategories(w http.ResponseWriter, r *http.Request) {
	// TODO: validate userID
	// userID := getAuthUserID(r)

	budgetID, err := strconv.ParseInt(chi.URLParam(r, "budgetID"), 10, 0)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	// categories, err := s.store.GetCategories(userID, budgetID)
	categories, err := s.store.GetCategoryTree(budgetID)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, categories)
}

func (s Server) handleCreateExpense(w http.ResponseWriter, r *http.Request) {
	userID := getAuthUserID(r)

	var req CreateExpenseRequest
	if err := getBody(r, &req); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	budgetID, err := strconv.ParseInt(chi.URLParam(r, "budgetID"), 10, 0)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	// TODO: validate that the user owns the budget and category

	expense := models.NewExpense(userID, budgetID, req.CategoryID, req.Description, req.Amount, req.Date)
	if err := s.store.CreateExpense(expense); err != nil {
		fmt.Printf("%v", err)
		return
	}

	writeJSON(w, http.StatusCreated, expense)
}

func (s Server) handleListExpenses(w http.ResponseWriter, r *http.Request) {
	userID := getAuthUserID(r)

	budgetID, err := strconv.ParseInt(chi.URLParam(r, "budgetID"), 10, 0)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	// TODO: validate user owns the budget
	expenses, err := s.store.GetExpenses(userID, budgetID)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, expenses)
}

// TODO: delete expense
// TODO: delete
