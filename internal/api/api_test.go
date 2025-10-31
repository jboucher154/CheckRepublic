package api

import (
	"bytes"
	"check_republic/internal/database"
	"check_republic/internal/models"
	"encoding/json"
	"fmt"
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
		if err := json.NewDecoder(response.Body).Decode(&got); err != nil {
			t.Fatalf("invalid JSON response: %v", err)
		}
		want := ChecklistResponse{
			ID:         1,
			Name:       "Leaving House",
			Complete:   false,
			Archived:   false,
			TemplateID: 1,
			Created:    "now",
			Updated:    "now",
			Items: []models.ChecklistItem{
				{ID: 101, Title: "Keys", Description: "get keys", Complete: false, ChecklistId: 1, Created: "now", Updated: "now"},
				{ID: 102, Title: "Wallet", Description: "in pocket", Complete: false, ChecklistId: 1, Created: "now", Updated: "now"},
				{ID: 103, Title: "Phone", Description: "charged", Complete: false, ChecklistId: 1, Created: "now", Updated: "now"},
			},
			Children: nil,
			IsChild:  false,
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
		if err := json.NewDecoder(response.Body).Decode(&got); err != nil {
			t.Fatalf("invalid JSON response: %v", err)
		}
		want := ChecklistResponse{
			ID:         2,
			Name:       "Bedtime",
			Complete:   false,
			Archived:   false,
			TemplateID: 2,
			Created:    "now",
			Updated:    "now",
			Items: []models.ChecklistItem{
				{ID: 2, Title: "Moisturize", Description: "apply lotion", Complete: false, ChecklistId: 2, Created: "now", Updated: "now"},
			},
			Children: nil,
			IsChild:  false,
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

func TestGetChecklistsHandler(t *testing.T) {
	mockDB := database.NewMockDB()

	server := &Server{DB: mockDB}
	t.Run("returns all checklists that are not children", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/checklists", nil)
		response := httptest.NewRecorder()
		// send request
		server.GetChecklistsHandler(response, request)
		//check responses
		var got ListChecklistsResponse
		if err := json.NewDecoder(response.Body).Decode(&got); err != nil {
			t.Fatalf("invalid JSON response: %v", err)
		}
		want := ListChecklistsResponse{
			Checklists: []models.Checklist{
				{
					ID:         1,
					Name:       "Leaving House",
					Complete:   false,
					Archived:   false,
					TemplateId: 1,
					Created:    "now",
					Updated:    "now",
					IsChild:    false,
				},
				{
					ID:         2,
					Name:       "Bedtime",
					Complete:   false,
					Archived:   false,
					TemplateId: 2,
					Created:    "now",
					Updated:    "now",
					IsChild:    false,
				},
			},
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("response mismatch: got %+v, want %+v", got, want)
		}
		assertStatus(t, response.Code, http.StatusOK)
	})
	t.Run("request checklists from empty DB", func(t *testing.T) {
		emptyMockDB := database.NewEmptyMockDB()
		emptyServer := &Server{DB: emptyMockDB}

		request, _ := http.NewRequest(http.MethodGet, "/checklists", nil)
		response := httptest.NewRecorder()
		// send request
		emptyServer.GetChecklistsHandler(response, request)
		//check responses
		got := response.Body.String()
		want := "no checklists found for user\n"

		assertStatus(t, response.Code, http.StatusNotFound)
		assertResponseString(t, got, want)
	})
	//TODO need to test user based retrieval eventually
}

func TestCreateChecklistHandler(t *testing.T) {

	mock := database.NewMockDB()
	server := &Server{DB: mock}

	emptyNewChecklist := NewChecklistRequest{
		Name: "New empty list",
	}

	t.Run("create new empty checklist", func(t *testing.T) {
		//prepare body
		bodyBytes, err := json.Marshal(emptyNewChecklist)
		if err != nil {
			t.Fatalf("failed to marshal request: %v", err)
		}
		request, _ := http.NewRequest(http.MethodPost, "/checklist", bytes.NewReader(bodyBytes))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()

		server.CreateChecklistHandler(response, request)
		assertStatus(t, response.Code, http.StatusOK)
		var checklist_res NewChecklistResponse
		err = json.NewDecoder(response.Body).Decode(&checklist_res)
		if err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}
		id := checklist_res.ID
		path := fmt.Sprintf("/checklist?id=%d", id)
		//request the id to check that exists in db
		c_request, _ := http.NewRequest(http.MethodGet, path, nil)
		c_response := httptest.NewRecorder()
		// send request
		server.GetChecklistHandler(c_response, c_request)
		//check responses
		assertStatus(t, c_response.Code, http.StatusOK)

		// check that the name is the same as the one in the request
		var got ChecklistResponse
		if err := json.NewDecoder(c_response.Body).Decode(&got); err != nil {
			t.Fatalf("invalid JSON response: %v", err)
		}
		want := ChecklistResponse{
			ID:         3,
			Name:       "New empty list",
			Complete:   false,
			Archived:   false,
			TemplateID: 0,
			Created:    "now",
			Updated:    "now",
			Items:      []models.ChecklistItem{},
			Children:   nil,
			IsChild:    false,
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("response mismatch: got %+v, want %+v", got, want)
		}
	})
}

func TestGetItemsHandler(t *testing.T) {
	mock := database.NewMockDB()
	server := &Server{DB: mock}
	t.Run("get items from list id 2", func(t *testing.T) {
		//prepare body
		request, _ := http.NewRequest(http.MethodGet, "/checklist-items?id=2", nil)
		response := httptest.NewRecorder()

		server.GetItemsHandler(response, request)
		assertStatus(t, response.Code, http.StatusOK)
		var got GetItemsResponse
		err := json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}
		// check that the name is the same as the one in the request

		want := GetItemsResponse{
			Items: []models.ChecklistItem{
				{
					ID:          2,
					Title:       "Moisturize",
					Description: "apply lotion",
					Complete:    false,
					ChecklistId: 2,
					Created:     "now",
					Updated:     "now",
				},
			},
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("response mismatch: got %+v, want %+v", got, want)
		}
	})
	t.Run("Test request for items when there are none", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/checklist-items?id=3", nil)
		response := httptest.NewRecorder()

		server.GetItemsHandler(response, request)
		assertStatus(t, response.Code, http.StatusNotFound)
		want := "no items found for checklist\n"
		got := response.Body.String()
		assertResponseString(t, got, want)
	})

}

