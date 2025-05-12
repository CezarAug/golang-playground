package services

import (
	"fmt"
	"stock/internal/models"
	"strconv"
	"time"
)

type Stock struct {
	items map[string]models.Item
	logs  []models.Log
}

func NewStock() *Stock {
	return &Stock{
		items: make(map[string]models.Item),
		logs:  []models.Log{},
	}
}

func (s *Stock) AddItem(item models.Item, user string) error {

	if item.Price <= 0 || item.Quantity <= 0 {
		return fmt.Errorf("failed to add item: [ID: %d]", item.ID)
	}

	existingItem, exists := s.items[strconv.Itoa(item.ID)]
	if exists {
		item.Quantity += existingItem.Quantity
	}

	s.items[strconv.Itoa(item.ID)] = item

	s.logs = append(s.logs, models.Log{
		Timestamp: time.Now(),
		Action:    "New item in stock",
		User:      user,
		ItemId:    item.ID,
		Quantity:  item.Quantity,
		Reason:    "Adding new items in stock",
	})

	return nil
}

func (s *Stock) ListItems() []models.Item {
	var itemList []models.Item
	for _, item := range s.items {
		itemList = append(itemList, item)
	}

	return itemList
}

func (s *Stock) RemoveItem(item models.Item, user string) error {
	existingItem, exists := s.items[strconv.Itoa(item.ID)]
	if item.Quantity <= 0 {
		return fmt.Errorf("invalid quantity: ", item.Quantity)
	}

	if !exists {
		return fmt.Errorf("item does not exist: [ID: %d]", item.ID)
	}

	if existingItem.Quantity < item.Quantity {
		return fmt.Errorf("cannot remove more items than current stock, quantity: ", item.Quantity)
	}

	existingItem.Quantity = existingItem.Quantity - item.Quantity
	s.items[strconv.Itoa(item.ID)] = existingItem

	s.logs = append(s.logs, models.Log{
		Timestamp: time.Now(),
		Action:    "Removed items",
		User:      user,
		ItemId:    item.ID,
		Quantity:  item.Quantity,
		Reason:    "Removing items in stock",
	})

	return nil
}

func (s *Stock) ViewAuditLogs() []models.Log {
	return s.logs
}

func (s *Stock) CalculateTotalCost() float64 {
	var totalCost float64

	for _, item := range s.items {
		totalCost += float64(item.Quantity) * item.Price
	}

	return totalCost
}

func Find[T any](data []T, comparator func(T) bool) ([]T, error) {
	var result []T
	for _, v := range data {
		if comparator(v) {
			result = append(result, v)
		}
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("no items were found")
	}
	return result, nil
}
