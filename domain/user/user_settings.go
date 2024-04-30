package user

import (
	"time"
)

type UserSettings struct {
	ID               int64     `json:"-"`
	UserID           int64     `json:"-"`
	DefaultAccountID int64     `json:"defaultAccountId"`
	Theme            string    `json:"theme"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

func NewUserSettings(userID, defaultAccountID int64, theme string) *UserSettings {
	return &UserSettings{
		UserID:           userID,
		DefaultAccountID: defaultAccountID,
		Theme:            theme,
		CreatedAt:        time.Now().UTC(),
		UpdatedAt:        time.Now().UTC(),
	}
}

func (s *UserSettings) Update(defaultAccountID int64, theme string) {
	s.DefaultAccountID = defaultAccountID
	s.Theme = theme
	s.UpdatedAt = time.Now().UTC()
}
