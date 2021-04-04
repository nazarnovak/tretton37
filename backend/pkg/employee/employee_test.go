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
				Imageportraiturl: "someurl",
				Name:             "Nazar",
				Office:           "Malmö",
				Github:           "nazarnovak",
				Twitter:          "nazarnovak",
				Linkedin:         "nazarnovak",
				Published:        true,
			},
			{
				Imageportraiturl: "someurl",
				Name:             "Nazar2",
				Office:           "Malmö",
				Github:           "nazarnovak2",
				Twitter:          "nazarnovak2",
				Linkedin:         "nazarnovak2",
				Published:        true,
			},
			{
				Imageportraiturl: "someurl",
				Name:             "Nazar3",
				Office:           "Malmö",
				Github:           "nazarnovak3",
				Twitter:          "nazarnovak3",
				Linkedin:         "nazarnovak3",
				Published:        true,
			},
		},
		ExpectedFilteredCount: 3,
	},
	{
		Name: "Filter 1 out (not published)",
		Employees: []Employee{
			{
				Imageportraiturl: "someurl",
				Name:             "Nazar",
				Office:           "Malmö",
				Github:           "nazarnovak",
				Twitter:          "nazarnovak",
				Linkedin:         "nazarnovak",
				Published:        true,
			},
			{
				Imageportraiturl: "someurl",
				Name:             "Nazar2",
				Office:           "Malmö",
				Github:           "nazarnovak2",
				Twitter:          "nazarnovak2",
				Linkedin:         "nazarnovak2",
				Published:        true,
			},
			{
				Imageportraiturl: "someurl",
				Name:             "Nazar3",
				Office:           "Malmö",
				Github:           "nazarnovak3",
				Twitter:          "nazarnovak3",
				Linkedin:         "nazarnovak3",
				Published:        false,
			},
		},
		ExpectedFilteredCount: 2,
	},
	{
		Name: "Filter 2 out (not published + no portrait url)",
		Employees: []Employee{
			{
				Imageportraiturl: "someurl",
				Name:             "Nazar",
				Office:           "Malmö",
				Github:           "nazarnovak",
				Twitter:          "nazarnovak",
				Linkedin:         "nazarnovak",
				Published:        true,
			},
			{
				Imageportraiturl: "",
				Name:             "Nazar2",
				Office:           "Malmö",
				Github:           "nazarnovak2",
				Twitter:          "nazarnovak2",
				Linkedin:         "nazarnovak2",
				Published:        true,
			},
			{
				Imageportraiturl: "someurl",
				Name:             "Nazar3",
				Office:           "Malmö",
				Github:           "nazarnovak3",
				Twitter:          "nazarnovak3",
				Linkedin:         "nazarnovak3",
				Published:        false,
			},
		},
		ExpectedFilteredCount: 1,
	},
}

func TestGetPublishedEmployees(t *testing.T) {
	for _, et := range employeesTests {
		employees := FilterEmployees(et.Employees)

		if len(employees) != et.ExpectedFilteredCount {
			t.Errorf("Test name: %s. Expecting employees: %d, got: %d", et.Name, et.ExpectedFilteredCount, len(employees))
		}
	}
}
