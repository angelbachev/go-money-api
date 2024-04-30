package delete_account

type DeleteAccountCommand struct {
	ID     int64
	UserID int64
	Force  bool
}
