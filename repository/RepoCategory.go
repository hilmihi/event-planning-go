package repository

import (
	"database/sql"
	"sirclo/api/entities"
)

type RepositoryCategory interface {
	GetCategories() ([]entities.Category, error)
}

type Repository_Category struct {
	db *sql.DB
}

func NewRepositoryCategory(db *sql.DB) *Repository_Category {
	return &Repository_Category{db: db}
}

// get category
func (rc *Repository_Category) GetCategories() ([]entities.Category, error) {
	var categories []entities.Category
	result, err := rc.db.Query(`select id, description from category where deleted_at is null order by id asc`)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	for result.Next() {
		var category entities.Category

		err = result.Scan(&category.Id, &category.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}
