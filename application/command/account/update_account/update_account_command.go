package update_account

type UpdateAccountCommand struct {
	ID           int64
	UserID       int64
	Name         string `json:"name"`
	Description  string `json:"description"`
	CurrencyCode string `json:"currencyCode"`
}
