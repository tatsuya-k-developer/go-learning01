package repositories

import "go-learning01/internal/models"

type IItemRepository interface {
	GetAll() []models.Item
	FindAll(name string) []models.Item
	FindById(id int) (models.Item, error)

	Save(item models.Item) (models.Item, error)
	Delete(id int) bool
}

type ItemRepository struct{}

func NewItemRepository() ItemRepository {
	return ItemRepository{}
}

func (self ItemRepository) GetAll() []models.Item {
	return []models.Item{}
}

func (self ItemRepository) FindAll(name string) []models.Item {
	return []models.Item{}
}

func (self ItemRepository) FindById(id int) (models.Item, error) {
	return models.Item{ID: 100, Name: "test", Price: 10000}, nil
}

func (self ItemRepository) Save(item models.Item) (models.Item, error) {
	return models.Item{ID: 100, Name: "test", Price: 10000}, nil
}

func (self ItemRepository) Delete(id int) bool {
	return true
}
