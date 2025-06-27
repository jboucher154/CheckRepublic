package database

import (
	"check_republic/internal/models"
)

type DataStore interface {
	GetChecklistByID(id int) (models.Checklist, []models.ChecklistItem, error)
}