func TestCreateItemHandler(t *testing.T) {

	mock := database.NewMockDB()
	server := &Server{DB: mock}

	newItem := NewItemRequest{
		ChecklistID: 3,
		Title:       "new item",
		Description: "just a new one",
	}
	t.Run("create new item for checklist", func(t *testing.T) {
		//prepare body
		bodyBytes, err := json.Marshal(newItem)
		if err != nil {
			t.Fatalf("failed to marshal request: %v", err)
		}
		request, _ := http.NewRequest(http.MethodPost, "/checklist-item", bytes.NewReader(bodyBytes))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()

		server.CreateItemHandler(response, request)
		assertStatus(t, response.Code, http.StatusOK)
		var item_res NewItemResponse
		err = json.NewDecoder(response.Body).Decode(&item_res)
		if err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}
		// get id to compare to items for checklist
		id := item_res.ID
		// path := fmt.Sprintf("/checklist-items?id=%d", id)
		//request items for the checklist, should be only item
		c_request, _ := http.NewRequest(http.MethodGet, "/checklist-item?id=3", nil)
		c_response := httptest.NewRecorder()
		// send request
		server.GetItemsHandler(c_response, c_request) //need this handler
		//check responses
		assertStatus(t, c_response.Code, http.StatusOK)

		var got GetItemsResponse
		if err := json.NewDecoder(c_response.Body).Decode(&got); err != nil {
			t.Fatalf("invalid JSON response: %v", err)
		}
		want := GetItemsResponse{
			Items: []models.ChecklistItem{
				{
					Title:       "new item",
					Description: "just a new one",
					ID:          id,
					Complete:    false,
					ChecklistId: 3,
					Created:     "now",
					Updated:     "now",
				},
			},
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("response mismatch: got %+v, want %+v", got, want)
		}
	})

}

