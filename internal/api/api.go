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
		ID: checklist.ID,
		Name: checklist.Name,
		Archived: checklist.Archived,
		TemplateID: checklist.TemplateId,
		Complete: checklist.Complete,
		Created: checklist.Created,
		Updated: checklist.Updated,
		Items: items,
		Children: nil,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *Server) CreateChecklistHandler(w http.ResponseWriter, r *http.Request) {

}