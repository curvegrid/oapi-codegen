// Package client provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package client

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// SchemaObject defines model for SchemaObject.
type SchemaObject struct {
	FirstName string `json:"firstName"`
	Role      string `json:"role"`
}

// Validate perform validation on the SchemaObject
func (s SchemaObject) Validate() error {
	// Run validate on a struct
	return validation.ValidateStruct(
		&s,
		validation.Field(
			&s.FirstName,
			validation.Required,
		),
		validation.Field(
			&s.Role,
			validation.Required,
		),
	)

}

// PostBothJSONBody defines parameters for PostBoth.
type PostBothJSONBody SchemaObject

// Validate perform validation on the PostBothJSONBody
func (s PostBothJSONBody) Validate() error {
	// Run validate on a scalar
	return validation.Validate(
		(SchemaObject)(s),
	)

}

// PostJsonJSONBody defines parameters for PostJson.
type PostJsonJSONBody SchemaObject

// Validate perform validation on the PostJsonJSONBody
func (s PostJsonJSONBody) Validate() error {
	// Run validate on a scalar
	return validation.Validate(
		(SchemaObject)(s),
	)

}

// PostBothRequestBody defines body for PostBoth for application/json ContentType.
type PostBothJSONRequestBody PostBothJSONBody

// PostJsonRequestBody defines body for PostJson for application/json ContentType.
type PostJsonJSONRequestBody PostJsonJSONBody

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
	// https://api.deepmap.com for example.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A callback for modifying requests which are generated before sending over
	// the network.
	RequestEditor RequestEditorFn
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
		client.Client = http.DefaultClient
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
		c.RequestEditor = fn
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// PostBoth request  with any body
	PostBothWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error)

	PostBoth(ctx context.Context, body PostBothJSONRequestBody) (*http.Response, error)

	// GetBoth request
	GetBoth(ctx context.Context) (*http.Response, error)

	// PostJson request  with any body
	PostJsonWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error)

	PostJson(ctx context.Context, body PostJsonJSONRequestBody) (*http.Response, error)

	// GetJson request
	GetJson(ctx context.Context) (*http.Response, error)

	// PostOther request  with any body
	PostOtherWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error)

	// GetOther request
	GetOther(ctx context.Context) (*http.Response, error)

	// GetJsonWithTrailingSlash request
	GetJsonWithTrailingSlash(ctx context.Context) (*http.Response, error)
}

func (c *Client) PostBothWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewPostBothRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) PostBoth(ctx context.Context, body PostBothJSONRequestBody) (*http.Response, error) {
	req, err := NewPostBothRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) GetBoth(ctx context.Context) (*http.Response, error) {
	req, err := NewGetBothRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) PostJsonWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewPostJsonRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) PostJson(ctx context.Context, body PostJsonJSONRequestBody) (*http.Response, error) {
	req, err := NewPostJsonRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) GetJson(ctx context.Context) (*http.Response, error) {
	req, err := NewGetJsonRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) PostOtherWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewPostOtherRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) GetOther(ctx context.Context) (*http.Response, error) {
	req, err := NewGetOtherRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) GetJsonWithTrailingSlash(ctx context.Context) (*http.Response, error) {
	req, err := NewGetJsonWithTrailingSlashRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

// NewPostBothRequest calls the generic PostBoth builder with application/json body
func NewPostBothRequest(server string, body PostBothJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostBothRequestWithBody(server, "application/json", bodyReader)
}

// NewPostBothRequestWithBody generates requests for PostBoth with any type of body
func NewPostBothRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/with_both_bodies")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// NewGetBothRequest generates requests for GetBoth
func NewGetBothRequest(server string) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/with_both_responses")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPostJsonRequest calls the generic PostJson builder with application/json body
func NewPostJsonRequest(server string, body PostJsonJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostJsonRequestWithBody(server, "application/json", bodyReader)
}

// NewPostJsonRequestWithBody generates requests for PostJson with any type of body
func NewPostJsonRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/with_json_body")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// NewGetJsonRequest generates requests for GetJson
func NewGetJsonRequest(server string) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/with_json_response")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPostOtherRequestWithBody generates requests for PostOther with any type of body
func NewPostOtherRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/with_other_body")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// NewGetOtherRequest generates requests for GetOther
func NewGetOtherRequest(server string) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/with_other_response")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetJsonWithTrailingSlashRequest generates requests for GetJsonWithTrailingSlash
func NewGetJsonWithTrailingSlashRequest(server string) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/with_trailing_slash/")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
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
	// PostBoth request  with any body
	PostBothWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*PostBothResponse, error)

	PostBothWithResponse(ctx context.Context, body PostBothJSONRequestBody) (*PostBothResponse, error)

	// GetBoth request
	GetBothWithResponse(ctx context.Context) (*GetBothResponse, error)

	// PostJson request  with any body
	PostJsonWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*PostJsonResponse, error)

	PostJsonWithResponse(ctx context.Context, body PostJsonJSONRequestBody) (*PostJsonResponse, error)

	// GetJson request
	GetJsonWithResponse(ctx context.Context) (*GetJsonResponse, error)

	// PostOther request  with any body
	PostOtherWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*PostOtherResponse, error)

	// GetOther request
	GetOtherWithResponse(ctx context.Context) (*GetOtherResponse, error)

	// GetJsonWithTrailingSlash request
	GetJsonWithTrailingSlashWithResponse(ctx context.Context) (*GetJsonWithTrailingSlashResponse, error)
}

type PostBothResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r PostBothResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostBothResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetBothResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r GetBothResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetBothResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostJsonResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r PostJsonResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostJsonResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetJsonResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r GetJsonResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetJsonResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostOtherResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r PostOtherResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostOtherResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetOtherResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r GetOtherResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetOtherResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetJsonWithTrailingSlashResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r GetJsonWithTrailingSlashResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetJsonWithTrailingSlashResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// PostBothWithBodyWithResponse request with arbitrary body returning *PostBothResponse
func (c *ClientWithResponses) PostBothWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*PostBothResponse, error) {
	rsp, err := c.PostBothWithBody(ctx, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParsePostBothResponse(rsp)
}

func (c *ClientWithResponses) PostBothWithResponse(ctx context.Context, body PostBothJSONRequestBody) (*PostBothResponse, error) {
	rsp, err := c.PostBoth(ctx, body)
	if err != nil {
		return nil, err
	}
	return ParsePostBothResponse(rsp)
}

// GetBothWithResponse request returning *GetBothResponse
func (c *ClientWithResponses) GetBothWithResponse(ctx context.Context) (*GetBothResponse, error) {
	rsp, err := c.GetBoth(ctx)
	if err != nil {
		return nil, err
	}
	return ParseGetBothResponse(rsp)
}

// PostJsonWithBodyWithResponse request with arbitrary body returning *PostJsonResponse
func (c *ClientWithResponses) PostJsonWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*PostJsonResponse, error) {
	rsp, err := c.PostJsonWithBody(ctx, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParsePostJsonResponse(rsp)
}

func (c *ClientWithResponses) PostJsonWithResponse(ctx context.Context, body PostJsonJSONRequestBody) (*PostJsonResponse, error) {
	rsp, err := c.PostJson(ctx, body)
	if err != nil {
		return nil, err
	}
	return ParsePostJsonResponse(rsp)
}

// GetJsonWithResponse request returning *GetJsonResponse
func (c *ClientWithResponses) GetJsonWithResponse(ctx context.Context) (*GetJsonResponse, error) {
	rsp, err := c.GetJson(ctx)
	if err != nil {
		return nil, err
	}
	return ParseGetJsonResponse(rsp)
}

// PostOtherWithBodyWithResponse request with arbitrary body returning *PostOtherResponse
func (c *ClientWithResponses) PostOtherWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*PostOtherResponse, error) {
	rsp, err := c.PostOtherWithBody(ctx, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParsePostOtherResponse(rsp)
}

// GetOtherWithResponse request returning *GetOtherResponse
func (c *ClientWithResponses) GetOtherWithResponse(ctx context.Context) (*GetOtherResponse, error) {
	rsp, err := c.GetOther(ctx)
	if err != nil {
		return nil, err
	}
	return ParseGetOtherResponse(rsp)
}

// GetJsonWithTrailingSlashWithResponse request returning *GetJsonWithTrailingSlashResponse
func (c *ClientWithResponses) GetJsonWithTrailingSlashWithResponse(ctx context.Context) (*GetJsonWithTrailingSlashResponse, error) {
	rsp, err := c.GetJsonWithTrailingSlash(ctx)
	if err != nil {
		return nil, err
	}
	return ParseGetJsonWithTrailingSlashResponse(rsp)
}

// ParsePostBothResponse parses an HTTP response from a PostBothWithResponse call
func ParsePostBothResponse(rsp *http.Response) (*PostBothResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &PostBothResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	}

	return response, nil
}

