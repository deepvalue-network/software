package errors

import "github.com/steve-care-software/products/bobby/domain/resources"

type errorStruct struct {
	resource resources.Immutable
	message  string
	code     uint
	parent   Error
}

func createError(
	resource resources.Immutable,
	message string,
	code uint,
) Error {
	return createErrorInternally(resource, message, code, nil)
}

func createErrorWithParent(
	resource resources.Immutable,
	message string,
	code uint,
	parent Error,
) Error {
	return createErrorInternally(resource, message, code, parent)
}

func createErrorInternally(
	resource resources.Immutable,
	message string,
	code uint,
	parent Error,
) Error {
	out := errorStruct{
		resource: resource,
		message:  message,
		code:     code,
		parent:   parent,
	}

	return &out
}

// Resource returns the resource
func (obj *errorStruct) Resource() resources.Immutable {
	return obj.resource
}

// Message returns the message
func (obj *errorStruct) Message() string {
	return obj.message
}

// Code returns the code
func (obj *errorStruct) Code() uint {
	return obj.code
}

// HasParent returns true if there is a parent, false otherwise
func (obj *errorStruct) HasParent() bool {
	return obj.parent != nil
}

// Parent returns the parent, if any
func (obj *errorStruct) Parent() Error {
	return obj.parent
}
