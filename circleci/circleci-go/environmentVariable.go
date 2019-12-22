package circlecigo

import (
	"encoding/json"
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)

type EnvironmentVariableService struct {
	sling *sling.Sling
	token string
}

func newEnvironmentVariableService(sling *sling.Sling, token string) *EnvironmentVariableService {
	return &EnvironmentVariableService{
		sling: sling.New(),
		token: token,
	}
}

// -----------------------------------------------------------------------------
// Input
// -----------------------------------------------------------------------------

type EnvironmentVariable struct {
	ProjectId   string
	Name        string
	Value       string
	ValueMasked string
}

type environmentVariableInputCreate struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type environmentVariableResponseType struct {
	Name        string `json:"name"`
	ValueMasked string `json:"value"`
}

// -----------------------------------------------------------------------------
// CRUD
// -----------------------------------------------------------------------------

func (service *EnvironmentVariableService) Create(environmentVariable *EnvironmentVariable) (*EnvironmentVariable, *http.Response, error) {
	if environmentVariable == nil {
		return nil, nil, fmt.Errorf("Cannot create an environment variable that is nil")
	}

	requestBody := &environmentVariableInputCreate{Name: environmentVariable.Name, Value: environmentVariable.Value}

	environmentVariableResponse := new(environmentVariableResponseType)
	genericError := new(json.RawMessage)
	resp, err := service.sling.New().Set("Accept", "application/json").Post("project/"+environmentVariable.ProjectId+"/envvar?circle-token="+service.token).BodyJSON(requestBody).Receive(environmentVariableResponse, genericError)
	if err != nil {
		return nil, nil, err
	}

	finalErr := relevantErrorFromStatusCode(resp, err)
	var responseEnvironmentVariable *EnvironmentVariable = nil
	if finalErr == nil {
		responseEnvironmentVariable = &EnvironmentVariable{
			ProjectId:   environmentVariable.ProjectId,
			Name:        environmentVariableResponse.Name,
			Value:       environmentVariable.Value,
			ValueMasked: environmentVariableResponse.ValueMasked,
		}
	}

	return responseEnvironmentVariable, resp, finalErr
}

func (service *EnvironmentVariableService) Read(environmentVariable *EnvironmentVariable) (*EnvironmentVariable, *http.Response, error) {
	if environmentVariable == nil {
		return nil, nil, fmt.Errorf("Cannot read an environment variable that is nil")
	}

	environmentVariableResponse := new(environmentVariableResponseType)
	genericError := new(json.RawMessage)
	resp, err := service.sling.New().Set("Accept", "application/json").Get("project/"+environmentVariable.ProjectId+"/envvar/"+environmentVariable.Name+"?circle-token="+service.token).Receive(environmentVariableResponse, genericError)
	if err != nil {
		return nil, nil, err
	}

	finalErr := relevantErrorFromStatusCode(resp, err)
	var responseEnvironmentVariable *EnvironmentVariable = nil
	if finalErr == nil {
		responseEnvironmentVariable = &EnvironmentVariable{
			ProjectId:   environmentVariable.ProjectId,
			Name:        environmentVariableResponse.Name,
			Value:       environmentVariable.Value,
			ValueMasked: environmentVariableResponse.ValueMasked,
		}
	}

	return responseEnvironmentVariable, resp, finalErr
}

func (service *EnvironmentVariableService) Delete(environmentVariable *EnvironmentVariable) (*http.Response, error) {
	if environmentVariable == nil {
		return nil, fmt.Errorf("Cannot delete an environment variable that is nil")
	}

	req, reqErr := service.sling.New().Delete("project/" + environmentVariable.ProjectId + "/envvar/" + environmentVariable.Name + "?circle-token=" + service.token).Request()
	if reqErr != nil {
		return nil, reqErr
	}
	resp, httpErr := http.DefaultClient.Do(req)

	return resp, relevantErrorFromStatusCode(resp, httpErr)
}