// ParseGetBothResponse parses an HTTP response from a GetBothWithResponse call
func ParseGetBothResponse(rsp *http.Response) (*GetBothResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetBothResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	}

	return response, nil
}

// ParsePostJsonResponse parses an HTTP response from a PostJsonWithResponse call
func ParsePostJsonResponse(rsp *http.Response) (*PostJsonResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &PostJsonResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	}

	return response, nil
}

// ParseGetJsonResponse parses an HTTP response from a GetJsonWithResponse call
func ParseGetJsonResponse(rsp *http.Response) (*GetJsonResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetJsonResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	}

	return response, nil
}

// ParsePostOtherResponse parses an HTTP response from a PostOtherWithResponse call
func ParsePostOtherResponse(rsp *http.Response) (*PostOtherResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &PostOtherResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	}

	return response, nil
}

// ParseGetOtherResponse parses an HTTP response from a GetOtherWithResponse call
func ParseGetOtherResponse(rsp *http.Response) (*GetOtherResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetOtherResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	}

	return response, nil
}

// ParseGetJsonWithTrailingSlashResponse parses an HTTP response from a GetJsonWithTrailingSlashWithResponse call
func ParseGetJsonWithTrailingSlashResponse(rsp *http.Response) (*GetJsonWithTrailingSlashResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetJsonWithTrailingSlashResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /with_both_bodies)
	PostBoth(ctx *PostBothContext) error

	// (GET /with_both_responses)
	GetBoth(ctx *GetBothContext) error

	// (POST /with_json_body)
	PostJson(ctx *PostJsonContext) error

	// (GET /with_json_response)
	GetJson(ctx *GetJsonContext) error

	// (POST /with_other_body)
	PostOther(ctx *PostOtherContext) error

	// (GET /with_other_response)
	GetOther(ctx *GetOtherContext) error

	// (GET /with_trailing_slash/)
	GetJsonWithTrailingSlash(ctx *GetJsonWithTrailingSlashContext) error
}

// PostBothContext is a context customized for PostBoth (POST /with_both_bodies).
type PostBothContext struct {
	echo.Context
}

// The body parsers
// ParseJSONBody tries to parse the body into the respective structure and validate it.
func (c *PostBothContext) ParseJSONBody() (PostBothJSONBody, error) {
	var resp PostBothJSONBody
	if err := c.Bind(&resp); err != nil {
		return resp, ValidationError{ParamType: "body", Err: errors.Wrap(err, "cannot parse as json")}
	}
	if err := resp.Validate(); err != nil {
		return resp, ValidationError{ParamType: "body", Err: err}
	}
	return resp, nil
}

// GetBothContext is a context customized for GetBoth (GET /with_both_responses).
type GetBothContext struct {
	echo.Context
}

// PostJsonContext is a context customized for PostJson (POST /with_json_body).
type PostJsonContext struct {
	echo.Context
}

// The body parsers
// ParseJSONBody tries to parse the body into the respective structure and validate it.
func (c *PostJsonContext) ParseJSONBody() (PostJsonJSONBody, error) {
	var resp PostJsonJSONBody
	if err := c.Bind(&resp); err != nil {
		return resp, ValidationError{ParamType: "body", Err: errors.Wrap(err, "cannot parse as json")}
	}
	if err := resp.Validate(); err != nil {
		return resp, ValidationError{ParamType: "body", Err: err}
	}
	return resp, nil
}

// GetJsonContext is a context customized for GetJson (GET /with_json_response).
type GetJsonContext struct {
	echo.Context
}

// PostOtherContext is a context customized for PostOther (POST /with_other_body).
type PostOtherContext struct {
	echo.Context
}

// The body parsers

// GetOtherContext is a context customized for GetOther (GET /with_other_response).
type GetOtherContext struct {
	echo.Context
}

// GetJsonWithTrailingSlashContext is a context customized for GetJsonWithTrailingSlash (GET /with_trailing_slash/).
type GetJsonWithTrailingSlashContext struct {
	echo.Context
}

