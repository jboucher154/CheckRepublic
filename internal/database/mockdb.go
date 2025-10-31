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
	NextID          int
	NextItemID      int
}

func NewEmptyMockDB() *MockDB {
	return &MockDB{}
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
				IsChild:    false,
			},
			2: {
				ID:         2,
				Name:       "Bedtime",
				Complete:   false,
				Archived:   false,
				TemplateId: 2,
				Created:    "now",
				Updated:    "now",
				IsChild:    false,
			},
			3: {
				ID:         3,
				Name:       "I'm a child",
				Complete:   false,
				Archived:   false,
				TemplateId: 3,
				Created:    "now",
				Updated:    "now",
				IsChild:    true,
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
				Title:       "Moisturize",
				Description: "apply lotion",
				Complete:    false,
				ChecklistId: 2,
				Created:     "now",
				Updated:     "now",
			},
		},
		ChecklistToItem: map[int][]int{
			1: {101, 102, 103},
			2: {2},
		},
		NextID:     3,
		NextItemID: 104,
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

// returns Checklist info only
// need to not return checklists that are children...
func (m *MockDB) GetChecklists(userID string) ([]models.Checklist, error) {
	var checklists []models.Checklist

	for _, checklist := range m.Checklists {
		if !checklist.IsChild {
			checklists = append(checklists, checklist)
		}
	}
	return checklists, nil
}

func (m *MockDB) CreateChecklist(name string) (int, error) {
	// title should be unique
	for _, checklist := range m.Checklists {
		if checklist.Name == name {
			return 0, fmt.Errorf("checklist already exists")
		}
	}
	// create new checklist with title
	id := m.NextID
	m.Checklists[id] = models.Checklist{
		ID:         id,
		Name:       name,
		Complete:   false,
		Archived:   false,
		TemplateId: 0,
		Created:    "now",
		Updated:    "now",
		IsChild:    false,
	}
	m.NextID++
	// return id of new checklist
	return id, nil
}

func (m *MockDB) CreateChecklistItem(checklistId int, title string, description string) (int, error) {
	itemID := m.NextItemID

	newItem := models.ChecklistItem{
		ID:          itemID,
		Title:       title,
		Description: description,
		Complete:    false,
		Created:     "now",
		Updated:     "now",
		ChecklistId: checklistId,
	}
	m.Items[itemID] = newItem
	m.NextItemID++
	m.ChecklistToItem[checklistId] = append(m.ChecklistToItem[checklistId], itemID)
	return itemID, nil
}

// TODO add getting child checklists as well (not items, just basic checklist info that can be expanded later)
func (m *MockDB) GetChecklistItems(checklistID int) ([]models.ChecklistItem, error) {
	var items []models.ChecklistItem
	itemIds, ok := m.ChecklistToItem[checklistID]

	if !ok || len(itemIds) == 0 {
		return items, ErrNotFound
	}
	//this is assuming all are items, not checklist children
	for _, itemId := range itemIds {
		items = append(items, m.Items[itemId])
	}

	return items, nil
}

func (m *MockDB) UpdateChecklistItem(id int, updateInfo map[string]string) (models.ChecklistItem, error) {
	//get item
	itemToUpdate, ok := m.Items[id]
	if !ok {
		return itemToUpdate, ErrNotFound
	}
	//update fields passed
	for key, value := range updateInfo {
		switch key {
			case "Title":
				itemToUpdate.Title = value
			case "Description":
				itemToUpdate.Description = value
			case "Complete":
				//toggle
				if updateInfo["Complete"] == "true" {
					itemToUpdate.Complete = true
				} else {
					itemToUpdate.Complete = false
				}
			default:
				return itemToUpdate, ErrIncorrectRequest
		}
	}
	//send copy of updated item
	m.Items[id] = itemToUpdate
	return itemToUpdate, nil
}

func (m *MockDB) UpdateChecklist(id int, updateInfo map[string]string) (models.Checklist, error) {
	//get item
	checklistToUpdate, ok := m.Checklists[id] //would doing by reference be better
	if !ok {
		return checklistToUpdate, ErrNotFound
	}
	//update fields passed
	for key, value := range updateInfo {
		switch key {
			case "Name":
				checklistToUpdate.Name = value
			case "IsChild":
				if updateInfo["IsChild"] == "true" {
					checklistToUpdate.IsChild = true
				} else {
					checklistToUpdate.IsChild = false
				}
			case "Complete":
				if updateInfo["Complete"] == "true" {
					checklistToUpdate.Complete = true
				} else {
					checklistToUpdate.Complete = false
				}
			case "Archived":
				if updateInfo["Archived"] == "true" {
					checklistToUpdate.Archived = true
				} else {
					checklistToUpdate.Archived = false
				}
			default:
				return checklistToUpdate, ErrIncorrectRequest
		}
	}
	//send copy of updated item
	m.Checklists[id] = checklistToUpdate
	return checklistToUpdate, nil
}

func (m *MockDB) DeleteChecklistItem(id int) error {
	item, ok := m.Items[id]
	if !ok {
		return ErrNotFound
	}
	//remove item
	delete(m.Items, id)
	//remove id from mapping
	itemIds := m.ChecklistToItem[item.ID]
	var newIdList []int

	for _, idInt := range itemIds {
		if idInt != id {
			newIdList = append(newIdList, idInt)
		}
	}
	m.ChecklistToItem[item.ID] = newIdList
	return nil
}
//not sure about the delete all children. maybe that sould always happen...
func (m *MockDB) DeleteChecklist(id int) error {
	_, ok := m.Checklists[id]
	if !ok {
		return ErrNotFound
	}
	//remove all items
	items := m.ChecklistToItem[id]
	for _, itemId := range items {
		delete(m.Items, itemId)
	}
	//TODO remove children

	// remove from item mapping list
	delete(m.ChecklistToItem, id)
	//delete checklist
	delete(m.Checklists, id)
	return nil
}