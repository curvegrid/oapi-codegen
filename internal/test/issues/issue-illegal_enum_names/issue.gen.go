// Package illegal_enum_names provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version (devel) DO NOT EDIT.
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
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
)

// Defines values for Bar.
const (
	BarBar Bar = "Bar"

	BarFoo Bar = "Foo"

	BarFoo1 Bar = "1Foo"

	BarFoo2 Bar = " Foo"

	BarFoo3 Bar = " Foo "

	BarFoo4 Bar = "_Foo_"

	BarFooBar Bar = "Foo Bar"

	BarFooBar1 Bar = "Foo-Bar"
)

// Bar defines model for Bar.
type Bar string

// Validate perform validation on the Bar
func (s Bar) Validate() error {
	// Run validate on an enum
	if err := validation.Validate(
		s,
		validation.In(
			BarBar, BarFoo, BarFoo1, BarFoo2, BarFoo3, BarFoo4, BarFooBar, BarFooBar1,
		),
		validation.Skip, // do not recurse infinitely
	); err != nil {
		return err
	}
	// Run validate on a scalar
	return validation.Validate(
		(string)(s),
	)

}

// GetFooResponseOK defines parameters for GetFoo.
type GetFooResponseOK []Bar

// Validate perform validation on the GetFooResponseOK
func (s GetFooResponseOK) Validate() error {
	// Run validate on a scalar
	return validation.Validate(
		([]Bar)(s),
		validation.Each(),
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
	// GetFoo request
	GetFoo(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetFoo(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetFooRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetFooRequest generates requests for GetFoo
func NewGetFooRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/foo")
	if operationPath[0] == '/' {
		operationPath = operationPath[1:]
	}
	operationURL := url.URL{
		Path: operationPath,
	}

	queryURL := serverURL.ResolveReference(&operationURL)

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
	// GetFoo request
	GetFooWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetFooResponse, error)
}

// GetFooResponseJSON200 represents a possible response for the GetFoo request.
type GetFooResponseJSON200 []Bar
type GetFooResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetFooResponseJSON200
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
func (c *ClientWithResponses) GetFooWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetFooResponse, error) {
	rsp, err := c.GetFoo(ctx, reqEditors...)
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
		var dest GetFooResponseJSON200
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
	return c.Context.JSON(200, resp)
}

// bindValidateBody decodes and validates the body of a request. It's highly inspired
// from the echo.DefaultBinder BindBody function.
// This is preferred over echo.Bind, since it grants more control over the binding
// functionality. Particularly, it returns a well-formatted ValidationError on invalid input.
func bindValidateBody(c echo.Context, i validation.Validatable) error {
	req := c.Request()
	if req.ContentLength != 0 {
		// Decode
		ctype := req.Header.Get(echo.HeaderContentType)
		switch {
		case strings.HasPrefix(ctype, echo.MIMEApplicationJSON):
			if err := json.NewDecoder(req.Body).Decode(i); err != nil {
				// Add some context to the error when possible
				switch e := err.(type) {
				case *json.UnmarshalTypeError:
					err = fmt.Errorf("cannot unmarshal a value of type %v into the field %v of type %v (offset %v)", e.Value, e.Field, e.Type, e.Offset)
				case *json.SyntaxError:
					err = fmt.Errorf("%v (offset %v)", err.Error(), e.Offset)
				}
				return &ValidationError{ParamType: "body", Err: err}
			}
		default:
			return echo.ErrUnsupportedMediaType
		}
	}

	// Validate
	if err := i.Validate(); err != nil {
		return &ValidationError{ParamType: "body", Err: err}
	}
	return nil
}

// ValidationError is the special validation error type, returned from failed validation runs.
type ValidationError struct {
	ParamType string // can be "path", "cookie", "header", "query" or "body"
	Param     string // which field? can be omitted, when we parse the entire struct at once
	Err       error
}

// Error implements the error interface.
func (v *ValidationError) Error() string {
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

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
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

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
