package api

//change things from the db model if not all info should be passed or it needs a different structure

import "check_republic/internal/models"

type ChecklistResponse struct {
	// models.Checklist					`json:"checklist"`
	ID			int						`json:"id"`
	Name      	string    	           	`json:"name"`
	Complete  	bool      	           	`json:"complete"`
	Archived  	bool      	           	`json:"archived"`
	TemplateID	int       	           	`json:"template_id"`
	Created   	string    	           	`json:"created"`
	Updated   	string    	           	`json:"updated"`
	IsChild		bool					`json:"is_child"`
	Items		[]models.ChecklistItem	`json:"items"`
	Children	[]ChecklistResponse		`json:"children"`
}
type NewChecklistRequest struct {
	Name	string	`json:"name"`
}

type ListChecklistsResponse struct {
	Checklists	[]models.Checklist	`json:"checklists"`
}


type NewChecklistResponse struct {
	ID	int	`json:"id"`
}

type NewItemRequest struct {
	ChecklistID	int		`json:"checklist_id"`
	Title		string	`json:"title"`
	Description	string	`json:"description"`
}

type NewItemResponse struct {
	ID	int	`json:"id"`
}