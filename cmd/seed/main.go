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
	userAcho := seedUser("acho_the_best@abv.bg", "Pass_123")
	budgetAcho1 := seedBudget(userAcho.ID, "Бюджетът на Ачо", "Това е моят бюджет")
	catRoot := seedCategory(userAcho.ID, budgetAcho1.ID, 0, "", "")
	catCar := seedCategory(userAcho.ID, budgetAcho1.ID, catRoot.ID, "Кола", "Разходи за кола")
	catGas := seedCategory(userAcho.ID, budgetAcho1.ID, catCar.ID, "Гориво", "Разходи за гориво")
	seedExpense(userAcho.ID, budgetAcho1.ID, catGas.ID, "", 7486, "2024-03-17")
	seedExpense(userAcho.ID, budgetAcho1.ID, catGas.ID, "", 11739, "2024-04-01")
	catParking := seedCategory(userAcho.ID, budgetAcho1.ID, catCar.ID, "Паркиране", "Разходи за паркиране")
	seedExpense(userAcho.ID, budgetAcho1.ID, catParking.ID, "Синя зона", 400, "2024-04-15")
	seedExpense(userAcho.ID, budgetAcho1.ID, catParking.ID, "Синя зона", 600, "2024-04-16")
	catService := seedCategory(userAcho.ID, budgetAcho1.ID, catCar.ID, "Ремонти и консумативи", "Разходи за ремонти и консумативи")
	seedExpense(userAcho.ID, budgetAcho1.ID, catService.ID, "Предни гуми Continental", 52500, "2024-04-15")
	seedCategory(userAcho.ID, budgetAcho1.ID, catCar.ID, "Данъци", "Данъци и такси за колата")
	catFood := seedCategory(userAcho.ID, budgetAcho1.ID, catRoot.ID, "Храна", "Разходи за храна и напитки")
	seedExpense(userAcho.ID, budgetAcho1.ID, catFood.ID, "", 5914, "2024-04-15")
	seedExpense(userAcho.ID, budgetAcho1.ID, catFood.ID, "", 1497, "2024-04-16")
	seedExpense(userAcho.ID, budgetAcho1.ID, catFood.ID, "", 11300, "2024-04-17")

	// budgetAcho2 := seedBudget(userAcho.ID, "Angel's additional budget", "This is my additional budget")

	// userRosie := seedUser("roskanq@gmail.com", "Pass_123")
	// budgetRosie := seedBudget(userRosie.ID, "Rosie's budget", "This is my budget")

	// userPavlin := seedUser("blade_777@abv.bg", "Pass_123")
	// budgetPavlin := seedBudget(userPavlin.ID, "Pavlin's budget", "This is my budget")
}

func seedUser(email, password string) *models.User {
	user, _ := models.NewUser(email, password)
	store.CreateUser(user)

	return user
}

func seedBudget(userID int64, name, description string) *models.Budget {
	budget := models.NewBudget(userID, name, description)
	store.CreateBudget(budget)

	return budget
}

func seedCategory(userID, budgetID, parentID int64, name, description string) *models.Category {
	category := models.NewCategory(userID, budgetID, parentID, name, description)
	store.CreateCategory(category)

	return category
}

func seedExpense(userID, budgetID, categoryID int64, description string, amount int64, dateStr string) *models.Expense {
	date, _ := time.Parse("2006-01-02", dateStr)
	expense := models.NewExpense(userID, budgetID, categoryID, description, amount, date)
	store.CreateExpense(expense)

	return expense
}
