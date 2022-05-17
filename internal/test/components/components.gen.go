// Package components provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version (devel) DO NOT EDIT.
package components

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Defines values for Enum1.
const (
	Enum1One Enum1 = "One"

	Enum1Three Enum1 = "Three"

	Enum1Two Enum1 = "Two"
)

// Defines values for Enum2.
const (
	Enum2One Enum2 = "One"

	Enum2Three Enum2 = "Three"

	Enum2Two Enum2 = "Two"
)

// Defines values for Enum3.
const (
	Enum3Bar Enum3 = "Bar"

	Enum3Enum1One Enum3 = "Enum1One"

	Enum3Foo Enum3 = "Foo"
)

// Defines values for Enum4.
const (
	Cat Enum4 = "Cat"

	Dog Enum4 = "Dog"

	Mouse Enum4 = "Mouse"
)

// Defines values for Enum5.
const (
	N5 Enum5 = 5

	N6 Enum5 = 6

	N7 Enum5 = 7
)

// Defines values for EnumParam1.
const (
	EnumParam1Both EnumParam1 = "both"

	EnumParam1False EnumParam1 = "false"

	EnumParam1True EnumParam1 = "true"
)

// Defines values for EnumParam2.
const (
	EnumParam2Both EnumParam2 = "both"

	EnumParam2False EnumParam2 = "false"

	EnumParam2True EnumParam2 = "true"
)

// Defines values for EnumParam3.
const (
	Alice EnumParam3 = "alice"

	Bob EnumParam3 = "bob"

	Eve EnumParam3 = "eve"
)

// Has additional properties of type int
type AdditionalPropertiesObject1 struct {
	Id                   int            `json:"id"`
	Name                 string         `json:"name"`
	Optional             *string        `json:"optional,omitempty"`
	AdditionalProperties map[string]int `json:"-"`
}

