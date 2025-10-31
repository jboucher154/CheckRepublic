package database

import (
	"check_republic/internal/models"
)

type DataStore interface {
	GetChecklistByID(id int) (models.Checklist, []models.ChecklistItem, error)
	GetChecklists(userID string) ([]models.Checklist, error)
	GetChecklistItems(checklistID int) ([]models.ChecklistItem, error)
	// GetChecklistTitles(userID string) ([]string, error)

	CreateChecklist(name string) (int, error)
	CreateChecklistItem(checklistId int, name string, description string) (int, error)
	// CreateChecklistWithItems(name string, items []models.ChecklistItem) (int, error)

	UpdateChecklist(id int, updateInfo map[string]string) (models.Checklist, error)
	UpdateChecklistItem(id int, updateInfo map[string]string) (models.ChecklistItem, error)

	// DeleteChecklist(id int, deleteAllChildren bool) error
	// DeleteChecklistItem(id int) error
}
