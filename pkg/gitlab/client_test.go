package gitlab

import (
	"context"
	"net/http"
	"testing"
)

func bearerAuth(ctx context.Context, req *http.Request) error {
	token := "<TOKEN>"
	req.Header.Set("Authorization", "Bearer "+token)
	return nil
}

func client() (*ClientWithResponses, error) {
	hc := http.Client{}

	return NewClientWithResponses("http://127.0.0.1", WithHTTPClient(&hc))
}

func newClient(t *testing.T) *ClientWithResponses {
	t.Skip()

	c, err := client()
	if err != nil {
		t.Errorf("%v", err)
	}

	return c
}

func assertHTTPStatus(t *testing.T, resp *http.Response, err error, msg any) {
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if resp.StatusCode < http.StatusOK || http.StatusMultipleChoices <= resp.StatusCode {
		t.Errorf("HTTP Status: %v %v", resp.StatusCode, msg)
	}
}

func assertResponseBody(t *testing.T, resp any) {
	if resp == nil {
		t.Errorf("Error: No Content")
	}
}

func TestMain(m *testing.M) {
	m.Run()
}

func TestGetApiV4ProjectsWithResponse(t *testing.T) {
	c := newClient(t)

	params := GetApiV4ProjectsParams{}
	resp, err := c.GetApiV4ProjectsWithResponse(context.TODO(), &params, bearerAuth)

	assertHTTPStatus(t, resp.HTTPResponse, err, string(resp.Body))
	assertResponseBody(t, &resp.JSON200)
}
