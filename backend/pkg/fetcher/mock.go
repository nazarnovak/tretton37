package fetcher

import "github.com/nazarnovak/tretton37/backend/pkg/employee"

type FetcherMock struct{}

func (fm FetcherMock) GetAll() ([]employee.Employee, error) {
	return []employee.Employee{
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
		{
			Name:        "Nazar3",
			Email:       "nazar3@nazar.com",
			Phonenumber: "+466791112233",
			Office:      "Malmö",
			Github:      "nazarnovak3",
			Twitter:     "nazarnovak3",
			Linkedin:    "nazarnovak3",
			Published:   true,
		},
		{
			Name:        "Nazar4",
			Email:       "nazar4@nazar.com",
			Phonenumber: "+466791112233",
			Office:      "Malmö",
			Github:      "nazarnovak4",
			Twitter:     "nazarnovak4",
			Linkedin:    "nazarnovak4",
			Published:   false,
		},
	}, nil
}
