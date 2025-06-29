package gitlab

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type GetApiV4WebCommitsPublicKeyResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

func (c *Client) GetApiV4WebCommitsPublicKey(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetApiV4WebCommitsPublicKeyRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
func NewGetApiV4WebCommitsPublicKeyRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/api/v4/web_commits/public_key")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
func (r GetApiV4WebCommitsPublicKeyResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}
func (r GetApiV4WebCommitsPublicKeyResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}
func (c *ClientWithResponses) GetApiV4WebCommitsPublicKeyWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetApiV4WebCommitsPublicKeyResponse, error) {
	rsp, err := c.GetApiV4WebCommitsPublicKey(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetApiV4WebCommitsPublicKeyResponse(rsp)
}
func ParseGetApiV4WebCommitsPublicKeyResponse(rsp *http.Response) (*GetApiV4WebCommitsPublicKeyResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetApiV4WebCommitsPublicKeyResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
