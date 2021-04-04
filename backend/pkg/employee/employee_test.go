package employee

import (
	"testing"
)

type employeeTest struct {
	Name                  string
	Employees             []Employee
	ExpectedFilteredCount int
}

var employeesTests = []employeeTest{
	{
		Name: "No filtering",
		Employees: []Employee{
			{
				Name:        "Nazar",
				Email:       "nazar@nazar.com",
				Phonenumber: "+466791112233",
				Office:      "Malmö",
				Github:      "nazarnovak",
				Twitter:     "nazarnovak",
				Linkedin:    "nazarnovak",
				Published:   true,
			},
			{
				Name:        "Nazar2",
				Email:       "nazar2@nazar.com",
				Phonenumber: "+466791112233",
				Office:      "Malmö",
				Github:      "nazarnovak2",
				Twitter:     "nazarnovak2",
				Linkedin:    "nazarnovak2",
				Published:   true,
			},
		},
		ExpectedFilteredCount: 2,
	},
	{
		Name: "Filter 1 out",
		Employees: []Employee{
			{
				Name:        "Nazar",
				Email:       "nazar@nazar.com",
				Phonenumber: "+466791112233",
				Office:      "Malmö",
				Github:      "nazarnovak",
				Twitter:     "nazarnovak",
				Linkedin:    "nazarnovak",
				Published:   true,
			},
			{
				Name:        "Nazar2",
				Email:       "nazar2@nazar.com",
				Phonenumber: "+466791112233",
				Office:      "Malmö",
				Github:      "nazarnovak2",
				Twitter:     "nazarnovak2",
				Linkedin:    "nazarnovak2",
				Published:   false,
			},
		},
		ExpectedFilteredCount: 1,
	},
}

func TestGetPublishedEmployees(t *testing.T) {
	for _, et := range employeesTests {
		employees := GetPublishedEmployees(et.Employees)

		if len(employees) != et.ExpectedFilteredCount {
			t.Errorf("Test name: %s. Expecting employees: %d, got: %d", et.Name, et.ExpectedFilteredCount, len(employees))
		}
	}
}