func TestUpdateItemHandler(t *testing.T) {

	mock := database.NewMockDB()
	server := &Server{DB: mock}

	title := "updated"
	description := "updated"
	complete := true
	id := 2

	itemUpdates := UpdateItemRequest{
		Id:          &id,
		Title:       &title,
		Description: &description,
		Complete:    &complete,
	}
	t.Run("update item for checklist", func(t *testing.T) {
		//prepare body
		bodyBytes, err := json.Marshal(itemUpdates)
		if err != nil {
			t.Fatalf("failed to marshal request: %v", err)
		}
		request, _ := http.NewRequest(http.MethodPatch, "/checklist-item", bytes.NewReader(bodyBytes))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()

		server.UpdateItemHandler(response, request)
		assertStatus(t, response.Code, http.StatusOK)

		var got UpdateItemResponse
		err = json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}
		want := UpdateItemResponse{
			Item: models.ChecklistItem{
				ID:          2,
				Title:       title,
				Description: description,
				Complete:    complete,
				ChecklistId: 2,
				Created:     "now",
				Updated:     "now",
			},
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("response mismatch: got %+v, want %+v", got, want)
		}

		// retrieve same and compare as well
		c_request, _ := http.NewRequest(http.MethodGet, "/checklist-item?id=2", nil)
		c_response := httptest.NewRecorder()
		// send request
		server.GetItemsHandler(c_response, c_request) //need this handler
		// //check responses
		assertStatus(t, c_response.Code, http.StatusOK)

		var c_got GetItemsResponse
		if err := json.NewDecoder(c_response.Body).Decode(&c_got); err != nil {
			t.Fatalf("invalid JSON response: %v", err)
		}
		c_want := GetItemsResponse{
			Items: []models.ChecklistItem{want.Item},
		}

		if !reflect.DeepEqual(c_want, c_got) {
			t.Errorf("response mismatch: got %+v, want %+v", got, want)
		}
	})

}

func TestUpdateChecklistHandler(t *testing.T) {

	mock := database.NewMockDB()
	server := &Server{DB: mock}

	title := "updated"
	isChild := false
	complete := true
	id := 3

	checklistUpdates := UpdateChecklistRequest{
		Id:       &id,
		Name:     &title,
		IsChild:  &isChild,
		Complete: &complete,
	}
	t.Run("update checklist", func(t *testing.T) {
		//prepare body
		bodyBytes, err := json.Marshal(checklistUpdates)
		if err != nil {
			t.Fatalf("failed to marshal request: %v", err)
		}
		request, _ := http.NewRequest(http.MethodPatch, "/checklist", bytes.NewReader(bodyBytes))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()

		server.UpdateChecklistHandler(response, request)
		assertStatus(t, response.Code, http.StatusOK)

		var got UpdateChecklistResponse
		err = json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}
		want := UpdateChecklistResponse{
			Checklist: models.Checklist{
				ID:         3,
				Name:       "updated",
				Complete:   true,
				Archived:   false,
				TemplateId: 3,
				Created:    "now",
				Updated:    "now",
				IsChild:    false,
			},
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("response mismatch: got %+v, want %+v", got, want)
		}

		// retrieve same and compare as well
		c_request, _ := http.NewRequest(http.MethodGet, "/checklist?id=3", nil)
		c_response := httptest.NewRecorder()
		// send request
		server.GetChecklistHandler(c_response, c_request) //need this handler
		// //check responses
		assertStatus(t, c_response.Code, http.StatusOK)

		var c_got ChecklistResponse
		if err := json.NewDecoder(c_response.Body).Decode(&c_got); err != nil {
			t.Fatalf("invalid JSON response: %v", err)
		}
		c_want := ChecklistResponse{
			ID:       3,
			Name:     "updated",
			Complete: true,
			Archived: false,
			TemplateID: 3,
			Created:  "now",
			Updated:  "now",
			IsChild:  false,
			Items:    []models.ChecklistItem{},
			Children: nil,
		}

		if !reflect.DeepEqual(c_want, c_got) {
			t.Errorf("response mismatch: got %+v, want %+v", got, want)
		}
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
