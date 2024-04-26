package main

import (
	"os"
	"time"

	"github.com/angelbachev/go-money-api/models"
	"github.com/angelbachev/go-money-api/storage"
	"github.com/joho/godotenv"
)

var store storage.Store

func init() {
	godotenv.Load(".env")
}
func main() {
	connStr := os.Getenv("MYSQL_CONNECTION_STRING")
	store, _ = storage.NewMySQLStore(connStr)

	seedData()
}

func seedData() {
	// userAcho := seedUser("acho_the_best@abv.bg", "Pass_123")
	// accountAcho1 := seedAccount(userAcho.ID, "Бюджетът на Ачо", "Това е моят бюджет", "BGN")
	// catRoot := seedCategory(userAcho.ID, accountAcho1.ID, 0, "", "")
	// catCar := seedCategory(userAcho.ID, accountAcho1.ID, catRoot.ID, "Кола", "Разходи за кола")
	// catGas := seedCategory(userAcho.ID, accountAcho1.ID, catCar.ID, "Гориво", "Разходи за гориво")
	// seedExpense(userAcho.ID, accountAcho1.ID, catGas.ID, "", 7486, "2024-03-17")
	// seedExpense(userAcho.ID, accountAcho1.ID, catGas.ID, "", 11739, "2024-04-01")
	// catParking := seedCategory(userAcho.ID, accountAcho1.ID, catCar.ID, "Паркиране", "Разходи за паркиране")
	// seedExpense(userAcho.ID, accountAcho1.ID, catParking.ID, "Синя зона", 400, "2024-04-15")
	// seedExpense(userAcho.ID, accountAcho1.ID, catParking.ID, "Синя зона", 600, "2024-04-16")
	// catService := seedCategory(userAcho.ID, accountAcho1.ID, catCar.ID, "Ремонти и консумативи", "Разходи за ремонти и консумативи")
	// seedExpense(userAcho.ID, accountAcho1.ID, catService.ID, "Предни гуми Continental", 52500, "2024-04-15")
	// seedCategory(userAcho.ID, accountAcho1.ID, catCar.ID, "Данъци", "Данъци и такси за колата")
	// catFood := seedCategory(userAcho.ID, accountAcho1.ID, catRoot.ID, "Храна", "Разходи за храна и напитки")
	// seedExpense(userAcho.ID, accountAcho1.ID, catFood.ID, "", 5914, "2024-04-15")
	// seedExpense(userAcho.ID, accountAcho1.ID, catFood.ID, "", 1497, "2024-04-16")
	// seedExpense(userAcho.ID, accountAcho1.ID, catFood.ID, "", 11300, "2024-04-17")

	// accountAcho2 := seedAccount(userAcho.ID, "Angel's additional account", "This is my additional account", "BGN")

	// userRosie := seedUser("roskanq@gmail.com", "Pass_123")
	// accountRosie := seedAccount(userRosie.ID, "Rosie's account", "This is my account", "BGN")

	// userPavlin := seedUser("blade_777@abv.bg", "Pass_123")
	// accountPavlin := seedAccount(userPavlin.ID, "Pavlin's account", "This is my account", "BGN")
}

func seedUser(email, password string) *models.User {
	user, _ := models.NewUser(email, password)
	store.CreateUser(user)

	return user
}

func seedAccount(userID int64, name, description, currencyCode string) *models.Account {
	account := models.NewAccount(userID, name, description, currencyCode)
	store.CreateAccount(account)

	return account
}

func seedCategory(userID, accountID, parentID int64, name, description string) *models.Category {
	category := models.NewCategory(userID, accountID, parentID, name, description)
	store.CreateCategory(category)

	return category
}

func seedExpense(userID, accountID, categoryID int64, description string, amount int64, dateStr string) *models.Expense {
	date, _ := time.Parse("2006-01-02", dateStr)
	expense := models.NewExpense(userID, accountID, categoryID, description, amount, date)
	store.CreateExpense(expense)

	return expense
}