// Does not allow additional properties
type AdditionalPropertiesObject2 struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// Allows any additional property
type AdditionalPropertiesObject3 struct {
	Name                 string                 `json:"name"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// Has anonymous field which has additional properties
type AdditionalPropertiesObject4 struct {
	Inner                AdditionalPropertiesObject4_Inner `json:"inner"`
	Name                 string                            `json:"name"`
	AdditionalProperties map[string]interface{}            `json:"-"`
}

// AdditionalPropertiesObject4_Inner defines model for AdditionalPropertiesObject4.Inner.
type AdditionalPropertiesObject4_Inner struct {
	Name                 string                 `json:"name"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// Has additional properties with schema for dictionaries
type AdditionalPropertiesObject5 struct {
	AdditionalProperties map[string]SchemaObject `json:"-"`
}

// Conflicts with Enum2, enum values need to be prefixed with type
// name.
type Enum1 string

// Conflicts with Enum1, enum values need to be prefixed with type
// name.
type Enum2 string

// Enum values conflict with Enums above, need to be prefixed
// with type name.
type Enum3 string

// No conflicts here, should have unmodified enums
type Enum4 string

// Numerical enum
type Enum5 int

// ObjectWithJsonField defines model for ObjectWithJsonField.
type ObjectWithJsonField struct {
	Name   string          `json:"name"`
	Value1 json.RawMessage `json:"value1"`
	Value2 json.RawMessage `json:"value2,omitempty"`
}

// SchemaObject defines model for SchemaObject.
type SchemaObject struct {
	FirstName string `json:"firstName"`

	// This property is required and readOnly, so the go model should have it as a pointer,
	// as it will not be included when it is sent from client to server.
	ReadOnlyRequiredProp  *string `json:"readOnlyRequiredProp,omitempty"`
	Role                  string  `json:"role"`
	WriteOnlyRequiredProp *int    `json:"writeOnlyRequiredProp,omitempty"`
}

// EnumParam1 defines model for EnumParam1.
type EnumParam1 string

// EnumParam2 defines model for EnumParam2.
type EnumParam2 string

// EnumParam3 defines model for EnumParam3.
type EnumParam3 string

// a parameter
type ParameterObject string

// ResponseObject defines model for ResponseObject.
type ResponseObject struct {
	Field SchemaObject `json:"Field"`
}

// RequestBody defines model for RequestBody.
type RequestBody struct {
	Field SchemaObject `json:"Field"`
}

// ParamsWithAddPropsParams_P1 defines parameters for ParamsWithAddProps.
type ParamsWithAddPropsParams_P1 struct {
	AdditionalProperties map[string]interface{} `json:"-"`
}

// ParamsWithAddPropsParams defines parameters for ParamsWithAddProps.
type ParamsWithAddPropsParams struct {
	// This parameter has additional properties
	P1 ParamsWithAddPropsParams_P1 `json:"p1"`

	// This parameter has an anonymous inner property which needs to be
	// turned into a proper type for additionalProperties to work
	P2 struct {
		Inner ParamsWithAddPropsParams_P2_Inner `json:"inner"`
	} `json:"p2"`
}

// ParamsWithAddPropsParams_P2_Inner defines parameters for ParamsWithAddProps.
type ParamsWithAddPropsParams_P2_Inner struct {
	AdditionalProperties map[string]string `json:"-"`
}

// BodyWithAddPropsJSONBody defines parameters for BodyWithAddProps.
type BodyWithAddPropsJSONBody struct {
	Inner                BodyWithAddPropsJSONBody_Inner `json:"inner"`
	Name                 string                         `json:"name"`
	AdditionalProperties map[string]interface{}         `json:"-"`
}

// BodyWithAddPropsJSONBody_Inner defines parameters for BodyWithAddProps.
type BodyWithAddPropsJSONBody_Inner struct {
	AdditionalProperties map[string]int `json:"-"`
}

// EnsureEverythingIsReferencedJSONRequestBody defines body for EnsureEverythingIsReferenced for application/json ContentType.
type EnsureEverythingIsReferencedJSONRequestBody RequestBody

// BodyWithAddPropsJSONRequestBody defines body for BodyWithAddProps for application/json ContentType.
type BodyWithAddPropsJSONRequestBody BodyWithAddPropsJSONBody

// Getter for additional properties for ParamsWithAddPropsParams_P1. Returns the specified
// element and whether it was found
func (a ParamsWithAddPropsParams_P1) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for ParamsWithAddPropsParams_P1
func (a *ParamsWithAddPropsParams_P1) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for ParamsWithAddPropsParams_P1 to handle AdditionalProperties
func (a *ParamsWithAddPropsParams_P1) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for ParamsWithAddPropsParams_P1 to handle AdditionalProperties
func (a ParamsWithAddPropsParams_P1) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for ParamsWithAddPropsParams_P2_Inner. Returns the specified
// element and whether it was found
func (a ParamsWithAddPropsParams_P2_Inner) Get(fieldName string) (value string, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for ParamsWithAddPropsParams_P2_Inner
func (a *ParamsWithAddPropsParams_P2_Inner) Set(fieldName string, value string) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]string)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for ParamsWithAddPropsParams_P2_Inner to handle AdditionalProperties
func (a *ParamsWithAddPropsParams_P2_Inner) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]string)
		for fieldName, fieldBuf := range object {
			var fieldVal string
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for ParamsWithAddPropsParams_P2_Inner to handle AdditionalProperties
func (a ParamsWithAddPropsParams_P2_Inner) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for BodyWithAddPropsJSONBody. Returns the specified
// element and whether it was found
func (a BodyWithAddPropsJSONBody) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for BodyWithAddPropsJSONBody
func (a *BodyWithAddPropsJSONBody) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for BodyWithAddPropsJSONBody to handle AdditionalProperties
func (a *BodyWithAddPropsJSONBody) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["inner"]; found {
		err = json.Unmarshal(raw, &a.Inner)
		if err != nil {
			return fmt.Errorf("error reading 'inner': %w", err)
		}
		delete(object, "inner")
	}

	if raw, found := object["name"]; found {
		err = json.Unmarshal(raw, &a.Name)
		if err != nil {
			return fmt.Errorf("error reading 'name': %w", err)
		}
		delete(object, "name")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for BodyWithAddPropsJSONBody to handle AdditionalProperties
func (a BodyWithAddPropsJSONBody) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["inner"], err = json.Marshal(a.Inner)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'inner': %w", err)
	}

	object["name"], err = json.Marshal(a.Name)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'name': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for BodyWithAddPropsJSONBody_Inner. Returns the specified
// element and whether it was found
func (a BodyWithAddPropsJSONBody_Inner) Get(fieldName string) (value int, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for BodyWithAddPropsJSONBody_Inner
func (a *BodyWithAddPropsJSONBody_Inner) Set(fieldName string, value int) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]int)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for BodyWithAddPropsJSONBody_Inner to handle AdditionalProperties
func (a *BodyWithAddPropsJSONBody_Inner) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]int)
		for fieldName, fieldBuf := range object {
			var fieldVal int
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for BodyWithAddPropsJSONBody_Inner to handle AdditionalProperties
func (a BodyWithAddPropsJSONBody_Inner) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for AdditionalPropertiesObject1. Returns the specified
// element and whether it was found
func (a AdditionalPropertiesObject1) Get(fieldName string) (value int, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for AdditionalPropertiesObject1
func (a *AdditionalPropertiesObject1) Set(fieldName string, value int) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]int)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for AdditionalPropertiesObject1 to handle AdditionalProperties
func (a *AdditionalPropertiesObject1) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["id"]; found {
		err = json.Unmarshal(raw, &a.Id)
		if err != nil {
			return fmt.Errorf("error reading 'id': %w", err)
		}
		delete(object, "id")
	}

	if raw, found := object["name"]; found {
		err = json.Unmarshal(raw, &a.Name)
		if err != nil {
			return fmt.Errorf("error reading 'name': %w", err)
		}
		delete(object, "name")
	}

	if raw, found := object["optional"]; found {
		err = json.Unmarshal(raw, &a.Optional)
		if err != nil {
			return fmt.Errorf("error reading 'optional': %w", err)
		}
		delete(object, "optional")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]int)
		for fieldName, fieldBuf := range object {
			var fieldVal int
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for AdditionalPropertiesObject1 to handle AdditionalProperties
func (a AdditionalPropertiesObject1) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["id"], err = json.Marshal(a.Id)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'id': %w", err)
	}

	object["name"], err = json.Marshal(a.Name)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'name': %w", err)
	}

	if a.Optional != nil {
		object["optional"], err = json.Marshal(a.Optional)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'optional': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for AdditionalPropertiesObject3. Returns the specified
// element and whether it was found
func (a AdditionalPropertiesObject3) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for AdditionalPropertiesObject3
func (a *AdditionalPropertiesObject3) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for AdditionalPropertiesObject3 to handle AdditionalProperties
func (a *AdditionalPropertiesObject3) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["name"]; found {
		err = json.Unmarshal(raw, &a.Name)
		if err != nil {
			return fmt.Errorf("error reading 'name': %w", err)
		}
		delete(object, "name")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for AdditionalPropertiesObject3 to handle AdditionalProperties
func (a AdditionalPropertiesObject3) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["name"], err = json.Marshal(a.Name)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'name': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for AdditionalPropertiesObject4. Returns the specified
// element and whether it was found
func (a AdditionalPropertiesObject4) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for AdditionalPropertiesObject4
func (a *AdditionalPropertiesObject4) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for AdditionalPropertiesObject4 to handle AdditionalProperties
func (a *AdditionalPropertiesObject4) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["inner"]; found {
		err = json.Unmarshal(raw, &a.Inner)
		if err != nil {
			return fmt.Errorf("error reading 'inner': %w", err)
		}
		delete(object, "inner")
	}

	if raw, found := object["name"]; found {
		err = json.Unmarshal(raw, &a.Name)
		if err != nil {
			return fmt.Errorf("error reading 'name': %w", err)
		}
		delete(object, "name")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for AdditionalPropertiesObject4 to handle AdditionalProperties
func (a AdditionalPropertiesObject4) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["inner"], err = json.Marshal(a.Inner)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'inner': %w", err)
	}

	object["name"], err = json.Marshal(a.Name)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'name': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for AdditionalPropertiesObject4_Inner. Returns the specified
// element and whether it was found
func (a AdditionalPropertiesObject4_Inner) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for AdditionalPropertiesObject4_Inner
func (a *AdditionalPropertiesObject4_Inner) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for AdditionalPropertiesObject4_Inner to handle AdditionalProperties
func (a *AdditionalPropertiesObject4_Inner) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["name"]; found {
		err = json.Unmarshal(raw, &a.Name)
		if err != nil {
			return fmt.Errorf("error reading 'name': %w", err)
		}
		delete(object, "name")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for AdditionalPropertiesObject4_Inner to handle AdditionalProperties
func (a AdditionalPropertiesObject4_Inner) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["name"], err = json.Marshal(a.Name)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'name': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for AdditionalPropertiesObject5. Returns the specified
// element and whether it was found
func (a AdditionalPropertiesObject5) Get(fieldName string) (value SchemaObject, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for AdditionalPropertiesObject5
func (a *AdditionalPropertiesObject5) Set(fieldName string, value SchemaObject) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]SchemaObject)
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for AdditionalPropertiesObject5 to handle AdditionalProperties
func (a *AdditionalPropertiesObject5) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]SchemaObject)
		for fieldName, fieldBuf := range object {
			var fieldVal SchemaObject
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for AdditionalPropertiesObject5 to handle AdditionalProperties
func (a AdditionalPropertiesObject5) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9xYW2/bOhL+KwPuPrJx4rS7gN/SG7YL9II2wHmog4IWRxF7ZFIlKbtG4f9+MKRsyRLl",
	"Ok5R4JyXxJbJmW9u38zoB8vMsjIatXds9oNVwoolerTh2ytdLz/Qkyv6JtFlVlVeGc1mTMD+LONM0aNv",
	"NdoN40yLJbIZCzdJxBXjzGUFLgWJQV0v2eyztzXyXJQOOVsYX7A7zvymoovOW6Xv2XbLWwTTRyGY/gIE",
	"149CcJ1CwESpMmSc4Yr+LswijeHDTsv7xVfMPEnIjPaow0dRVaXKBGGafHUE7EdHVWVNhdYrDBF9rbCU",
	"9OHfFnM2Y/+atPGfxEtu8in8b3SRfovfamVREuQo4Y4ee/zuJ1UpVE9l34AtP89vHYt3KND550Y2xnzc",
	"P9j87VwSZbjKaLczJn75h4T4BpxaViXCzkgwrbIGBQm6kVLRFVF+2FsRYQXOEYmfO/qV9niPlg3U/084",
	"aO9C6yEwOdBlUNoz3nOdkmnZMSkHVnNmqqgg5ZJDnwYRnDS0Fb7zCD/ihem4FxrqOjT8pUEH2ngQZWnW",
	"aR881u5fZNr1uGmBmvsZRQY5EHqTsGozsOkB2B8G++nDYIdM1EZvlqZ2kFNpwbpQWQHFWI4O46M12p+p",
	"/aXmn3Y94uLnePHZseo+nblOr/u18gVEIZAbC1Jl4ZCNDh9Aj4PLoOW/MDovVeYbgWG44ED9HFairKn4",
	"ECV4AwuEymKuvqOMZ0nFXJOvLuaamn4zBLzXVD63a0N/C4uYGAIinulJeK5+G57ESPSqozlrsLXQHIiF",
	"WSFPgZrrPSoYggqGRWSvDSF7LuworqdDXO/MHo6DAi1ycIWpSwmFWCHUemmkyhXK4Dt3oPuFoFbx0twz",
	"zt6a2o075FlCcb1EqzJRBsFduc/4f/h/W0kd3o3Z/Yfyxf+d0fuGflKFcxa8H1I3N3YpPJuxMDPwkaPT",
	"E46mKb/RlKr8gzIdYM+Vdf7dmAEWhXyvy83HRiORw9Cxt4Vye/oH5WAHEISWsJPBwRnwBcK9gaWRWB6E",
	"XXkgwoDKkPcth7kWjp6uVVmGNrqgSSEra0lVU6CmH5UDh9pDbs0SslLRZ2/AoV2hjWm7U78j56GNpkwb",
	"v7bKY8r6Xpp0TkYl/RgFDbzj6hHHjqm8C1Od0rkh7bSoaIdt4rG3b24Jr1eeLGG36Dx8Ci6gzEDrYpiu",
	"Li4vLuOshFpUis3Y9cXlBe2ElfBFSIcJaldbfIIrtBtfKH3/RLknFnO0qDMMyX+PfiQHUMsQP8DvynkX",
	"Ay48tM0DMqEpkplF4VGC0uAL5ebaVZiFdGlCXdlaExOxANeGgfuNDLRGAF/t8b1xH1t03dVkM9a/DraX",
	"SXd16W8C08vLR4z/uVrhz3rosba85Sw3tT1fxFMS8bXLW8fkpKiOkkU/woirkJehc50t4zrIWJvzJUwD",
	"a/aIcbgoVaXIsDCl3K0xuahLP55FTaJMevti3CgnYa92X6iTfhFSfqHccKPlcwNUgrE9t29+QkEIWBi5",
	"aSbVhlbSk1WiWsL27iioNzKQSRhpO6+WPqfJfHdifDQOylIvDKor1qW/yLptjRwbnAe9y/lNoLS4wbIt",
	"PwWt7kz5YS5ue1N0Io08Ls48c+1rqwMReUPtJ5yMow/Npim0dHNt7J/jHpge9cCDNopEX+oncmoTuNtu",
	"73pk1snnDqGd/VrhoFro58q4RGKHNQAayu1mMmEQStOvUlnMfNLXnEpgro/GlGomdTdRDsTyvWKwZ766",
	"On0BPDXCnbHzzC1wt//vUmDbT8Ptb82J7Xb7VwAAAP//Ati2mNYWAAA=",
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
