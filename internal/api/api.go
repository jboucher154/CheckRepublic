package api

import (
	"check_republic/internal/database"
	"fmt"
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
	checklist, err := s.DB.GetChecklistByID(id)
	if err != nil {
		http.Error(w, "id not found", http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, checklist)
}

// func (p *ChecklistServer) ChecklistHandler(w http.ResponseWriter, r *http.Request) {

// 	checklist, _ := strconv.Atoi(r.PathValue("id"))
// 	fmt.Fprint(w, p.store.GetChecklistByID(checklist))
// }
