package create_account

type CreateAccountCommand struct {
	UserID       int64
	Name         string `json:"name"`
	Description  string `json:"description"`
	CurrencyCode string `json:"currencyCode"`
}
