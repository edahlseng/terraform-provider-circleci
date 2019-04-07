package circlecigo

import (
	"fmt"
	"net/http"
)

func relevantErrorFromStatusCode(resp *http.Response, httpError error) error {
	if httpError != nil {
		return httpError
	}

	if 200 <= resp.StatusCode && resp.StatusCode <= 299 {
		return nil
	}

	if resp.StatusCode == 404 {
		return fmt.Errorf("Not found")
	}

	return fmt.Errorf("Unknown API error encountered. Status Code: %d", resp.StatusCode)
}
