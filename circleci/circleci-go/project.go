package circlecigo

import (
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
	"strings"
)

type ProjectService struct {
	sling *sling.Sling
	token string
}

func newProjectService(sling *sling.Sling, token string) *ProjectService {
	return &ProjectService{
		sling: sling.New().Path("project/"),
		token: token,
	}
}

// -----------------------------------------------------------------------------
// Input
// -----------------------------------------------------------------------------

type Project struct {
	VcsType  string
	Username string
	Name     string
}

// -----------------------------------------------------------------------------
// Response
// -----------------------------------------------------------------------------

type ProjectResponse struct{}

type ProjectResponseError struct{}

// -----------------------------------------------------------------------------
// Helpers
// -----------------------------------------------------------------------------

func projectFromId(projectId string) (*Project, error) {
	parts := strings.Split(projectId, "/")
	if len(parts) < 3 || 3 < len(parts) {
		return nil, fmt.Errorf("Invalid project ID")
	}
	return &Project{VcsType: parts[0], Username: parts[1], Name: parts[2]}, nil
}

func ProjectIdFromProject(project Project) string {
	return fmt.Sprintf("%s/%s/%s", project.VcsType, project.Username, project.Name)
}

// -----------------------------------------------------------------------------
// CRUD
// -----------------------------------------------------------------------------

func (service *ProjectService) Create(project *Project) (*Project, *http.Response, error) {
	if project == nil {
		return nil, nil, fmt.Errorf("Cannot create a project that is nil")
	}

	req, reqErr := service.sling.New().Post(ProjectIdFromProject(*project) + "/follow?circle-token=" + service.token).Request()
	if reqErr != nil {
		return nil, nil, reqErr
	}
	resp, httpErr := http.DefaultClient.Do(req)
	finalErr := relevantErrorFromStatusCode(resp, httpErr)
	var responseProject *Project = nil
	if finalErr == nil {
		responseProject = project
	}

	return responseProject, resp, finalErr
}

func (service *ProjectService) Read(projectId string) (*Project, *http.Response, error) {
	project, err := projectFromId(projectId)
	if err != nil {
		return nil, nil, err
	}

	req, reqErr := service.sling.New().Head(projectId + "?circle-token=" + service.token).Request()
	if reqErr != nil {
		return nil, nil, reqErr
	}
	resp, httpErr := http.DefaultClient.Do(req) // Need to make request ourselves, as Sling doesn't seem to handle the empty body from a Head request very well
	finalErr := relevantErrorFromStatusCode(resp, httpErr)
	var responseProject *Project = nil
	if finalErr == nil {
		responseProject = project
	}

	return responseProject, resp, finalErr
}
