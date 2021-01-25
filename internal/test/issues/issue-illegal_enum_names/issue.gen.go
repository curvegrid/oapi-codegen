// Package illegal_enum_names provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package illegal_enum_names

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
)

// Bar defines model for Bar.
type Bar string

// List of Bar
const (
	Bar_Bar      Bar = "Bar"
	Bar_Foo      Bar = "Foo"
	Bar_Foo_Bar  Bar = "Foo Bar"
	Bar_Foo_Bar1 Bar = "Foo-Bar"
	Bar__Foo     Bar = "1Foo"
	Bar__Foo1    Bar = " Foo"
	Bar__Foo_    Bar = " Foo "
	Bar__Foo_1   Bar = "_Foo_"
)

// Validate perform validation on the Bar
func (s Bar) Validate() error {
	// Run validate on a scalar
	return validation.Validate(
		(string)(s),
		validation.In(
			" Foo", " Foo ", "1Foo", "Bar", "Foo", "Foo Bar", "Foo-Bar", "_Foo_",
		),
	)

}

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
	// GetFoo request
	GetFoo(ctx context.Context) (*http.Response, error)
}

func (c *Client) GetFoo(ctx context.Context) (*http.Response, error) {
	req, err := NewGetFooRequest(c.Server)
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

// NewGetFooRequest generates requests for GetFoo
func NewGetFooRequest(server string) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/foo")
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
	// GetFoo request
	GetFooWithResponse(ctx context.Context) (*GetFooResponse, error)
}

type GetFooResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]Bar
}

// Status returns HTTPResponse.Status
func (r GetFooResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetFooResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetFooWithResponse request returning *GetFooResponse
func (c *ClientWithResponses) GetFooWithResponse(ctx context.Context) (*GetFooResponse, error) {
	rsp, err := c.GetFoo(ctx)
	if err != nil {
		return nil, err
	}
	return ParseGetFooResponse(rsp)
}

// ParseGetFooResponse parses an HTTP response from a GetFooWithResponse call
func ParseGetFooResponse(rsp *http.Response) (*GetFooResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetFooResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []Bar
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /foo)
	GetFoo(ctx *GetFooContext) error
}

// GetFooContext is a context customized for GetFoo (GET /foo).
type GetFooContext struct {
	echo.Context
}

// Responses

// OK responses with the appropriate code and the JSON response.
func (c *GetFooContext) OK(resp GetFooResponseOK) error {
	return c.JSON(200, resp)
}

// GetFooResponseOK is the response type for GetFoo's "200" response.
type GetFooResponseOK = []Bar

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

	securityHandler SecurityHandler
}

type (
	// SecurityScheme is a security scheme name
	SecurityScheme string

	// SecurityScopes is a list of security scopes
	SecurityScopes []string

	// SecurityReq is a map of security scheme names and their respective scopes
	SecurityReq map[SecurityScheme]SecurityScopes

	// SecurityHandler defines a function to handle the security requirements
	// defined in the OpenAPI specification.
	SecurityHandler func(echo.Context, SecurityReq) error
)

// GetFoo converts echo context to params.
func (w *ServerInterfaceWrapper) GetFoo(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetFoo(&GetFooContext{ctx})
	return err
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
func RegisterHandlers(router EchoRouter, si ServerInterface, sh SecurityHandler, m ...echo.MiddlewareFunc) {
	RegisterHandlersWithBaseURL(router, si, "", sh, m...)
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string, sh SecurityHandler, m ...echo.MiddlewareFunc) {

	wrapper := ServerInterfaceWrapper{
		Handler:         si,
		securityHandler: sh,
	}

	router.GET(baseURL+"/foo", wrapper.GetFoo, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/0yQzU4DMQyEX6UaOIbsUm45cijiGRCqoq23DeraUWKQqirvjpyFQi4z+ZnYn6+YZMnC",
	"xFoRrqjTiZbY7XMsJsSfC8IbdiJw/dCZ39zcw+oe/z/YrBuTDRz2O5E93h30kgkBVUviI1prDolnsTqa",
	"9Gx33ns4fFGpSRgBox/9iOYgmTjmhIAnP/otHHLUU291mKX/cSQ1kUwlahJ+PSDghXTtplDNwpV6ZDuO",
	"JpOwEvdUzPmcpp4bPqrV/h2HuaS09OB9oRkBd8Pf4IafqQ0G326UsZR4WSEPVKeSsq5Ihtj6+g4AAP//",
	"Jkt64H8BAAA=",
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
