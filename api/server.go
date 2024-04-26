package api

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

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
		tmpl := template.Must(template.ParseFiles("./views/index.html"))
		tmpl.Execute(w, map[string]any{"ReleasedAt": os.Getenv("RELEASED_AT")})
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

		r.Put("/user/settings", s.handleUpdateUserSettings)
		r.Get("/user/settings", s.handleGetUserSettings)

		r.Post("/accounts", s.handleCreateAccount)
		r.Put("/accounts/{accountID}", s.handleUpdateAccount)
		r.Get("/accounts", s.handleListAccounts)
		r.Delete("/accounts/{accountID}", s.handleDeleteAccount)

		r.Post("/accounts/{accountID}/categories", s.handleCreateCategory)
		r.Get("/accounts/{accountID}/categories", s.handleListCategories)
		r.Get("/accounts/{accountID}/categories/{categoryID}", s.handleGetCategory)
		r.Delete("/accounts/{accountID}/categories/{categoryID}", s.handleDeleteCategory)
		r.Put("/accounts/{accountID}/categories/{categoryID}", s.handleUpdateCategory)

		r.Post("/accounts/{accountID}/expenses", s.handleCreateExpense)
		r.Get("/accounts/{accountID}/expenses", s.handleListExpenses)
		r.Delete("/accounts/{accountID}/expenses/{expenseID}", s.handleDeleteExpense)
		r.Put("/accounts/{accountID}/expenses/{expenseID}", s.handleUpdateExpense)
		r.Post("/accounts/{accountID}/expenses/import", s.handleImportExpenses)

	})

	// Public routes
	apiRouter.Group(func(r chi.Router) {
		r.Post("/users", s.handleRegisterUser)
		r.Post("/auth/tokens", s.handleLoginUser)
		r.Get("/category-icons", s.handleListCategoryIcons)
	})

	// Mounting the new Sub Router on the main router
	r.Mount("/api", apiRouter)

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "files"))
	fmt.Println(filesDir)
	FileServer(r, "/files", filesDir)

	return r
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func getBody(r *http.Request, data interface{}) error {
	// var f interface{}
	// json.NewDecoder(r.Body).Decode(&f)

	// fmt.Printf("form: %+v", f)
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

	// Create default account and categories
	account := models.NewAccount(user.ID, "Акаунт", "Описание на акаунт", "BGN")
	s.store.CreateAccount(account)
	settings := models.NewUserSettings(user.ID, account.ID, "default")
	s.store.CreateUserSettings(settings)

	rootCategory := models.NewCategory(user.ID, account.ID, 0, "", "", "")
	s.store.CreateCategory(rootCategory)
	categories := []map[string]any{
		{"name": "Храна", "description": "Храна, ресторанти и др.", "icon": "food.svg"},
		{"name": "Сметки", "description": "Наем, данъци, комунални услуги", "icon": "taxes.svg"},
		{"name": "Транспорт", "description": "Градски транспорт, кола, такси, др.", "icon": "transport.svg"},
		{"name": "Забавления", "description": "", "icon": "entertainment.svg"},
		{"name": "Хоби", "description": "", "icon": "hobbies.svg"},
		{"name": "Спорт", "description": "Спортно оборудване, карти за спорт и др.", "icon": "sport.svg"},
		{"name": "Облекло", "description": "", "icon": "clothes.svg"},
		{"name": "Образование", "description": "Книги, уроци, курсове, семинари и др.", "icon": "education.svg"},
		{"name": "Здраве", "description": "Медицински прегледи, лекарства и др.", "icon": "health.svg"},
		{"name": "Почивка", "description": "Хотели, самолетни билети, екскурзоводи, музеи и др.", "icon": "travelling.svg"},
		{"name": "Дом", "description": "Home", "icon": "home.svg"},
		{"name": "Други", "description": "Всичко останало", "icon": "others.svg"},
	}
	for _, cat := range categories {
		category := models.NewCategory(user.ID, account.ID, rootCategory.ID, cat["name"].(string), cat["description"].(string), cat["icon"].(string))
		s.store.CreateCategory(category)
	}

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

	settings, err := s.store.GetUserSettings(user.ID)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": "Unable to login user"})
		return
	}

	writeJSON(w, http.StatusCreated, map[string]any{
		"accessToken": createJWT(user.ID),
		"settings":    settings,
	})
}

func (s Server) handleUpdateUserSettings(w http.ResponseWriter, r *http.Request) {
	userID := getAuthUserID(r)

	var req UpdateUserSettingsRequest
	if err := getBody(r, &req); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	settings, err := s.store.GetUserSettings(userID)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	settings.Update(req.DefaultAccountID, req.Theme)
	s.store.UpdateUserSettings(settings)
	writeJSON(w, http.StatusOK, settings)
}

