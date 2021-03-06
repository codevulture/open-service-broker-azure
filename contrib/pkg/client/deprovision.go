package client

import (
	"fmt"
	"net/http"
)

// Deprovision initiates deprovisioning of a service instance
func Deprovision(
	host string,
	port int,
	username string,
	password string,
	instanceID string,
) error {
	url := fmt.Sprintf(
		"%s/v2/service_instances/%s",
		getBaseURL(host, port),
		instanceID,
	)
	req, err := http.NewRequest(
		http.MethodDelete,
		url,
		nil,
	)
	if err != nil {
		return fmt.Errorf("error building request: %s", err)
	}
	if username != "" || password != "" {
		addAuthHeader(req, username, password)
	}
	q := req.URL.Query()
	q.Add("accepts_incomplete", "true")
	req.URL.RawQuery = q.Encode()
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error executing deprovision call: %s", err)
	}
	defer resp.Body.Close() // nolint: errcheck
	if resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf(
			"unanticipated http response code %d",
			resp.StatusCode,
		)
	}
	return nil
}
