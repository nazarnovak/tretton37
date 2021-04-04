package employee

// GetPublishedEmployees filters out employees and gets only the ones that have published field set to true. This function could
// be extended to be a bigger filter if needed.
func GetPublishedEmployees(employees []Employee) []Employee {
	filteredEmployees := []Employee{}

	for _, e := range employees {
		if !e.Published {
			continue
		}

		filteredEmployees = append(filteredEmployees, e)
	}

	return filteredEmployees
}
