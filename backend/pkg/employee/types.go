package employee

// Employee is the struct that handles the employees we show to the user in our app.
type Employee struct {
	Name      string `json:"name"`
	Office    string `json:"office"`
	Github    string `json:"github"`
	Linkedin  string `json:"linkedin"`
	Twitter   string `json:"twitter"`
	Published bool   `json:"published"`
}
