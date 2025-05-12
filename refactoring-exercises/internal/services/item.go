package services

import (
	"errors"
	"myapi/internal/models"
)

func CreateItem(item *models.Item) (*models.Item, error) {
	if item.Nome == "" {
		return nil, errors.New("name cannot be empty")
	}

	return item, nil
}
