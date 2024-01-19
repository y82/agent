// Package dataplane provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package dataplane

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	externalRef0 "github.com/nginx/agent/v3/api/http/common"
	"github.com/oapi-codegen/runtime"
)

// Defines values for ConfigurationStatusType.
const (
	FAILED         ConfigurationStatusType = "FAILED"
	INPROGESS      ConfigurationStatusType = "IN_PROGESS"
	ROLLBACKFAILED ConfigurationStatusType = "ROLLBACK_FAILED"
	SUCCESS        ConfigurationStatusType = "SUCCESS"
)

// Configuration defines model for Configuration.
type Configuration struct {
	Location *string `json:"location,omitempty"`
}

// ConfigurationStatus defines model for ConfigurationStatus.
type ConfigurationStatus struct {
	CorrelationId *string    `json:"correlationId,omitempty"`
	LastUpdated   *time.Time `json:"lastUpdated,omitempty"`
	Message       *string    `json:"message,omitempty"`

	// Status The type of configuration status
	Status *ConfigurationStatusType `json:"status,omitempty"`
}

// ConfigurationStatusType The type of configuration status
type ConfigurationStatusType string

// CorrelationId defines model for CorrelationId.
type CorrelationId struct {
	CorrelationId *string `json:"correlationId,omitempty"`
}

// UpdateInstanceConfigurationJSONRequestBody defines body for UpdateInstanceConfiguration for application/json ContentType.
type UpdateInstanceConfigurationJSONRequestBody = Configuration

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetInstances request
	GetInstances(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateInstanceConfigurationWithBody request with any body
	UpdateInstanceConfigurationWithBody(ctx context.Context, instanceId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateInstanceConfiguration(ctx context.Context, instanceId string, body UpdateInstanceConfigurationJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetInstanceConfigurationStatus request
	GetInstanceConfigurationStatus(ctx context.Context, instanceId string, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetInstances(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetInstancesRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateInstanceConfigurationWithBody(ctx context.Context, instanceId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateInstanceConfigurationRequestWithBody(c.Server, instanceId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateInstanceConfiguration(ctx context.Context, instanceId string, body UpdateInstanceConfigurationJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateInstanceConfigurationRequest(c.Server, instanceId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetInstanceConfigurationStatus(ctx context.Context, instanceId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetInstanceConfigurationStatusRequest(c.Server, instanceId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetInstancesRequest generates requests for GetInstances
func NewGetInstancesRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/instances")
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

// NewUpdateInstanceConfigurationRequest calls the generic UpdateInstanceConfiguration builder with application/json body
func NewUpdateInstanceConfigurationRequest(server string, instanceId string, body UpdateInstanceConfigurationJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateInstanceConfigurationRequestWithBody(server, instanceId, "application/json", bodyReader)
}

// NewUpdateInstanceConfigurationRequestWithBody generates requests for UpdateInstanceConfiguration with any type of body
func NewUpdateInstanceConfigurationRequestWithBody(server string, instanceId string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "instanceId", runtime.ParamLocationPath, instanceId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/instances/%s/configurations", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewGetInstanceConfigurationStatusRequest generates requests for GetInstanceConfigurationStatus
func NewGetInstanceConfigurationStatusRequest(server string, instanceId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "instanceId", runtime.ParamLocationPath, instanceId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/instances/%s/configurations/status", pathParam0)
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

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetInstancesWithResponse request
	GetInstancesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetInstancesResponse, error)

	// UpdateInstanceConfigurationWithBodyWithResponse request with any body
	UpdateInstanceConfigurationWithBodyWithResponse(ctx context.Context, instanceId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateInstanceConfigurationResponse, error)

	UpdateInstanceConfigurationWithResponse(ctx context.Context, instanceId string, body UpdateInstanceConfigurationJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateInstanceConfigurationResponse, error)

	// GetInstanceConfigurationStatusWithResponse request
	GetInstanceConfigurationStatusWithResponse(ctx context.Context, instanceId string, reqEditors ...RequestEditorFn) (*GetInstanceConfigurationStatusResponse, error)
}

type GetInstancesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]externalRef0.Instance
	JSON500      *externalRef0.InternalServerError
}

// Status returns HTTPResponse.Status
func (r GetInstancesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetInstancesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateInstanceConfigurationResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *CorrelationId
	JSON404      *externalRef0.NotFound
	JSON500      *externalRef0.InternalServerError
}

// Status returns HTTPResponse.Status
func (r UpdateInstanceConfigurationResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateInstanceConfigurationResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetInstanceConfigurationStatusResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ConfigurationStatus
	JSON404      *externalRef0.NotFound
	JSON500      *externalRef0.InternalServerError
}

// Status returns HTTPResponse.Status
func (r GetInstanceConfigurationStatusResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetInstanceConfigurationStatusResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetInstancesWithResponse request returning *GetInstancesResponse
func (c *ClientWithResponses) GetInstancesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetInstancesResponse, error) {
	rsp, err := c.GetInstances(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetInstancesResponse(rsp)
}

// UpdateInstanceConfigurationWithBodyWithResponse request with arbitrary body returning *UpdateInstanceConfigurationResponse
func (c *ClientWithResponses) UpdateInstanceConfigurationWithBodyWithResponse(ctx context.Context, instanceId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateInstanceConfigurationResponse, error) {
	rsp, err := c.UpdateInstanceConfigurationWithBody(ctx, instanceId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateInstanceConfigurationResponse(rsp)
}

func (c *ClientWithResponses) UpdateInstanceConfigurationWithResponse(ctx context.Context, instanceId string, body UpdateInstanceConfigurationJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateInstanceConfigurationResponse, error) {
	rsp, err := c.UpdateInstanceConfiguration(ctx, instanceId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateInstanceConfigurationResponse(rsp)
}

// GetInstanceConfigurationStatusWithResponse request returning *GetInstanceConfigurationStatusResponse
func (c *ClientWithResponses) GetInstanceConfigurationStatusWithResponse(ctx context.Context, instanceId string, reqEditors ...RequestEditorFn) (*GetInstanceConfigurationStatusResponse, error) {
	rsp, err := c.GetInstanceConfigurationStatus(ctx, instanceId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetInstanceConfigurationStatusResponse(rsp)
}

// ParseGetInstancesResponse parses an HTTP response from a GetInstancesWithResponse call
func ParseGetInstancesResponse(rsp *http.Response) (*GetInstancesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetInstancesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []externalRef0.Instance
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest externalRef0.InternalServerError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseUpdateInstanceConfigurationResponse parses an HTTP response from a UpdateInstanceConfigurationWithResponse call
func ParseUpdateInstanceConfigurationResponse(rsp *http.Response) (*UpdateInstanceConfigurationResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateInstanceConfigurationResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest CorrelationId
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest externalRef0.NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest externalRef0.InternalServerError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetInstanceConfigurationStatusResponse parses an HTTP response from a GetInstanceConfigurationStatusWithResponse call
func ParseGetInstanceConfigurationStatusResponse(rsp *http.Response) (*GetInstanceConfigurationStatusResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetInstanceConfigurationStatusResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ConfigurationStatus
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest externalRef0.NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest externalRef0.InternalServerError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}
