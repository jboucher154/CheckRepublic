package database

import (
	"check_republic/internal/models"
	"fmt"
)

type MockDB struct {
	// data map[int]string
	Checklists      map[int]models.Checklist
	Items           map[int]models.ChecklistItem
	ChecklistToItem map[int][]int
}

func NewMockDB() *MockDB {
	return &MockDB{
		Checklists: map[int]models.Checklist{
			1: {
				ID:         1,
				Name:       "Leaving House",
				Complete:   false,
				Archived:   false,
				TemplateId: 1,
				Created:    "now",
				Updated:    "now",
			},
			2: {
				ID:         2,
				Name:       "Bedtime",
				Complete:   false,
				Archived:   false,
				TemplateId: 2,
				Created:    "now",
				Updated:    "now",
			},
		},
		Items: map[int]models.ChecklistItem{
			101: {
				ID:          101,
				Title:       "Keys",
				Description: "get keys",
				Complete:    false,
				ChecklistId: 1,
				Created:     "now",
				Updated:     "now",
			},
			102: {
				ID:          102,
				Title:       "Wallet",
				Description: "in pocket",
				Complete:    false,
				ChecklistId: 1,
				Created:     "now",
				Updated:     "now",
			},
			103: {
				ID:          103,
				Title:       "Phone",
				Description: "charged",
				Complete:    false,
				ChecklistId: 1,
				Created:     "now",
				Updated:     "now",
			},
			2: {
				ID:          2,
				Title:       "Mosturize",
				Description: "apply lotion",
				Complete:    false,
				ChecklistId: 2,
				Created:     "now",
				Updated:     "now",
			},
		},
		ChecklistToItem: map[int][]int{
			1:{101, 102, 103},
			2:{2},
		},
	}
}


func (m *MockDB) GetChecklistByID(id int) (models.Checklist, []models.ChecklistItem, error) {
	checklist, err := m.Checklists[id]
	if !err {
		return models.Checklist{}, nil, fmt.Errorf("not found")
	}
	itemIDs := m.ChecklistToItem[id]
	items := []models.ChecklistItem{}
	for _, itemID := range itemIDs {
		if item, ok := m.Items[itemID]; ok {
			items = append(items, item)
		}
	}
	return checklist, items, nil
}
