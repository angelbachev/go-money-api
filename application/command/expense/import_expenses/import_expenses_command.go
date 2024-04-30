package import_expenses

import "mime/multipart"

type ImportExpensesCommand struct {
	UserID    int64
	AccountID int64
	File      multipart.File
}
