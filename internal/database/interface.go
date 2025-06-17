package database

type DataStore interface {
	GetChecklistByID(id int) (string, error)
}