func (s Server) handleGetUserSettings(w http.ResponseWriter, r *http.Request) {
	userID := getAuthUserID(r)

	settings, err := s.store.GetUserSettings(userID)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, settings)
}

func (s Server) handleCreateAccount(w http.ResponseWriter, r *http.Request) {
	userID := getAuthUserID(r)

	var req CreateAccountRequest
	if err := getBody(r, &req); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	account := models.NewAccount(userID, req.Name, req.Description, req.CurrencyCode)

	s.store.CreateAccount(account)
	writeJSON(w, http.StatusCreated, account)
}

func (s Server) handleUpdateAccount(w http.ResponseWriter, r *http.Request) {
	userID := getAuthUserID(r)

	accountID, err := strconv.ParseInt(chi.URLParam(r, "accountID"), 10, 0)
	if err != nil {
		writeJSON(w, http.StatusNotFound, err.Error())
		return
	}

	var req UpdateAccountRequest
	if err := getBody(r, &req); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	account, err := s.store.GetAccountByID(userID, accountID)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	account.Update(req.Name, req.Description, req.CurrencyCode)
	err = s.store.UpdateAccount(account)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, account)
}

func (s Server) handleListAccounts(w http.ResponseWriter, r *http.Request) {
	// TODO: validate userID
	userID := getAuthUserID(r)

	accounts, err := s.store.GetAccounts(userID)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, accounts)
}

func (s Server) handleDeleteAccount(w http.ResponseWriter, r *http.Request) {
	userID := getAuthUserID(r)

	accountID, err := strconv.ParseInt(chi.URLParam(r, "accountID"), 10, 0)
	if err != nil {
		writeJSON(w, http.StatusNotFound, err.Error())
		return
	}

	_, err = s.store.GetAccountByID(userID, accountID)
	if err != nil {
		writeJSON(w, http.StatusNotFound, err.Error())
		return
	}

	categoryIDs, err := s.store.GetCategories(accountID)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	// TODO: handle subcategories

	filters := &models.ExpenseFilters{
		CategoryIDs: categoryIDs,
	}

	// TODO: validate that the user owns the account and category

	expenses, err := s.store.GetExpenses(userID, accountID, filters, 0, 0)
	if err != nil {
		writeJSON(w, http.StatusNotFound, err.Error())
		return
	}

	q := r.URL.Query()
	force, _ := strconv.ParseInt(q.Get("force"), 10, 0)
	if force == 0 && (len(categoryIDs) > 1 || len(expenses) > 0) {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "Account is not empty"})
		return
	}

	for _, exp := range expenses {
		s.store.DeleteExpense(exp.ID)
	}

	for _, cat := range categoryIDs {
		s.store.DeleteCategory(cat)
	}

	s.store.DeleteAccount(accountID)

	w.WriteHeader(http.StatusNoContent)
}

func (s Server) handleCreateCategory(w http.ResponseWriter, r *http.Request) {
	userID := getAuthUserID(r)

	var req CreateCategoryRequest
	if err := getBody(r, &req); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	accountID, err := strconv.ParseInt(chi.URLParam(r, "accountID"), 10, 0)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	if req.ParentID == 0 {
		req.ParentID, err = s.store.GetRootCategoryID(accountID)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, err.Error())
			return
		}
	}

	category := models.NewCategory(userID, accountID, req.ParentID, req.Name, req.Description, req.Icon)

	s.store.CreateCategory(category)
	writeJSON(w, http.StatusCreated, category)
}

func (s Server) handleListCategories(w http.ResponseWriter, r *http.Request) {
	// TODO: validate userID
	// userID := getAuthUserID(r)

	accountID, err := strconv.ParseInt(chi.URLParam(r, "accountID"), 10, 0)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	categories, err := s.store.GetCategoryTree(accountID)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, categories)
}

func (s Server) handleGetCategory(w http.ResponseWriter, r *http.Request) {
	// TODO: validate userID
	// userID := getAuthUserID(r)

	// accountID, err := strconv.ParseInt(chi.URLParam(r, "accountID"), 10, 0)
	// if err != nil {
	// 	writeJSON(w, http.StatusBadRequest, err.Error())
	// 	return
	// }

	categoryID, err := strconv.ParseInt(chi.URLParam(r, "categoryID"), 10, 0)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	category, err := s.store.GetSingleCategoryTree(categoryID)
	fmt.Printf("category %v", category)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	// TODO: validate account

	writeJSON(w, http.StatusOK, category)
}

