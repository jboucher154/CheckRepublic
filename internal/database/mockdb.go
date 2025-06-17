package database

import "fmt"

type MockDB struct {
	data map[int]string
}

func NewMockDB() *MockDB {
	return &MockDB{
		data: map[int]string{
			1: "checklist 1",
			2: "checklist 2",
			3: "checklist 3",
		},
	}
}

func (m *MockDB) GetChecklistByID(id int) (string, error) {
	item, err := m.data[id]
	if !err { //is this err eval wrong?
		return "", fmt.Errorf("not found")
	}
	return item, nil
}
