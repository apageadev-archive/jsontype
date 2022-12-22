package jsontype

import (
	"fmt"
	"strings"

	"github.com/goccy/go-json"
	"github.com/gookit/validate"
)

// SchemaManager provides methods for managing schemas
// it also allows for custom types to be used withing the schemas
// as this is the centralized store for all schemas
type SchemaManager struct {
	Schemas map[string]*Schema
}

// A Schema defines an entity and is the atomic unit of the JSONType package
type Schema struct {
	Type                     string              `json:"type,omitempty" validate:"required"`
	Description              string              `json:"description,omitempty"`
	Extends                  string              `json:"extends,omitempty"`
	Properties               map[string]Property `json:"properties" validate:"required"`
	OptionalProperties       map[string]Property `json:"optional_properties,omitempty"`
	AllowUndefinedProperties bool                `json:"allow_undefined_properties,omitempty" default:"false"`
}

// A property defines a field within a schema. Rules for the property are used
// to enforce validation, and ensures data integrity.
type Property struct {
	Type        string                 `json:"type" validate:"required|in:number,string,list,array,bool,object"`
	Description string                 `json:"description,omitempty"`
	Rules       map[string]interface{} `json:"rules"`
}

// NewSchemaManager creates and returns an initialized SchemaManager that is empty.
func NewSchemaManager() *SchemaManager {
	return &SchemaManager{
		Schemas: make(map[string]*Schema),
	}
}

// LoadSchema loads a schema into the SchemaManager
// NOTE: LoadSchema will overwrite any existing schema with the same type
func (sm *SchemaManager) LoadSchema(schemaDef []byte) error {
	s := &Schema{}
	err := json.Unmarshal(schemaDef, s)
	if err != nil {
		return err
	}

	v := validate.Struct(s)
	if !v.Validate() {
		return fmt.Errorf("schema is invalid: %s", v.Errors.One())
	}

	schemaType := strings.ToLower(s.Type)
	sm.Schemas[schemaType] = s
	return nil
}

// ListSchemas returns a list of all schemaTypes in the SchemaManager
func (sm *SchemaManager) ListSchemas() []string {
	schemaTypes := make([]string, 0, len(sm.Schemas))
	for schemaType := range sm.Schemas {
		schemaTypes = append(schemaTypes, schemaType)
	}
	return schemaTypes
}

// GetSchema returns a schema from the SchemaManager
func (sm *SchemaManager) GetSchema(schemaType string) (*Schema, error) {
	schemaType = strings.ToLower(schemaType)
	if schema, ok := sm.Schemas[schemaType]; ok {
		return schema, nil
	}
	return nil, fmt.Errorf("schema %s not found", schemaType)
}

// DeleteSchema deletes a schema from the SchemaManager
func (sm *SchemaManager) DeleteSchema(schemaType string) error {
	schemaType = strings.ToLower(schemaType)
	if _, ok := sm.Schemas[schemaType]; ok {
		delete(sm.Schemas, schemaType)
		return nil
	}
	return fmt.Errorf("schema %s not found", schemaType)
}

// SchemaCount returns the number of schemas in the SchemaManager
func (sm *SchemaManager) SchemaCount() int {
	return len(sm.Schemas)
}

// ToString returns a string representation of the schema
func (s *Schema) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}

// Validate will validate the provided JSON document against the schema
func (s *Schema) Validate(document []byte) error {

	// first we need to validate the document is valid JSON
	var jsondata interface{}
	err := json.Unmarshal(document, &jsondata)
	if err != nil {
		return err
	}

	// next we need to validate the document against the schema's properties
	for property, p := range s.Properties {

		// TODO: handle optional properties

		// Check if the property exists in the data
		if _, ok := jsondata.(map[string]interface{})[property]; !ok {
			return fmt.Errorf("required property %s is missing", property)
		}

		// if we are not allowing additional properties, then we should check if
		// there are any additional properties in the data
		if !s.AllowUndefinedProperties {
			for key := range jsondata.(map[string]interface{}) {
				if _, ok := s.Properties[key]; !ok {
					return fmt.Errorf("property %s is not defined in the schema", key)
				}
			}
		}

		// Check if the property is the correct type
		value := jsondata.(map[string]interface{})[property]
		if !IsType(value, p.Type) {
			return fmt.Errorf("property %s is not of type %s", property, p.Type)
		}

		// validate value against rules
		for ruleType, ruleArg := range p.Rules {
			err := Evaluate(property, ruleType, ruleArg, value)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
