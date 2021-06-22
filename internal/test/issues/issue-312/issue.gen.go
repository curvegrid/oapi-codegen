// Package issue_312 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version (devel) DO NOT EDIT.
package issue_312

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// Error defines model for Error.
type Error struct {

	// Error code
	Code int32 `json:"code"`

	// Error message
	Message string `json:"message"`
}

// Validate perform validation on the Error
func (s Error) Validate() error {
	// Run validate on a struct
	return validation.ValidateStruct(
		&s,
		validation.Field(
			&s.Code,
			validation.Required,
		),
		validation.Field(
			&s.Message,
			validation.Required,
		),
	)

}

// Pet defines model for Pet.
type Pet struct {

	// The name of the pet.
	Name string `json:"name"`
}

// Validate perform validation on the Pet
func (s Pet) Validate() error {
	// Run validate on a struct
	return validation.ValidateStruct(
		&s,
		validation.Field(
			&s.Name,
			validation.Required,
		),
	)

}

// PetNames defines model for PetNames.
type PetNames struct {

	// The names of the pets.
	Names []string `json:"names"`
}

// Validate perform validation on the PetNames
func (s PetNames) Validate() error {
	// Run validate on a struct
	return validation.ValidateStruct(
		&s,
		validation.Field(
			&s.Names,
			validation.Required,
			eachWithIndirection(),
		),
	)

}

// validation.Each does not handle a pointer to slices/arrays or maps.
// This does the job.
func eachWithIndirection(rules ...validation.Rule) validation.Rule {
	return validation.By(func(value interface{}) error {
		v, isNil := validation.Indirect(value)
		if isNil {
			return nil
		}
		return validation.Each(rules...).Validate(v)
	})
}

// GetPetPathPetId defines parameters for GetPet.
type GetPetPathPetId string

// Validate perform validation on the GetPetPathPetId
func (s GetPetPathPetId) Validate() error {
	// Run validate on a scalar
	return validation.Validate(
		(string)(s),
	)

}

// GetPetResponseOK defines parameters for GetPet.
type GetPetResponseOK = Pet

// ValidatePetsJSONBody defines parameters for ValidatePets.
type ValidatePetsJSONBody PetNames

// Validate perform validation on the ValidatePetsJSONBody
func (s ValidatePetsJSONBody) Validate() error {
	// Run validate on a scalar
	return validation.Validate(
		(PetNames)(s),
	)

}

// ValidatePetsResponseOK defines parameters for ValidatePets.
type ValidatePetsResponseOK []Pet

// Validate perform validation on the ValidatePetsResponseOK
func (s ValidatePetsResponseOK) Validate() error {
	// Run validate on a scalar
	return validation.Validate(
		([]Pet)(s),
		eachWithIndirection(),
	)

}

// ValidatePetsResponseDefault defines parameters for ValidatePets.
type ValidatePetsResponseDefault = Error