func (s Server) handleDeleteCategory(w http.ResponseWriter, r *http.Request) {
	userID := getAuthUserID(r)

	accountID, err := strconv.ParseInt(chi.URLParam(r, "accountID"), 10, 0)
	if err != nil {
		writeJSON(w, http.StatusNotFound, err.Error())
		return
	}

	categoryID, err := strconv.ParseInt(chi.URLParam(r, "categoryID"), 10, 0)
	if err != nil {
		writeJSON(w, http.StatusNotFound, err.Error())
		return
	}

	categoryIDs, err := s.store.GetListCategoryIDsAndTheirSubcategories([]int64{categoryID})
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	// TODO: handle subcategories

	filters := &models.ExpenseFilters{
		CategoryIDs: categoryIDs,
	}

	// TODO: validate that the user owns the account and category

	expenses, err := s.store.GetExpenses(userID, accountID, filters, 0, 0)
	if err != nil {
		writeJSON(w, http.StatusNotFound, err.Error())
		return
	}

	q := r.URL.Query()
	force, _ := strconv.ParseInt(q.Get("force"), 10, 0)
	if force == 0 && (len(categoryIDs) > 1 || len(expenses) > 0) {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "Category is not empty"})
		return
	}

	for _, exp := range expenses {
		s.store.DeleteExpense(exp.ID)
	}

	for _, cat := range categoryIDs {
		s.store.DeleteCategory(cat)
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s Server) handleUpdateCategory(w http.ResponseWriter, r *http.Request) {
	// userID := getAuthUserID(r)

	accountID, err := strconv.ParseInt(chi.URLParam(r, "accountID"), 10, 0)
	if err != nil {
		writeJSON(w, http.StatusNotFound, err.Error())
		return
	}

	categoryID, err := strconv.ParseInt(chi.URLParam(r, "categoryID"), 10, 0)
	if err != nil {
		writeJSON(w, http.StatusNotFound, err.Error())
		return
	}

	var req UpdateCategoryRequest
	if err := getBody(r, &req); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	category, err := s.store.GetCategoryByID(categoryID)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	if category == nil || category.AccountID != accountID {
		writeJSON(w, http.StatusNotFound, err.Error())
		return
	}

	if req.ParentID == 0 {
		req.ParentID, err = s.store.GetRootCategoryID(accountID)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, err.Error())
			return
		}
	}
	category.Update(req.ParentID, req.Name, req.Description, req.Icon)
	err = s.store.UpdateCategory(category)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, category)
}

func (s Server) handleCreateExpense(w http.ResponseWriter, r *http.Request) {
	userID := getAuthUserID(r)

	var req CreateExpenseRequest
	if err := getBody(r, &req); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	accountID, err := strconv.ParseInt(chi.URLParam(r, "accountID"), 10, 0)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	// TODO: validate that the user owns the account and category

	expense := models.NewExpense(userID, accountID, req.CategoryID, req.Description, req.Amount, req.Date)
	if err := s.store.CreateExpense(expense); err != nil {
		fmt.Printf("%v", err)
		return
	}

	writeJSON(w, http.StatusCreated, expense)
}

func (s Server) handleUpdateExpense(w http.ResponseWriter, r *http.Request) {
	userID := getAuthUserID(r)
	// TODO: validate new category belongs to the same account
	accountID, err := strconv.ParseInt(chi.URLParam(r, "accountID"), 10, 0)
	if err != nil {
		writeJSON(w, http.StatusNotFound, err.Error())
		return
	}

	expenseID, err := strconv.ParseInt(chi.URLParam(r, "expenseID"), 10, 0)
	if err != nil {
		writeJSON(w, http.StatusNotFound, err.Error())
		return
	}

	var req UpdateExpenseRequest
	if err := getBody(r, &req); err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	expense, err := s.store.GetExpenseByID(userID, accountID, expenseID)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	if expense == nil || expense.AccountID != accountID {
		writeJSON(w, http.StatusNotFound, err.Error())
		return
	}
	expense.Update(req.CategoryID, req.Description, req.Amount, req.Date)
	err = s.store.UpdateExpense(expense)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, expense)
}

