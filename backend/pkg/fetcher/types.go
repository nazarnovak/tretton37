package fetcher

// Employee is the struct that handles the employees we fetch from the external API.
type Employee []struct {
	Name               string      `json:"name"`
	Email              string      `json:"email"`
	Phonenumber        string      `json:"phoneNumber"`
	Office             string      `json:"office"`
	Manager            string      `json:"manager"`
	Orgunit            string      `json:"orgUnit"`
	Maintext           string      `json:"mainText"`
	Github             string      `json:"gitHub"`
	Twitter            string      `json:"twitter"`
	Stackoverflow      interface{} `json:"stackOverflow"`
	Linkedin           string      `json:"linkedIn"`
	Imageportraiturl   string      `json:"imagePortraitUrl"`
	Imagewallofleeturl string      `json:"imageWallOfLeetUrl"`
	Highlighted        bool        `json:"highlighted"`
	Published          bool        `json:"published"`
}
