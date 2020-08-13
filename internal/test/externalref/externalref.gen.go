// Package externalref provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package externalref

import (
	externalRef0 "github.com/deepmap/oapi-codegen/internal/test/externalref/packageA"
	externalRef1 "github.com/deepmap/oapi-codegen/internal/test/externalref/packageB"
	"github.com/go-ozzo/ozzo-validation/v4"
)

// Container defines model for Container.
type Container struct {
	ObjectA *externalRef0.ObjectA `json:"object_a,omitempty"`
	ObjectB *externalRef1.ObjectB `json:"object_b,omitempty"`
}

// Validate perform validation on the Container
func (s Container) Validate() error {
	// Run validate on a struct
	return validation.ValidateStruct(
		&s,
		validation.Field(
			&s.ObjectA,
		),
		validation.Field(
			&s.ObjectB,
		),
	)

}