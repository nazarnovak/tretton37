package fetcher

import "github.com/nazarnovak/tretton37/backend/pkg/employee"

type FetcherMock struct{}

func (fm FetcherMock) GetAll() ([]employee.Employee, error) {
	return []employee.Employee{
		{
			Name:      "Nazar",
			Office:    "Malmö",
			Github:    "nazarnovak",
			Twitter:   "nazarnovak",
			Linkedin:  "nazarnovak",
			Published: true,
		},
		{
			Name:      "Nazar2",
			Office:    "Malmö",
			Github:    "nazarnovak2",
			Twitter:   "nazarnovak2",
			Linkedin:  "nazarnovak2",
			Published: true,
		},
		{
			Name:      "Nazar3",
			Office:    "Malmö",
			Github:    "nazarnovak3",
			Twitter:   "nazarnovak3",
			Linkedin:  "nazarnovak3",
			Published: true,
		},
		{
			Name:      "Nazar4",
			Office:    "Malmö",
			Github:    "nazarnovak4",
			Twitter:   "nazarnovak4",
			Linkedin:  "nazarnovak4",
			Published: false,
		},
	}, nil
}
