package service

import (
	"sirclo/api/entities"
	"sirclo/api/repository"
)

type ServiceCategory interface {
	ServiceCategoriesGet() ([]entities.Category, error)
}

type serviceCategory struct {
	repo repository.RepositoryCategory
}

func NewCategoryService(repo repository.RepositoryCategory) *serviceCategory {
	return &serviceCategory{repo: repo}
}

// get all category
func (sc *serviceCategory) ServiceCategoriesGet() ([]entities.Category, error) {
	categories, err := sc.repo.GetCategories()
	if err != nil {
		return categories, err
	}
	return categories, nil
}
