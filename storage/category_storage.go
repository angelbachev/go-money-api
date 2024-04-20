package storage

import (
	"database/sql"
	"fmt"

	"github.com/angelbachev/go-money-api/models"
	_ "github.com/go-sql-driver/mysql"
)

type CategoryStore interface {
	CreateCategory(cateory *models.Category) error
	GetCategories(userID, budgetID int64) ([]*models.Category, error)
	GetCategoryByID(id int64) (*models.Category, error)
	GetCategoryTree(budgetID int64) (*models.CategoryTree, error)
}

func (s MySQLStore) CreateCategory(category *models.Category) error {
	query := `
	INSERT INTO categories (user_id, budget_id, parent_id, name, description, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	result, err := s.db.Exec(
		query,
		category.UserID,
		category.BudgetID,
		category.ParentID,
		category.Name,
		category.Description,
		category.CreatedAt,
		category.UpdatedAt,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	category.ID = id

	return nil
}

func (s MySQLStore) GetCategories(userID, budgetID int64) ([]*models.Category, error) {
	query := `SELECT * FROM categories WHERE user_id = ? AND budget_id = ? ORDER BY name ASC`
	rows, err := s.db.Query(query, userID, budgetID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	categories := []*models.Category{}
	for rows.Next() {
		var category models.Category
		err = rows.Scan(
			&category.ID,
			&category.UserID,
			&category.BudgetID,
			&category.ParentID,
			&category.Name,
			&category.Description,
			&category.CreatedAt,
			&category.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		categories = append(categories, &category)
	}
	return categories, nil
}

func (s MySQLStore) GetCategoryByID(id int64) (*models.Category, error) {
	query := `SELECT * FROM categories WHERE id = ? LIMIT 1`
	row := s.db.QueryRow(query, id)

	return scanIntoCategory(row)
}

func (s MySQLStore) GetCategoryTree(budgetID int64) (*models.CategoryTree, error) {
	query := `
		WITH RECURSIVE tree_path (id, user_id, parent_id, name, description, path, created_at, updated_at) AS
		(
			SELECT id, user_id, parent_id, name, description, CONCAT(name, '/') as path, created_at, updated_at
    		FROM categories
    		WHERE budget_id = ? AND parent_id = 0 -- the tree node for given budget
			UNION ALL
			SELECT t.id, t.user_id, t.parent_id, t.name, t.description, CONCAT(tp.path, t.name, '/'), t.created_at, t.updated_at
			FROM tree_path AS tp 
    		JOIN categories AS t ON tp.id = t.parent_id AND t.budget_id = ?
		)
		SELECT * FROM tree_path
		ORDER BY path;
	`
	rows, err := s.db.Query(query, budgetID, budgetID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var categoryNodes []*models.CategoryTree

	for rows.Next() {
		var path string

		categoryNode := models.CategoryTree{Children: []*models.CategoryTree{}}

		err = rows.Scan(
			&categoryNode.ID,
			&categoryNode.UserID,
			&categoryNode.ParentID,
			&categoryNode.Name,
			&categoryNode.Description,
			&path,
			&categoryNode.CreatedAt,
			&categoryNode.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		categoryNodes = append(categoryNodes, &categoryNode)
	}
	return buildTree(categoryNodes), nil
}

func buildTree(categories []*models.CategoryTree) *models.CategoryTree {

	// a map, to keep track of each individual subtree.
	// Using a pointer to the CategoryNode struct so as to ensure that there's
	// only a single copy of each struct
	subtrees := map[int64]*models.CategoryTree{}

	var rootID int64
	// populate the map: every node is the root of its own subtree
	for _, cat := range categories {
		if cat.ParentID == 0 {
			rootID = cat.ID
		}
		subtrees[cat.ID] = cat
	}

	// iterate over the list of categories
	for _, cat := range categories {

		// if this is not the root node, it belongs to other category
		if cat.ParentID != 0 {

			// look up their immediate parent
			subtree := subtrees[cat.ParentID]

			// add them as a direct children
			subtree.Children = append(subtree.Children, cat)

		}

	}

	fmt.Printf("flat list: %v", categories)
	fmt.Printf("nested list: %v", subtrees)

	// At the end of the day, now, the tree is fully populated
	// return the root node for the entire tree
	return subtrees[rootID]
}

func scanIntoCategory(row *sql.Row) (*models.Category, error) {
	var category models.Category
	switch err := row.Scan(
		&category.ID,
		&category.BudgetID,
		&category.ParentID,
		&category.Name,
		&category.Description,
		&category.CreatedAt,
		&category.UpdatedAt,
	); err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return &category, nil
	default:
		return nil, err
	}
}