// ValidatePetsJSONRequestBody defines body for ValidatePets for application/json ContentType.
type ValidatePetsJSONRequestBody ValidatePetsJSONBody

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
	// GetPet request
	GetPet(ctx context.Context, petId GetPetPathPetId, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ValidatePets request  with any body
	ValidatePetsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	ValidatePets(ctx context.Context, body ValidatePetsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetPet(ctx context.Context, petId GetPetPathPetId, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetPetRequest(c.Server, petId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ValidatePetsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewValidatePetsRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ValidatePets(ctx context.Context, body ValidatePetsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewValidatePetsRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetPetRequest generates requests for GetPet
func NewGetPetRequest(server string, petId GetPetPathPetId) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "petId", runtime.ParamLocationPath, petId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/pets/%s", pathParam0)
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

// NewValidatePetsRequest calls the generic ValidatePets builder with application/json body
func NewValidatePetsRequest(server string, body ValidatePetsJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewValidatePetsRequestWithBody(server, "application/json", bodyReader)
}

// NewValidatePetsRequestWithBody generates requests for ValidatePets with any type of body
func NewValidatePetsRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/pets:validate")
	if operationPath[0] == '/' {
		operationPath = operationPath[1:]
	}
	operationURL := url.URL{
		Path: operationPath,
	}

	queryURL := serverURL.ResolveReference(&operationURL)

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

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
	// GetPet request
	GetPetWithResponse(ctx context.Context, petId GetPetPathPetId, reqEditors ...RequestEditorFn) (*GetPetResponse, error)

	// ValidatePets request  with any body
	ValidatePetsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*ValidatePetsResponse, error)

	ValidatePetsWithResponse(ctx context.Context, body ValidatePetsJSONRequestBody, reqEditors ...RequestEditorFn) (*ValidatePetsResponse, error)
}

// GetPetResponseJSON200 represents a possible response for the GetPet request.
type GetPetResponseJSON200 Pet
type GetPetResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetPetResponseJSON200
}

// Status returns HTTPResponse.Status
func (r GetPetResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetPetResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ValidatePetsResponseJSON200 represents a possible response for the ValidatePets request.
type ValidatePetsResponseJSON200 []Pet

// ValidatePetsResponseJSONDefault represents a possible response for the ValidatePets request.
type ValidatePetsResponseJSONDefault Error
type ValidatePetsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ValidatePetsResponseJSON200
	JSONDefault  *ValidatePetsResponseJSONDefault
}

// Status returns HTTPResponse.Status
func (r ValidatePetsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ValidatePetsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetPetWithResponse request returning *GetPetResponse
func (c *ClientWithResponses) GetPetWithResponse(ctx context.Context, petId GetPetPathPetId, reqEditors ...RequestEditorFn) (*GetPetResponse, error) {
	rsp, err := c.GetPet(ctx, petId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetPetResponse(rsp)
}

// ValidatePetsWithBodyWithResponse request with arbitrary body returning *ValidatePetsResponse
func (c *ClientWithResponses) ValidatePetsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*ValidatePetsResponse, error) {
	rsp, err := c.ValidatePetsWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseValidatePetsResponse(rsp)
}

func (c *ClientWithResponses) ValidatePetsWithResponse(ctx context.Context, body ValidatePetsJSONRequestBody, reqEditors ...RequestEditorFn) (*ValidatePetsResponse, error) {
	rsp, err := c.ValidatePets(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseValidatePetsResponse(rsp)
}

// ParseGetPetResponse parses an HTTP response from a GetPetWithResponse call
func ParseGetPetResponse(rsp *http.Response) (*GetPetResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetPetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetPetResponseJSON200
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseValidatePetsResponse parses an HTTP response from a ValidatePetsWithResponse call
func ParseValidatePetsResponse(rsp *http.Response) (*ValidatePetsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &ValidatePetsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ValidatePetsResponseJSON200
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest ValidatePetsResponseJSONDefault
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get pet given identifier.
	// (GET /pets/{petId})
	GetPet(ctx *GetPetContext, petId GetPetPathPetId) error
	// Validate pets
	// (POST /pets:validate)
	ValidatePets(ctx *ValidatePetsContext) error
}

// GetPetContext is a context customized for GetPet (GET /pets/{petId}).
type GetPetContext struct {
	echo.Context
}

// Responses

// OK responses with the appropriate code and the JSON response.
func (c *GetPetContext) OK(resp GetPetResponseOK) error {
	return c.Context.JSON(200, resp)
}

// ValidatePetsContext is a context customized for ValidatePets (POST /pets:validate).
type ValidatePetsContext struct {
	echo.Context
}

// The body parsers
// ParseJSONBody tries to parse the body into the respective structure and validate it.
func (c *ValidatePetsContext) ParseJSONBody() (ValidatePetsJSONBody, error) {
	var resp ValidatePetsJSONBody
	return resp, bindValidateBody(c.Context, &resp)
}

// Responses

// OK responses with the appropriate code and the JSON response.
func (c *ValidatePetsContext) OK(resp ValidatePetsResponseOK) error {
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

// GetPet converts echo context to params.
func (w *ServerInterfaceWrapper) GetPet(ctx echo.Context) error {
	var err error

	// ------------- Path parameter "petId" -------------
	var petId GetPetPathPetId

	err = runtime.BindStyledParameterWithLocation("simple", false, "petId", runtime.ParamLocationPath, ctx.Param("petId"), &petId)
	if err != nil {
		return errors.WithStack(&ValidationError{ParamType: "path", Param: "petId", Err: errors.Wrap(err, "invalid format")})
	}

	if err := petId.Validate(); err != nil {
		return errors.WithStack(&ValidationError{ParamType: "path", Param: "petId", Err: err})
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPet(&GetPetContext{ctx}, petId)
	return err
}

// ValidatePets converts echo context to params.
func (w *ServerInterfaceWrapper) ValidatePets(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ValidatePets(&ValidatePetsContext{ctx})
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

	router.GET(baseURL+"/pets/:petId", wrapper.GetPet, m...)
	router.POST(baseURL+"/pets:validate", wrapper.ValidatePets, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/6xUPW/bMBD9K8S1o2A58aaxRVF4KTIUXQIPrHSSmZof5Z2MGgb/e3Gk7NqxjAzNFIX3",
	"eO/de0cfofU2eIeOCZojULtFq/Pnlxh9lI8QfcDIBvNx6zuUvx1SG01g4x00BaxyrYLeR6sZGjCOV49Q",
	"AR8Cln9xwAipAotEerjb6FQ+XyWOxg2QUgURf48mYgfNM0yEJ/gmVfCEfCvaaTvD9X2LSirK94q3qALy",
	"4k3K3GpzRvmfL9gyFOJv2ha+W3a6T08X/CQCDKPN+FdKzqQ6Rn2YVUYz0gRnXO9vFXzeYvuLVFGrkFod",
	"jBtETtBRW2SMJIYY3knDNdGIavXwqBiJoYI9RiqdHhbLxVIU+oBOBwMNrPJRBUHzNk9Ty3z1MSCvuyQH",
	"Q4lKyLUoWnfQwFdkiVDunSU0z0cwQiO9oJrihNwJLk3gOGI1LfGMgWkjYAreUQnkcbksO+0YXRajQ9iZ",
	"NsupX0hmO170+xixhwY+1P9eTT09mVpUZ6+vPd7rnekk2pwXjdbqeChzyqkazB6dMh06Nr3BuMi47FWT",
	"72rOqxs88W2CPyZE3h2oXnl5qj6VoviExJ98d3jPqcvWp+t1lCTSf7p9fgdv2n7zMu7HQJBrvR53/G4u",
	"lN/KGdrR4Z+ALWOncMJcLsF1fCml9DcAAP//Z6kfG5EFAAA=",
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