// ValidationError is the special validation error type, returned from failed validation runs.
type ValidationError struct {
	ParamType string // can be "path", "cookie", "header", "query" or "body"
	Param     string // which field? can be omitted, when we parse the entire struct at once
	Err       error
}

// Error implements the error interface.
func (v ValidationError) Error() string {
	if v.Param == "" {
		return fmt.Sprintf("validation failed for '%s': %v", v.ParamType, v.Err)
	}
	return fmt.Sprintf("validation failed for %s parameter '%s': %v", v.ParamType, v.Param, v.Err)
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface

	middlewares []echo.MiddlewareFunc
}

// handlePostBoth converts echo context to params.
func (w *ServerInterfaceWrapper) handlePostBoth(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostBoth(&PostBothContext{ctx})
	return err
}

// PostBoth creates a handler function for the endpoint.
func (w *ServerInterfaceWrapper) PostBoth() echo.HandlerFunc {
	securityReqs := BindSecurityRequirements()
	// Wrap handler in middlewares
	handler := echo.HandlerFunc(w.handlePostBoth)
	for i := len(w.middlewares); i > 0; i-- {
		handler = w.middlewares[i-1](handler)
	}
	// Put securityReqs on top
	return securityReqs(handler)
}

// handleGetBoth converts echo context to params.
func (w *ServerInterfaceWrapper) handleGetBoth(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetBoth(&GetBothContext{ctx})
	return err
}

// GetBoth creates a handler function for the endpoint.
func (w *ServerInterfaceWrapper) GetBoth() echo.HandlerFunc {
	securityReqs := BindSecurityRequirements()
	// Wrap handler in middlewares
	handler := echo.HandlerFunc(w.handleGetBoth)
	for i := len(w.middlewares); i > 0; i-- {
		handler = w.middlewares[i-1](handler)
	}
	// Put securityReqs on top
	return securityReqs(handler)
}

// handlePostJson converts echo context to params.
func (w *ServerInterfaceWrapper) handlePostJson(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostJson(&PostJsonContext{ctx})
	return err
}

// PostJson creates a handler function for the endpoint.
func (w *ServerInterfaceWrapper) PostJson() echo.HandlerFunc {
	securityReqs := BindSecurityRequirements()
	// Wrap handler in middlewares
	handler := echo.HandlerFunc(w.handlePostJson)
	for i := len(w.middlewares); i > 0; i-- {
		handler = w.middlewares[i-1](handler)
	}
	// Put securityReqs on top
	return securityReqs(handler)
}

// handleGetJson converts echo context to params.
func (w *ServerInterfaceWrapper) handleGetJson(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetJson(&GetJsonContext{ctx})
	return err
}

// GetJson creates a handler function for the endpoint.
func (w *ServerInterfaceWrapper) GetJson() echo.HandlerFunc {
	securityReqs := BindSecurityRequirements(SecurityRequirement{SecuritySchemeOpenId, []string{"json.read", "json.admin"}})
	// Wrap handler in middlewares
	handler := echo.HandlerFunc(w.handleGetJson)
	for i := len(w.middlewares); i > 0; i-- {
		handler = w.middlewares[i-1](handler)
	}
	// Put securityReqs on top
	return securityReqs(handler)
}

// handlePostOther converts echo context to params.
func (w *ServerInterfaceWrapper) handlePostOther(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostOther(&PostOtherContext{ctx})
	return err
}

// PostOther creates a handler function for the endpoint.
func (w *ServerInterfaceWrapper) PostOther() echo.HandlerFunc {
	securityReqs := BindSecurityRequirements()
	// Wrap handler in middlewares
	handler := echo.HandlerFunc(w.handlePostOther)
	for i := len(w.middlewares); i > 0; i-- {
		handler = w.middlewares[i-1](handler)
	}
	// Put securityReqs on top
	return securityReqs(handler)
}

// handleGetOther converts echo context to params.
func (w *ServerInterfaceWrapper) handleGetOther(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetOther(&GetOtherContext{ctx})
	return err
}

// GetOther creates a handler function for the endpoint.
func (w *ServerInterfaceWrapper) GetOther() echo.HandlerFunc {
	securityReqs := BindSecurityRequirements()
	// Wrap handler in middlewares
	handler := echo.HandlerFunc(w.handleGetOther)
	for i := len(w.middlewares); i > 0; i-- {
		handler = w.middlewares[i-1](handler)
	}
	// Put securityReqs on top
	return securityReqs(handler)
}