func (s Server) handleListExpenses(w http.ResponseWriter, r *http.Request) {
	userID := getAuthUserID(r)

	accountID, err := strconv.ParseInt(chi.URLParam(r, "accountID"), 10, 0)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	q := r.URL.Query()
	minAm, err := strconv.ParseInt(q.Get("minAmount"), 10, 0)
	var minAmount *int64
	if err == nil && minAm != 0 {
		minAmount = &minAm
	}
	maxAm, err := strconv.ParseInt(q.Get("maxAmount"), 10, 0)
	var maxAmount *int64
	if err == nil && maxAm != 0 {
		maxAmount = &maxAm
	}
	minDt, err := time.Parse(time.RFC3339, q.Get("minDate"))
	var minDate *time.Time
	if err == nil {
		minDate = &minDt
	}
	maxDt, err := time.Parse(time.RFC3339, q.Get("maxDate"))
	var maxDate *time.Time
	if err == nil {
		maxDate = &maxDt
	}

	var categories []int64
	for _, cat := range q["categoryIds[]"] {
		ct, err := strconv.ParseInt(cat, 10, 0)
		if err != nil {
			// todo: handle non integer value error
		}

		categories = append(categories, ct)
	}
	var categoryIDs []int64
	if len(categories) > 0 {
		categoryIDs, err = s.store.GetListCategoryIDsAndTheirSubcategories(categories)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, err.Error())
			return
		}
	}

	// TODO: handle subcategories

	filters := &models.ExpenseFilters{
		MinAmount:   minAmount,
		MaxAmount:   maxAmount,
		MinDate:     minDate,
		MaxDate:     maxDate,
		CategoryIDs: categoryIDs,
	}

	p := q.Get("page")
	if p == "" || p == "0" {
		p = "1"
	}
	page, err := strconv.ParseInt(p, 10, 0)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	l := q.Get("limit")
	if l == "" || l == "0" {
		l = "10"
	}
	limit, err := strconv.ParseInt(l, 10, 0)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	// TODO: validate user owns the account
	expenses, err := s.store.GetExpenses(userID, accountID, filters, page, limit)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	totalCount, err := s.store.GetExpensesCount(userID, accountID, filters)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{"items": expenses, "totalCount": totalCount})
}

func (s Server) handleDeleteExpense(w http.ResponseWriter, r *http.Request) {
	userID := getAuthUserID(r)

	accountID, err := strconv.ParseInt(chi.URLParam(r, "accountID"), 10, 0)
	if err != nil {
		writeJSON(w, http.StatusNotFound, err.Error())
		return
	}

	expenseID, err := strconv.ParseInt(chi.URLParam(r, "expenseID"), 10, 0)
	if err != nil {
		writeJSON(w, http.StatusNotFound, err.Error())
		return
	}

	// TODO: validate that the user owns the account and category

	_, err = s.store.GetExpenseByID(userID, accountID, expenseID)
	if err != nil {
		writeJSON(w, http.StatusNotFound, err.Error())
		return
	}

	if err := s.store.DeleteExpense(expenseID); err != nil {
		writeJSON(w, http.StatusNotFound, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
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

func (s Server) handleListCategoryIcons(w http.ResponseWriter, r *http.Request) {
	icons, err := os.ReadDir("./files/images/categories")
	if err != nil {
		writeJSON(w, http.StatusNotFound, err.Error())
		return
	}

	var iconNames []string
	for _, icon := range icons {
		iconNames = append(iconNames, icon.Name())
	}

	writeJSON(w, http.StatusOK, iconNames)
}

func (s Server) handleImportExpenses(w http.ResponseWriter, r *http.Request) {
	userID := getAuthUserID(r)

	accountID, err := strconv.ParseInt(chi.URLParam(r, "accountID"), 10, 0)
	if err != nil {
		writeJSON(w, http.StatusNotFound, err.Error())
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		writeJSON(w, http.StatusNotFound, err.Error())
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	existingCategories, err := s.store.GetCategoryNames(accountID)
	if err != nil {
		writeJSON(w, http.StatusNotFound, err.Error())
		return
	}

	rootCategoryID, err := s.store.GetRootCategoryID(accountID)
	if err != nil {
		writeJSON(w, http.StatusNotFound, err.Error())
		return
	}

	// Print the CSV data
	for _, row := range data {
		if row[0] == "" || row[2] == "" {
			continue // empty line
		}
		date, err := time.Parse("02.01.2006", row[0])
		if err != nil {
			// TODO: handle error
			continue
		}

		re := regexp.MustCompile("[0-9]+.?,?[0-9]+")
		amount, err := strconv.ParseFloat(strings.Replace(re.FindString(row[2]), ",", ".", 1), 64)
		if err != nil {
			// TODO: handle error
			continue
		}
		categoryName := strings.TrimSpace(row[1])
		categoryID, ok := existingCategories[strings.ToLower(categoryName)]
		if !ok {
			category := models.NewCategory(userID, accountID, rootCategoryID, categoryName, "", "")
			err = s.store.CreateCategory(category)
			if err != nil {
				// TODO: handle error
				continue
			}
			categoryID = category.ID
			existingCategories[strings.ToLower(categoryName)] = categoryID
		}

		expense := models.NewExpense(userID, accountID, categoryID, strings.TrimSpace(row[3]), int64(amount*100), date)
		s.store.CreateExpense(expense)
	}

	// TODO: prevent entering the same expense multiple times
	// TODO: return error lines
	// TODO: add support for different currencies

	writeJSON(w, http.StatusNoContent, "")
}
