package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)
	


func TestGETChecklist(t *testing.T) {
	t.Run("returns checklist with id 1", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/checklist/1", nil)
		response := httptest.NewRecorder()
		// send request
		ChecklistHandler(response, request)
		//check responses
		got := response.Body.String()
		want := "this is a checklist"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})
}