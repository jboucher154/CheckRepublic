package models

type Checklist struct {
	ID			int		`json:"id"`
	Name		string	`json:"name"`
	Complete	bool	`json:"complete"`
	Archived	bool	`json:"archived"`
	TemplateId	int		`json:"template_id"`
	Created		string	`json:"created"`
	Updated		string	`json:"updated"`
	IsChild		bool	`json:"is_child"`
}

type ChecklistItem  struct {
	ID			int		`json:"id"`
	Title		string 	`json:"title"`
	Description	string 	`json:"description"`
	Complete	bool	`json:"complete"`
	ChecklistId	int		`json:"checklist_id"`
	Created		string	`json:"created"`
	Updated		string	`json:"updated"`
}

type ChecklistItems struct {
	ID				int		`json:"id"`
	ChecklistId		int		`json:"checklist_id"`
	ItemId			int		`json:"item_id"`
	SubChecklistId	int		`json:"sub_checklist_id"`
	Created			string	`json:"created"`
	Updated			string	`json:"updated"`
}

type TemplateChecklist struct {
	ID			int		`json:"id"`
	Name		string	`json:"name"`
	Created		string	`json:"created"`
	Updated		string	`json:"updated"`
}

type TemplateChecklistItem  struct {
	ID					int		`json:"id"`
	Title				string 	`json:"title"`
	Description			string 	`json:"descritpion"`
	TemplateChecklistId	int		`json:"template_checklist_id"`
	Created				string	`json:"created"`
	Updated				string	`json:"updated"`
}

type TemplateChecklistItems struct {
	ID						int		`json:"id"`
	TemplateChecklistId		int		`json:"template_checklist_id"`
	ItemId					int		`json:"item_id"`
	TemplateSubChecklistId	int		`json:"template_sub_checklist_id"`
	Created					string	`json:"created"`
	Updated					string	`json:"updated"`
}
