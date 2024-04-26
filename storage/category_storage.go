package storage

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/angelbachev/go-money-api/models"
	_ "github.com/go-sql-driver/mysql"
)

type CategoryStore interface {
	CreateCategory(category *models.Category) error
	GetCategories(accountID int64) ([]int64, error)
	GetCategoryNames(accountID int64) (map[string]int64, error)
	GetCategoryByID(id int64) (*models.Category, error)
	GetCategoryTree(accountID int64) (*models.CategoryTree, error)
	GetSingleCategoryTree(id int64) (*models.CategoryTree, error)
	GetListCategoryIDsAndTheirSubcategories(ids []int64) ([]int64, error)
	DeleteCategory(id int64) error
	UpdateCategory(category *models.Category) error
	GetRootCategoryID(accountID int64) (int64, error)
}

func (s MySQLStore) CreateCategory(category *models.Category) error {
	query := `
	INSERT INTO categories (user_id, account_id, parent_id, name, description, icon, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	result, err := s.db.Exec(
		query,
		category.UserID,
		category.AccountID,
		category.ParentID,
		category.Name,
		category.Description,
		category.Icon,
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

func (s MySQLStore) GetCategories(accountID int64) ([]int64, error) {
	query := `SELECT id FROM categories WHERE account_id = ?`
	rows, err := s.db.Query(query, accountID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	categories := []int64{}
	for rows.Next() {
		var category int64
		err = rows.Scan(&category)

		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}
	return categories, nil
}

func (s MySQLStore) GetCategoryNames(accountID int64) (map[string]int64, error) {
	query := `SELECT id, name FROM categories WHERE account_id = ? AND name != ""`
	rows, err := s.db.Query(query, accountID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	categoryNames := map[string]int64{}
	for rows.Next() {
		var name string
		var id int64
		err = rows.Scan(&id, &name)

		if err != nil {
			return nil, err
		}

		categoryNames[strings.ToLower(name)] = id
	}
	return categoryNames, nil
}

func (s MySQLStore) GetCategoryByID(id int64) (*models.Category, error) {
	query := `SELECT * FROM categories WHERE id = ? LIMIT 1`
	row := s.db.QueryRow(query, id)

	return scanIntoCategory(row)
}

func (s MySQLStore) GetCategoryTree(accountID int64) (*models.CategoryTree, error) {
	query := `
		WITH RECURSIVE tree_path (id, user_id, parent_id, name, description, icon, path, created_at, updated_at) AS
		(
			SELECT id, user_id, parent_id, name, description, icon, CONCAT(name, '/') as path, created_at, updated_at
    		FROM categories
    		WHERE account_id = ? AND parent_id = 0 -- the tree node for given account
			UNION ALL
			SELECT t.id, t.user_id, t.parent_id, t.name, t.description, t.icon, CONCAT(tp.path, t.name, '/'), t.created_at, t.updated_at
			FROM tree_path AS tp 
    		JOIN categories AS t ON tp.id = t.parent_id AND t.account_id = ?
		)
		SELECT * FROM tree_path
		ORDER BY path;
	`
	rows, err := s.db.Query(query, accountID, accountID)

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
			&categoryNode.Icon,
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

func (s MySQLStore) GetSingleCategoryTree(id int64) (*models.CategoryTree, error) {
	query := `
		WITH RECURSIVE tree_path (id, user_id, parent_id, name, description, icon, path, created_at, updated_at) AS
		(
			SELECT id, user_id, parent_id, name, description, icon, CONCAT(name, '/') as path, created_at, updated_at
    		FROM categories
    		WHERE id = ? -- the given category
			UNION ALL
			SELECT t.id, t.user_id, t.parent_id, t.name, t.description, icon, CONCAT(tp.path, t.name, '/'), t.created_at, t.updated_at
			FROM tree_path AS tp 
    		JOIN categories AS t ON tp.id = t.parent_id
		)
		SELECT * FROM tree_path
		ORDER BY path;
	`
	rows, err := s.db.Query(query, id)

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
			&categoryNode.Icon,
			&path,
			&categoryNode.CreatedAt,
			&categoryNode.UpdatedAt,
		)

		if err != nil {
			fmt.Printf("err %v", err)
			return nil, err
		}
		fmt.Printf("cat %v", categoryNode)

		categoryNodes = append(categoryNodes, &categoryNode)
		fmt.Printf("cats %v", categoryNodes)

	}

	return buildTree(categoryNodes), nil
}

func (s MySQLStore) GetListCategoryIDsAndTheirSubcategories(ids []int64) ([]int64, error) {
	query := `
		WITH RECURSIVE tree_path (id, parent_id) AS
		(
			SELECT id, parent_id
			FROM categories
			WHERE id IN (%s) -- the given category
			UNION ALL
			SELECT t.id, t.parent_id
			FROM tree_path AS tp 
			JOIN categories AS t ON tp.id = t.parent_id
		)
		SELECT * FROM tree_path;
	`

	placeholders := make([]string, len(ids))
	args := make([]any, len(ids))
	for i, id := range ids {
		placeholders[i] = "?"
		args[i] = id
	}

	rows, err := s.db.Query(fmt.Sprintf(query, strings.Join(placeholders, ",")), args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var categoryIDs []int64

	for rows.Next() {
		var id, parentId int64

		err = rows.Scan(&id, &parentId)

		if err != nil {
			return nil, err
		}

		categoryIDs = append(categoryIDs, id)
	}

	return categoryIDs, nil
}

func buildTree(categories []*models.CategoryTree) *models.CategoryTree {

	// a map, to keep track of each individual subtree.
	// Using a pointer to the CategoryNode struct so as to ensure that there's
	// only a single copy of each struct
	subtrees := map[int64]*models.CategoryTree{}

	var rootID int64
	// populate the map: every node is the root of its own subtree
	for idx, cat := range categories {
		if idx == 0 {
			rootID = cat.ID
		}
		subtrees[cat.ID] = cat
	}
	fmt.Println("root", rootID)
	// iterate over the list of categories
	for idx, cat := range categories {

		// if this is not the root node, it belongs to other category
		if idx > 0 && cat.ParentID > 0 {
			// look up their immediate parent
			subtree := subtrees[cat.ParentID]
			// add them as a direct children
			subtree.Children = append(subtree.Children, cat)
		}

	}

	// At the end of the day, now, the tree is fully populated
	// return the root node for the entire tree
	return subtrees[rootID]
}

func scanIntoCategory(row *sql.Row) (*models.Category, error) {
	var category models.Category
	switch err := row.Scan(
		&category.ID,
		&category.UserID,
		&category.AccountID,
		&category.ParentID,
		&category.Name,
		&category.Description,
		&category.Icon,
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

func (s MySQLStore) DeleteCategory(id int64) error {
	query := "DELETE FROM categories WHERE id = ?"
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s MySQLStore) GetRootCategoryID(accountID int64) (int64, error) {
	query := `SELECT id FROM categories WHERE account_id = ? AND parent_id = 0 LIMIT 1`
	row := s.db.QueryRow(query, accountID)

	var id int64
	err := row.Scan(&id)

	return id, err
}

func (s MySQLStore) UpdateCategory(category *models.Category) error {
	query := "UPDATE categories SET parent_id = ?, name = ?, description = ?, icon = ?, updated_at = ? WHERE id = ?"
	_, err := s.db.Exec(
		query,
		category.ParentID,
		category.Name,
		category.Description,
		category.Icon,
		category.UpdatedAt,
		category.ID,
	)

	return err
}

func buildIconUrl(icon string) string {
	if icon == "" {
		return ""
	}
	path := os.Getenv("CATEGORY_ICON_PATH")
	return path + icon
}
