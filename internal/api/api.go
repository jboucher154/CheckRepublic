package api

import (
	"net/http"
	"fmt"
)

func ChecklistHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "this is a checklist")
}