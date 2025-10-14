package api

import (
	"check_republic/internal/database"
	"encoding/json"
	"net/http"
	"strconv"
)

type Server struct {
	DB database.DataStore
}

func (s *Server) GetChecklistHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id") //lookup how this works
	id, err := strconv.Atoi(idStr)
	if err != nil {
		if idStr == "" {
			http.Error(w, "missing id", http.StatusBadRequest)
			return
		}
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	checklist, items, err := s.DB.GetChecklistByID(id)
	if err != nil {
		http.Error(w, "id not found", http.StatusNotFound)
		return
	}
	response := ChecklistResponse{
		ID:         checklist.ID,
		Name:       checklist.Name,
		Archived:   checklist.Archived,
		TemplateID: checklist.TemplateId,
		Complete:   checklist.Complete,
		Created:    checklist.Created,
		Updated:    checklist.Updated,
		Items:      items,
		Children:   nil,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *Server) CreateChecklistHandler(w http.ResponseWriter, r *http.Request) {
	// get body, validate that it is json for new checklist
	var newChecklist NewChecklistRequest
	// unmarshall json
	if err := json.NewDecoder(r.Body).Decode(&newChecklist); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	// create checklist
	id, err := s.DB.CreateChecklist(newChecklist.Name)
	if err != nil {
		http.Error(w, "failed to create checklist", http.StatusInternalServerError)
		return
	}
	// return id of new checklist
	response := NewChecklistResponse{ID: id}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
