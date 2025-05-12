package repositories

import (
	"myapi/internal/config"
	"myapi/internal/models"
)

type ItemRepository struct{}

func NewItemRepository() *ItemRepository {
	return &ItemRepository{}
}

func (r *ItemRepository) ListAll() ([]models.Item, error) {
	var items []models.Item
	if err := config.DB.Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (r *ItemRepository) FindById(id int) (models.Item, error) {
	var item models.Item
	if err := config.DB.First(&item, id).Error; err != nil {
		return item, err
	}

	return item, nil
}
