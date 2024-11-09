package service

import (
	"go-learning01/internal/models"
	"go-learning01/internal/repositories"
	"math/rand"
)

type IItemService interface {
	GetAll() []models.Item
	FindAll(name string) []models.Item
	FindById(id int) (models.Item, error)

	Create(name string, price int) (models.Item, error)
	UpdateName(id int, name string) (models.Item, error)
	Delete(id int) bool
}

type ItemService struct {
	repo repositories.IItemRepository
}

func NewItemService(repo repositories.IItemRepository) *ItemService {
	return &ItemService{
		repo: repo,
	}
}

func (self ItemService) GetAll() []models.Item {
	return self.repo.GetAll()
}

func (self ItemService) FindAll(name string) []models.Item {
	return self.repo.FindAll(name)
}

func (self ItemService) FindById(id int) (models.Item, error) {
	return self.repo.FindById(id)
}

func (self ItemService) Create(name string, price int) (models.Item, error) {

	model := models.Item{
		ID:    rand.Int(),
		Name:  name,
		Price: price,
	}

	return self.repo.Save(model)
}

func (self ItemService) UpdateName(id int, name string) (models.Item, error) {
	model, err := self.repo.FindById(id)

	if err != nil {
		return model, err
	}

	model.Name = name

	return self.repo.Save(model)
}

func (self ItemService) Delete(id int) bool {
	return self.repo.Delete(id)
}
