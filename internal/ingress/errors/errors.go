/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package errors

import (
	"errors"
	"fmt"
)

var (
	// ErrMissingAnnotations the ingress rule does not contain annotations
	// This is an error only when annotations are being parsed
	ErrMissingAnnotations = errors.New("ingress rule without annotations")

	// ErrInvalidAnnotationName the ingress rule does contains an invalid
	// annotation name
	ErrInvalidAnnotationName = errors.New("invalid annotation name")
)

// NewInvalidAnnotationConfiguration returns a new InvalidConfiguration error for use when
// annotations are not correctly configured
func NewInvalidAnnotationConfiguration(name, reason string) error {
	return InvalidConfigurationError{
		Name: fmt.Sprintf("the annotation %v does not contain a valid configuration: %v", name, reason),
	}
}

// NewInvalidAnnotationContent returns a new InvalidContent error
func NewInvalidAnnotationContent(name string, val interface{}) error {
	return InvalidContentError{
		Name: fmt.Sprintf("the annotation %v does not contain a valid value (%v)", name, val),
	}
}

// NewLocationDenied returns a new LocationDenied error
func NewLocationDenied(reason string) error {
	return LocationDeniedError{
		Reason: fmt.Errorf("location denied, reason: %v", reason),
	}
}

// InvalidConfigurationError
type InvalidConfigurationError struct {
	Name string
}

func (e InvalidConfigurationError) Error() string {
	return e.Name
}

// InvalidContentError
type InvalidContentError struct {
	Name string
}

func (e InvalidContentError) Error() string {
	return e.Name
}

// LocationDeniedError
type LocationDeniedError struct {
	Reason error
}

func (e LocationDeniedError) Error() string {
	return e.Reason.Error()
}

// IsLocationDenied checks if the err is an error which
// indicates a location should return HTTP code 503
func IsLocationDenied(e error) bool {
	_, ok := e.(LocationDeniedError)
	return ok
}

// IsMissingAnnotations checks if the err is an error which
// indicates the ingress does not contain annotations
func IsMissingAnnotations(e error) bool {
	return e == ErrMissingAnnotations
}

// IsInvalidContent checks if the err is an error which
// indicates an annotations value is not valid
func IsInvalidContent(e error) bool {
	_, ok := e.(InvalidContentError)
	return ok
}

// New returns a new error
func New(m string) error {
	return errors.New(m)
}

// Errorf formats according to a format specifier and returns the string
// as a value that satisfies error.
func Errorf(format string, args ...interface{}) error {
	return fmt.Errorf(format, args...)
}

type ValidationError struct {
	Reason error
}

type RiskyAnnotationError struct {
	Reason error
}

func (e ValidationError) Error() string {
	return e.Reason.Error()
}

// NewValidationError returns a new LocationDenied error
func NewValidationError(annotation string) error {
	return ValidationError{
		Reason: fmt.Errorf("annotation %s contains invalid value", annotation),
	}
}

// IsValidationError checks if the err is an error which
// indicates that some annotation value is invalid
func IsValidationError(e error) bool {
	_, ok := e.(ValidationError)
	return ok
}

// NewRiskyAnnotations returns a new LocationDenied error
func NewRiskyAnnotations(name string) error {
	return RiskyAnnotationError{
		Reason: fmt.Errorf("annotation group %s contains risky annotation based on ingress configuration", name),
	}
}

// IsRiskyAnnotationError checks if the err is an error which
// indicates that some annotation value is invalid
func IsRiskyAnnotationError(e error) bool {
	_, ok := e.(ValidationError)
	return ok
}

func (e RiskyAnnotationError) Error() string {
	return e.Reason.Error()
}
