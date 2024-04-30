package account

type AccountRepositoryInterface interface {
	CreateAccount(account *Account) error
	UpdateAccount(account *Account) error
	GetAccountByID(userID, accountID int64) (*Account, error)
	GetAccounts(userID int64) ([]*Account, error)
	DeleteAccount(id int64) error
}
