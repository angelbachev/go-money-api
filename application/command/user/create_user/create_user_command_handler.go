package create_user

import (
	"errors"

	"github.com/angelbachev/go-money-api/domain/account"
	"github.com/angelbachev/go-money-api/domain/category"
	"github.com/angelbachev/go-money-api/domain/user"
)

type CreateUserCommandHandler struct {
	userRepository         user.UserRepositoryInterface
	userSettingsRepository user.UserSettingsRepositoryInterface
	accountRepository      account.AccountRepositoryInterface
	categoryRepository     category.CategoryRepositoryInterface
}

func NewCreateUserCommandHandler(
	userRepository user.UserRepositoryInterface,
	userSettingsRepository user.UserSettingsRepositoryInterface,
	accountRepository account.AccountRepositoryInterface,
	categoryRepository category.CategoryRepositoryInterface,
) *CreateUserCommandHandler {
	return &CreateUserCommandHandler{
		userRepository:         userRepository,
		userSettingsRepository: userSettingsRepository,
		accountRepository:      accountRepository,
		categoryRepository:     categoryRepository,
	}
}

func (h CreateUserCommandHandler) Handle(command CreateUserCommand) (*CreateUserResponse, error) {
	// TODO: validate email and password

	existingUser, _ := h.userRepository.GetUserByEmail(command.Email)
	if existingUser != nil {
		return nil, errors.New("Unable to register user")
	}

	u, err := user.NewUser(command.Email, command.Password)
	if err != nil {
		return nil, err
	}

	h.userRepository.CreateUser(u)

	// TODO: handle errors
	h.createDefaultAccountAndCategories(u)

	return &CreateUserResponse{User: u}, nil
}

func (h CreateUserCommandHandler) createDefaultAccountAndCategories(u *user.User) {
	// Create default account and categories
	account := account.NewAccount(u.ID, "Акаунт", "Описание на акаунт", "BGN")
	h.accountRepository.CreateAccount(account)
	settings := user.NewUserSettings(u.ID, account.ID, "default")
	h.userSettingsRepository.CreateUserSettings(settings)

	h.createCategories(u.ID, account.ID)
}

func (h CreateUserCommandHandler) createCategories(userID, accountID int64) {
	rootCategory := category.NewCategory(userID, accountID, 0, "", "", "")
	h.categoryRepository.CreateCategory(rootCategory)
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
		category := category.NewCategory(userID, accountID, rootCategory.ID, cat["name"].(string), cat["description"].(string), cat["icon"].(string))
		h.categoryRepository.CreateCategory(category)
	}
}
