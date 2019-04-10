package circlecigo

import (
	"fmt"
	"io/ioutil"
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

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return fmt.Errorf("Unknown API error encountered. Status Code: %d. Unable to read response body: %s.", resp.StatusCode, readErr)
	}
	closeErr := resp.Body.Close()
	if closeErr != nil {
		return fmt.Errorf("Unknown API error encountered. Status Code: %d. Unable to close response body: %s.", resp.StatusCode, closeErr)
	}

	return fmt.Errorf("Unknown API error encountered. Status Code: %d. Response Body: %s", resp.StatusCode, string(body))
}
