package circlecigo

import (
	"github.com/dghubble/sling"
	"net/http"
	"time"
)

type Client struct {
	sling                *sling.Sling
	Projects             *ProjectService
	SshKeys              *SshKeyService
	EnvironmentVariables *EnvironmentVariableService
}

func NewClient(authToken, baseURL string) *Client {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	base := sling.New().Client(client).Base(baseURL).Set("Content-Type", "application/json")

	return &Client{
		sling:                base,
		Projects:             newProjectService(base, authToken),
		SshKeys:              newSshKeyService(base, authToken),
		EnvironmentVariables: newEnvironmentVariableService(base, authToken),
	}
}
