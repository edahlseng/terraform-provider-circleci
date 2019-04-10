package circlecigo

import (
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)

type SshKeyService struct {
	sling *sling.Sling
	token string
}

func newSshKeyService(sling *sling.Sling, token string) *SshKeyService {
	return &SshKeyService{
		sling: sling.New(),
		token: token,
	}
}

// -----------------------------------------------------------------------------
// Input
// -----------------------------------------------------------------------------

type SshKey struct {
	ProjectId      string
	Hostname       string
	PrivateKey     string
	FingerprintMd5 string
}

type sshKeyJsonInputCreate struct {
	Hostname   string `json:"hostname"`
	PrivateKey string `json:"private_key"`
}

type sshKeyJsonInputDelete struct {
	Hostname    string `json:"hostname"`
	Fingerprint string `json:"fingerprint"`
}

// -----------------------------------------------------------------------------
// CRUD
// -----------------------------------------------------------------------------

func (service *SshKeyService) Create(sshKey *SshKey) (*SshKey, *http.Response, error) {
	if sshKey == nil {
		return nil, nil, fmt.Errorf("Cannot create an SSH key that is nil")
	}

	requestBody := &sshKeyJsonInputCreate{Hostname: sshKey.Hostname, PrivateKey: sshKey.PrivateKey}

	req, reqErr := service.sling.New().Post("project/" + sshKey.ProjectId + "/ssh-key?circle-token=" + service.token).BodyJSON(requestBody).Request()
	if reqErr != nil {
		return nil, nil, reqErr
	}
	resp, httpErr := http.DefaultClient.Do(req)
	finalErr := relevantErrorFromStatusCode(resp, httpErr)
	var responseSshKey *SshKey = nil
	if finalErr == nil {
		responseSshKey = sshKey
	}

	return responseSshKey, resp, finalErr
}

func (service *SshKeyService) Delete(sshKey *SshKey) (*http.Response, error) {
	if sshKey == nil {
		return nil, fmt.Errorf("Cannot delete an SSH key that is nil")
	}

	requestBody := &sshKeyJsonInputDelete{Hostname: sshKey.Hostname, Fingerprint: sshKey.FingerprintMd5}

	req, reqErr := service.sling.New().Delete("project/" + sshKey.ProjectId + "/ssh-key?circle-token=" + service.token).BodyJSON(requestBody).Request()
	if reqErr != nil {
		return nil, reqErr
	}
	resp, httpErr := http.DefaultClient.Do(req)

	return resp, relevantErrorFromStatusCode(resp, httpErr)
}
