package api

import (
	"check_republic/internal/database"
	"check_republic/internal/models"
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
	//how to add child
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
		IsChild:    checklist.IsChild,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *Server) GetChecklistsHandler(w http.ResponseWriter, r *http.Request) {
	//TODO gets user id from auth eventually
	var checklists []models.Checklist
	var checklistsWrapper ListChecklistsResponse

	checklists, err := s.DB.GetChecklists("placeholder")
	if err != nil {
		http.Error(w, "unable to retrieve checklists", http.StatusInternalServerError)
		return
	}
	if len(checklists) == 0 {
		http.Error(w, "no checklists found for user", http.StatusNotFound)
		return
	}
	checklistsWrapper.Checklists = checklists
	json.NewEncoder(w).Encode(checklistsWrapper) //should this have an error check?

}
func (s *Server) CreateChecklistHandler(w http.ResponseWriter, r *http.Request) {
	// get body, validate that it is json for new checklist
	var newChecklist NewChecklistRequest
	// unmarshal json
	if err := json.NewDecoder(r.Body).Decode(&newChecklist); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	// create checklist
	id, err := s.DB.CreateChecklist(newChecklist.Name)
	if err != nil {
		//should this be error 409 if specifically for dup name?
		http.Error(w, "failed to create checklist", http.StatusInternalServerError)
		return
	}
	// return id of new checklist
	response := NewChecklistResponse{ID: id}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

//TODO make get checklist items for checklist id

func (s *Server) CreateItemHandler(w http.ResponseWriter, r *http.Request) {
	var newItem NewItemRequest

	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, "incorrect body format for new item creation", http.StatusNotAcceptable)
		return
	}
	id, err := s.DB.CreateChecklistItem(newItem.ChecklistID, newItem.Title, newItem.Description)
	if err != nil {
		http.Error(w, "failed to create checklist item", http.StatusInternalServerError)
		return
	}
	response := NewItemResponse{ID: id}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
