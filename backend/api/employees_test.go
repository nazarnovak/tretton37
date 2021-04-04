package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zenazn/goji/web"

	"github.com/nazarnovak/tretton37/backend/pkg/employee"
	"github.com/nazarnovak/tretton37/backend/pkg/fetcher"
)

func TestEmployees(t *testing.T) {
	mux := web.New()
	mux.Get("/api/employees", EmployeesHandler(fetcher.FetcherMock{}))

	req, _ := http.NewRequest("GET", "/api/employees", nil)
	rr := httptest.NewRecorder()

	mux.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expecting status: %d got: %d", http.StatusOK, rr.Code)
	}

	var employees []employee.Employee
	if err := json.Unmarshal(rr.Body.Bytes(), &employees); err != nil {
		t.Errorf("Unexpected error: %w", err)
	}

	if len(employees) != 3 {
		t.Errorf("Unexpected employees count. Expecting: %d, got: %d", 3, len(employees))
	}
}
