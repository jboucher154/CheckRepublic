package api

import (
	"bytes"
	// "check_republic/internal/api"
	"check_republic/internal/database"
	"check_republic/internal/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetChecklistHandler(t *testing.T) {

	mockDB := database.NewMockDB()

	server := &Server{DB: mockDB}
	t.Run("returns checklist with id 1", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/checklist?id=1", nil)
		response := httptest.NewRecorder()
		// send request
		server.GetChecklistHandler(response, request)
		//check responses
		var got ChecklistResponse
		if err := json.NewDecoder(response.Body).Decode(&got);err != nil {
			t.Fatalf("invalid JSON response: %v", err)
		}
		want := ChecklistResponse{
				ID:       1,
				Name:     "Leaving House",
				Complete: false,
				Archived: false,
				TemplateID: 1,
				Created: "now",
				Updated: "now",
				Items: []models.ChecklistItem{
					{ID: 101, Title: "Keys", Description: "get keys", Complete: false, ChecklistId: 1, Created: "now", Updated:  "now"},
					{ID: 102, Title: "Wallet", Description: "in pocket", Complete: false, ChecklistId: 1, Created: "now", Updated:  "now"},
					{ID: 103, Title: "Phone", Description: "charged", Complete: false, ChecklistId: 1, Created: "now", Updated:  "now"},
				},
				Children: nil,
			}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("response mismatch: got %+v, want %+v", got, want)
		}
		assertStatus(t, response.Code, http.StatusOK)
	})
	t.Run("returns checklist with id 2", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/checklist?id=2", nil)
		response := httptest.NewRecorder()
		// send request
		server.GetChecklistHandler(response, request)
		//check responses
				var got ChecklistResponse
		if err := json.NewDecoder(response.Body).Decode(&got);err != nil {
			t.Fatalf("invalid JSON response: %v", err)
		}
		want := ChecklistResponse{
				ID:       2,
				Name:     "Bedtime",
				Complete: false,
				Archived: false,
				TemplateID: 2,
				Created: "now",
				Updated: "now",
				Items: []models.ChecklistItem{
					{ID: 2, Title: "Mosturize", Description: "apply lotion", Complete: false, ChecklistId: 2, Created: "now", Updated:  "now"},
				},
				Children: nil,
			}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("response mismatch: got %+v, want %+v", got, want)
		}
		assertStatus(t, response.Code, http.StatusOK)
	})
	t.Run("test bad request", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/checklist?id=fs", nil)
		response := httptest.NewRecorder()
		// send request
		server.GetChecklistHandler(response, request)
		//check responses
		got := response.Body.String()
		want := "invalid id\n"

		assertResponseString(t, got, want)
		assertStatus(t, response.Code, http.StatusBadRequest)
	})
	t.Run("test missing id", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/checklist", nil)
		response := httptest.NewRecorder()
		// send request
		server.GetChecklistHandler(response, request)
		//check responses
		got := response.Body.String()
		want := "missing id\n"

		assertResponseString(t, got, want)
		assertStatus(t, response.Code, http.StatusBadRequest)
	})
	t.Run("test id not found", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/checklist?id=999", nil)
		response := httptest.NewRecorder()
		// send request
		server.GetChecklistHandler(response, request)
		//check responses
		got := response.Body.String()
		want := "id not found\n"
		assertResponseString(t, got, want)
		assertStatus(t, response.Code, http.StatusNotFound)
	})
}


func TestCreateChecklistHandler(t *testing.T) {

	mock := database.NewMockDB()
	server := &Server{DB:mock}

	emptyNewChecklist := NewChecklistRequest{
		Name: "New empty list",
	}

	t.Run("create new empty checklist", func(t *testing.T) {
		//prepare body
		bodyBytes, err := json.Marshal(emptyNewChecklist)
		if err != nil {
			t.Fatalf("failed to marshl request: %v", err)
		}
		request, _ := http.NewRequest(http.MethodPost,"/checklist", bytes.NewReader(bodyBytes))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()

		server.CreateChecklistHandler(response, request)

	})

}

/*   HELPER FUNCTIONS   */

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func assertResponseString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
