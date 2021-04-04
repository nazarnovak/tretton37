package fetcher

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/nazarnovak/tretton37/backend/pkg/employee"
)

const (
	// These could be moved to some kind of a config, like a JSON file for example, that will make it easier to keep track of
	// these values and change them in case we need to
	apiEmployeesURL     = "https://api.1337co.de/v3/employees"
	headerAuthorization = "api-key 14:2021-03-31:tina.nielsen@tretton37.com 5aa6b2b66fbda96e398938b422906ce486e245de34f1212e64c9899006e00523"
)

type Tretton37 struct{}

func (t Tretton37) GetAll() ([]employee.Employee, error) {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, apiEmployeesURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", headerAuthorization)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var employees []employee.Employee
	if err := json.Unmarshal(body, &employees); err != nil {
		return nil, err
	}

	return employees, nil
}
