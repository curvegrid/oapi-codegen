// Package api provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"strings"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Returns all pets
	// (GET /pets)
	FindPets(ctx *FindPetsContext, params FindPetsParams) error
	// Creates a new pet
	// (POST /pets)
	AddPet(ctx *AddPetContext) error
	// Deletes a pet by ID
	// (DELETE /pets/{id})
	DeletePet(ctx *DeletePetContext, id DeletePetPathId) error
	// Returns a pet by ID
	// (GET /pets/{id})
	FindPetById(ctx *FindPetByIdContext, id FindPetByIdPathId) error
}

// FindPetsContext is a context customized for FindPets (GET /pets).
type FindPetsContext struct {
	echo.Context
}

// Responses

// OK responses with the appropriate code and the JSON response.
func (c *FindPetsContext) OK(resp FindPetsResponseOK) error {
	return c.JSON(200, resp)
}

// FindPetsResponseOK is the response type for FindPets's "200" response.
type FindPetsResponseOK = []Pet

// AddPetContext is a context customized for AddPet (POST /pets).
type AddPetContext struct {
	echo.Context
}

// The body parsers
// ParseJSONBody tries to parse the body into the respective structure and validate it.
func (c *AddPetContext) ParseJSONBody() (AddPetJSONBody, error) {
	var resp AddPetJSONBody
	if err := c.Bind(&resp); err != nil {
		return resp, ValidationError{ParamType: "body", Err: errors.Wrap(err, "cannot parse as json")}
	}
	if err := resp.Validate(); err != nil {
		return resp, ValidationError{ParamType: "body", Err: err}
	}
	return resp, nil
}

// Responses

// OK responses with the appropriate code and the JSON response.
func (c *AddPetContext) OK(resp Pet) error {
	return c.JSON(200, resp)
}

// DeletePetContext is a context customized for DeletePet (DELETE /pets/{id}).
type DeletePetContext struct {
	echo.Context
}

// Responses

// FindPetByIdContext is a context customized for FindPetById (GET /pets/{id}).
type FindPetByIdContext struct {
	echo.Context
}

// Responses

// OK responses with the appropriate code and the JSON response.
func (c *FindPetByIdContext) OK(resp Pet) error {
	return c.JSON(200, resp)
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

// handleFindPets converts echo context to params.
func (w *ServerInterfaceWrapper) handleFindPets(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params FindPetsParams
	// ------------- Optional query parameter "tags" -------------

	err = runtime.BindQueryParameter("form", true, false, "tags", ctx.QueryParams(), &params.Tags)
	if err != nil {
		return errors.WithStack(ValidationError{ParamType: "query", Err: errors.Wrap(err, "invalid format")})
	}

	if err := params.Validate(); err != nil {
		return errors.WithStack(ValidationError{ParamType: "query", Err: err})
	}
	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return errors.WithStack(ValidationError{ParamType: "query", Err: errors.Wrap(err, "invalid format")})
	}

	if err := params.Validate(); err != nil {
		return errors.WithStack(ValidationError{ParamType: "query", Err: err})
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FindPets(&FindPetsContext{ctx}, params)
	return err
}

// FindPets creates a handler function for the endpoint.
func (w *ServerInterfaceWrapper) FindPets() echo.HandlerFunc {
	securityReqs := BindSecurityRequirements()
	// Wrap handler in middlewares
	handler := echo.HandlerFunc(w.handleFindPets)
	for i := len(w.middlewares); i > 0; i-- {
		handler = w.middlewares[i-1](handler)
	}
	// Put securityReqs on top
	return securityReqs(handler)
}

// handleAddPet converts echo context to params.
func (w *ServerInterfaceWrapper) handleAddPet(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddPet(&AddPetContext{ctx})
	return err
}

// AddPet creates a handler function for the endpoint.
func (w *ServerInterfaceWrapper) AddPet() echo.HandlerFunc {
	securityReqs := BindSecurityRequirements()
	// Wrap handler in middlewares
	handler := echo.HandlerFunc(w.handleAddPet)
	for i := len(w.middlewares); i > 0; i-- {
		handler = w.middlewares[i-1](handler)
	}
	// Put securityReqs on top
	return securityReqs(handler)
}

// handleDeletePet converts echo context to params.
func (w *ServerInterfaceWrapper) handleDeletePet(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id DeletePetPathId

	err = runtime.BindStyledParameter("simple", false, "id", ctx.Param("id"), &id)
	if err != nil {
		return errors.WithStack(ValidationError{ParamType: "path", Param: "id", Err: errors.Wrap(err, "invalid format")})
	}

	if err := id.Validate(); err != nil {
		return errors.WithStack(ValidationError{ParamType: "path", Param: "id", Err: err})
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeletePet(&DeletePetContext{ctx}, id)
	return err
}

// DeletePet creates a handler function for the endpoint.
func (w *ServerInterfaceWrapper) DeletePet() echo.HandlerFunc {
	securityReqs := BindSecurityRequirements()
	// Wrap handler in middlewares
	handler := echo.HandlerFunc(w.handleDeletePet)
	for i := len(w.middlewares); i > 0; i-- {
		handler = w.middlewares[i-1](handler)
	}
	// Put securityReqs on top
	return securityReqs(handler)
}

// handleFindPetById converts echo context to params.
func (w *ServerInterfaceWrapper) handleFindPetById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id FindPetByIdPathId

	err = runtime.BindStyledParameter("simple", false, "id", ctx.Param("id"), &id)
	if err != nil {
		return errors.WithStack(ValidationError{ParamType: "path", Param: "id", Err: errors.Wrap(err, "invalid format")})
	}

	if err := id.Validate(); err != nil {
		return errors.WithStack(ValidationError{ParamType: "path", Param: "id", Err: err})
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FindPetById(&FindPetByIdContext{ctx}, id)
	return err
}

// FindPetById creates a handler function for the endpoint.
func (w *ServerInterfaceWrapper) FindPetById() echo.HandlerFunc {
	securityReqs := BindSecurityRequirements()
	// Wrap handler in middlewares
	handler := echo.HandlerFunc(w.handleFindPetById)
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

	router.GET("/pets", wrapper.FindPets())
	router.POST("/pets", wrapper.AddPet())
	router.DELETE("/pets/:id", wrapper.DeletePet())
	router.GET("/pets/:id", wrapper.FindPetById())

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

	"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",
	"u/vdPwIAAP//v4qmX1cSAAA=",
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
