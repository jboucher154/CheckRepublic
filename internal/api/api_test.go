package api

import (
	"check_republic/internal/database"
	"net/http"
	"net/http/httptest"
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
		got := response.Body.String()
		want := "checklist 1\n"

		assertResponseString(t, got, want)
		assertStatus(t, response.Code, http.StatusOK)
	})
	t.Run("returns checklist with id 2", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/checklist?id=2", nil)
		response := httptest.NewRecorder()
		// send request
		server.GetChecklistHandler(response, request)
		//check responses
		got := response.Body.String()
		want := "checklist 2\n"

		assertResponseString(t, got, want)
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