// handleGetJsonWithTrailingSlash converts echo context to params.
func (w *ServerInterfaceWrapper) handleGetJsonWithTrailingSlash(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetJsonWithTrailingSlash(&GetJsonWithTrailingSlashContext{ctx})
	return err
}

// GetJsonWithTrailingSlash creates a handler function for the endpoint.
func (w *ServerInterfaceWrapper) GetJsonWithTrailingSlash() echo.HandlerFunc {
	securityReqs := BindSecurityRequirements(SecurityRequirement{SecuritySchemeOpenId, []string{"json.read", "json.admin"}})
	// Wrap handler in middlewares
	handler := echo.HandlerFunc(w.handleGetJsonWithTrailingSlash)
	for i := len(w.middlewares); i > 0; i-- {
		handler = w.middlewares[i-1](handler)
	}
	// Put securityReqs on top
	return securityReqs(handler)
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface, middlewares ...echo.MiddlewareFunc) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,

		middlewares: middlewares,
	}

	router.POST("/with_both_bodies", wrapper.PostBoth())
	router.GET("/with_both_responses", wrapper.GetBoth())
	router.POST("/with_json_body", wrapper.PostJson())
	router.GET("/with_json_response", wrapper.GetJson())
	router.POST("/with_other_body", wrapper.PostOther())
	router.GET("/with_other_response", wrapper.GetOther())
	router.GET("/with_trailing_slash/", wrapper.GetJsonWithTrailingSlash())

}

// SecurityScheme represents a security scheme used in the server.
type SecurityScheme string

// ScopesKey returns the key of the scopes in the Context.
func (ss SecurityScheme) ScopesKey() string {
	return string(ss) + ".Scopes"
}

// Scopes collect the scopes defined in the Context.
func (ss SecurityScheme) Scopes(c echo.Context) ([]string, bool) {
	val := c.Get(ss.ScopesKey())
	scopes, ok := val.([]string)
	return scopes, ok
}

// SecurityRequirement is a requirement of an endpoint on the allowed scopes a scheme can be used.
type SecurityRequirement struct {
	Scheme SecurityScheme
	Scopes []string
}

// BindSecurityRequirements returns an echo middleware that sets the scopes of the security schemes.
func BindSecurityRequirements(reqs ...SecurityRequirement) echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			for _, req := range reqs {
				ctx.Set(req.Scheme.ScopesKey(), req.Scopes)
			}
			return h(ctx)
		}
	}
}

// All security schemes defined.
const ()

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8yUz24TMRDGX2U1cFyyKdz2CAdUJAgikTiEqHK8k9jVrm1mJq1W0b47GicliahCkGjV",
	"SzTO/NE338/rLdjYpRgwCEO9BbYOO5PDaQ4ny1u0oudEMSGJx5xdeWL5YjrUg/QJoQYW8mENQwkU28cS",
	"msGfG0/YQD3fVZVHoxaDlviwitrcIFvySXwMUMPMeS4EWbi4dygOqRCHxYfWY5DChGYffvfiviGnGBi5",
	"MITFGgOSEWwKG4nQStv/CFBC6y0GzjpDXgQ+X89UvXhR+TBDlmKKdIcEJdwh8U7K1Wg8GmthTBhM8lDD",
	"u9F4dAUlJCMu+1Pde3E3y5h/mr1pKXK2Uo00utd1AzV8jSzvozjYuYN6anqtszEIhtxiUmq9zU3VLauM",
	"B1gavSZcQQ2vqgPNao+yOuGo/h6PilZQ3rAQmu505CpSZwRqWPpgqIfyD5gnNIU2mP/YOw912LSt1hw5",
	"cZTdwhof8eIjHqw4qn07Hr9UE4bDjipJaffnWX9S5c/C+p8IZfUP2XOAfut/QkAqi9FuyEsP9XwLk4RZ",
	"wBx07ojQNFDuYtN0PsBiWBx2ifo+XIBionUXs3i2j2Un/xIWhwXOw/hfV1zI+NaH9Q23hl31t2uij/Fs",
	"3zLVjhd6b4bhVwAAAP//2pHiCAkHAAA=",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
