package api

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
	Items		[]models.ChecklistItem	`json:"items"`
	Children	[]ChecklistResponse		`json:"children"`
}

type NewChecklistRequest struct {
	Name	string	`json:"name"`
}

type NewChecklistResponse struct {
	ID	int	`json:"id"`
}