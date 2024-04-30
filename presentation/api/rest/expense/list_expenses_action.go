package expense

import (
	"net/http"
	"strconv"
	"time"

	"github.com/angelbachev/go-money-api/application/query/expense/list_expenses"
	"github.com/angelbachev/go-money-api/presentation/api/rest"
	"github.com/go-chi/chi/v5"
)

type ListExpensesAction struct {
	*rest.BaseAction
	handler *list_expenses.ListExpensesQueryHandler
}

func NewListExpensesAction(handler *list_expenses.ListExpensesQueryHandler) *ListExpensesAction {
	return &ListExpensesAction{
		BaseAction: rest.NewBaseAction("Get", "/accounts/{accountID}/expenses", false),
		handler:    handler,
	}
}

func (a ListExpensesAction) Handle(w http.ResponseWriter, r *http.Request) {
	accountID, err := strconv.ParseInt(chi.URLParam(r, "accountID"), 10, 0)
	if err != nil {
		a.Error(w, http.StatusBadRequest, err)
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

	p := q.Get("page")
	if p == "" || p == "0" {
		p = "1"
	}
	page, err := strconv.ParseInt(p, 10, 0)
	if err != nil {
		a.Error(w, http.StatusBadRequest, err)
		return
	}

	l := q.Get("limit")
	if l == "" || l == "0" {
		l = "10"
	}
	limit, err := strconv.ParseInt(l, 10, 0)
	if err != nil {
		a.Error(w, http.StatusBadRequest, err)
		return
	}

	// TODO: validate userID
	query := list_expenses.ListExpensesQuery{
		UserID:      a.GetAuthUserID(r),
		AccountID:   accountID,
		MinAmount:   minAmount,
		MaxAmount:   maxAmount,
		MinDate:     minDate,
		MaxDate:     maxDate,
		CategoryIDs: categories,
		Page:        page,
		Limit:       limit,
	}

	response, err := a.handler.Handle(query)
	if err != nil {
		a.Error(w, http.StatusBadRequest, err)
		return
	}

	a.JSON(w, http.StatusOK, response)
}
