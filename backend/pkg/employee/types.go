package employee

// Employee is the struct that handles the employees we show to the user in our app.
type Employee struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phonenumber string `json:"phoneNumber"`
	Office      string `json:"office"`
	Github      string `json:"gitHub"`
	Twitter     string `json:"twitter"`
	Linkedin    string `json:"linkedIn"`
	Published   bool   `json:"published"`
}
