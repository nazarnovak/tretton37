package employee

// FilterEmployees filters out employees and gets only the ones that have published field set to true and a portrait picture.
// This function could be extended to be a bigger filter if needed.
func FilterEmployees(employees []Employee) []Employee {
	filteredEmployees := []Employee{}

	for _, e := range employees {
		if !e.Published {
			continue
		}

		if e.Imageportraiturl == "" {
			continue
		}

		filteredEmployees = append(filteredEmployees, e)
	}

	return filteredEmployees
}
