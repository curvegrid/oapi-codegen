// Package api provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

import (
	"github.com/go-ozzo/ozzo-validation/v4"
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

// NewPet defines model for NewPet.
type NewPet struct {

	// Name of the pet
	Name string `json:"name"`

	// Type of the pet
	Tag *string `json:"tag,omitempty"`
}

// Validate perform validation on the NewPet
func (s NewPet) Validate() error {
	// Run validate on a struct
	return validation.ValidateStruct(
		&s,
		validation.Field(
			&s.Name,
			validation.Required,
		),
		validation.Field(
			&s.Tag,
		),
	)

}

// Pet defines model for Pet.
type Pet struct {
	// Embedded struct due to allOf(#/components/schemas/NewPet)
	NewPet
	// Embedded fields due to inline allOf schema

	// Unique id of the pet
	Id int64 `json:"id"`
}

// Validate perform validation on the Pet
func (s Pet) Validate() error {
	// Run validate on a struct
	return validation.ValidateStruct(
		&s, validation.Field(&s.NewPet),
		validation.Field(
			&s.Id,
			validation.Required,
		),
	)

}

// FindPetsParams defines parameters for FindPets.
type FindPetsParams struct {

	// tags to filter by
	Tags *[]string `json:"tags,omitempty"`

	// maximum number of results to return
	Limit *int32 `json:"limit,omitempty"`
}

// Validate perform validation on the FindPetsParams
func (s FindPetsParams) Validate() error {
	// Run validate on a struct
	return validation.ValidateStruct(
		&s,
		validation.Field(
			&s.Tags,

			validation.Each(),
		),
		validation.Field(
			&s.Limit,
		),
	)

}

// AddPetJSONBody defines parameters for AddPet.
type AddPetJSONBody NewPet

// Validate perform validation on the AddPetJSONBody
func (s AddPetJSONBody) Validate() error {
	// Run validate on a scalar
	return validation.Validate(
		(NewPet)(s),
	)

}

// AddPetRequestBody defines body for AddPet for application/json ContentType.
type AddPetJSONRequestBody AddPetJSONBody
