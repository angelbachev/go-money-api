package update_user_settings

type UpdateUserSettingsCommand struct {
	UserID           int64
	DefaultAccountID int64  `json:"defaultAccountId"`
	Theme            string `json:"theme"`
}
