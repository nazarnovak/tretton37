package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nazarnovak/tretton37/backend/pkg/employee"
)

type EmployeeGetter interface {
	GetAll() ([]employee.Employee, error)
}

func EmployeesHandler(eg EmployeeGetter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// First we get all employees from the API
		employees, err := eg.GetAll()
		if err != nil {
			fmt.Println(fmt.Errorf("Error fetching employees: %w", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Now we need to filter some of the employees
		filteredEmployees := employee.FilterEmployees(employees)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(filteredEmployees)
	}
}
