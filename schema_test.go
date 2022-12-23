package jsontype_test

import (
	"testing"

	"github.com/goccy/go-reflect"

	"github.com/apageadev/jsontype"
)

func TestCreateSchemaManager(t *testing.T) {
	sm := jsontype.NewSchemaManager()
	if sm == nil {
		t.Fatal("failed to create schema manager")
	}
}

func TestSchemaManagerLoadSchema(t *testing.T) {
	sm := jsontype.NewSchemaManager()
	if sm == nil {
		t.Fatal("failed to create schema manager")
	}

	err := sm.LoadSchema([]byte(`{"type": "Person", "properties": {"name": {"type": "string"}}}`))
	if err != nil {
		t.Fatal(err)
	}

	if sm.SchemaCount() != 1 {
		t.Fatal("failed to load schema")
	}

	// test bad schema
	err = sm.LoadSchema([]byte(`{`))
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestSchemaManagerLoadBadSchema(t *testing.T) {
	sm := jsontype.NewSchemaManager()
	if sm == nil {
		t.Fatal("failed to create schema manager")
	}
	err := sm.LoadSchema([]byte(`{}`))
	t.Log(err)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestSchemaManagerGetSchema(t *testing.T) {
	sm := jsontype.NewSchemaManager()
	if sm == nil {
		t.Fatal("failed to create schema manager")
	}

	err := sm.LoadSchema([]byte(`{"type": "Person", "properties": {"name": {"type": "string"}}}`))
	if err != nil {
		t.Fatal(err)
	}

	schema, err := sm.GetSchema("person")
	if err != nil {
		t.Fatal(err)
	}

	if schema == nil {
		t.Fatal("failed to get schema")
	}

	if schema.Type != "Person" {
		t.Fatal("failed to get schema")
	}

	// test bad schema name
	_, err = sm.GetSchema("bad")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestSchemaManagerDeleteSchema(t *testing.T) {
	sm := jsontype.NewSchemaManager()
	if sm == nil {
		t.Fatal("failed to create schema manager")
	}

	err := sm.LoadSchema([]byte(`{"type": "Person", "properties": {"name": {"type": "string"}}}`))
	if err != nil {
		t.Fatal(err)
	}

	if sm.SchemaCount() != 1 {
		t.Fatal("failed to load schema")
	}

	err = sm.DeleteSchema("person")
	if err != nil {
		t.Fatal(err)
	}

	if sm.SchemaCount() != 0 {
		t.Fatal("failed to delete schema")
	}

	// test bad schema name
	err = sm.DeleteSchema("bad")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestSchemaManagerListSchemas(t *testing.T) {
	sm := jsontype.NewSchemaManager()
	if sm == nil {
		t.Fatal("failed to create schema manager")
	}

	err := sm.LoadSchema([]byte(`{"type": "Person", "properties": {"name": {"type": "string"}}}`))
	if err != nil {
		t.Fatal(err)
	}

	if sm.SchemaCount() != 1 {
		t.Fatal("failed to load schema")
	}

	schemas := sm.ListSchemas()
	if len(schemas) != 1 {
		t.Fatal("failed to list schemas")
	}

	if schemas[0] != "person" {
		t.Fatal("failed to list schemas")
	}
}

func TestSchemaValidate(t *testing.T) {

	// create our schema manager
	sm := jsontype.NewSchemaManager()
	if sm == nil {
		t.Fatal("failed to create schema manager")
	}

	// load our schema into the schema manager
	err := sm.LoadSchema([]byte(`{"type":"Person","properties":{"name":{"type":"string","rules":{"min_length":1,"max_length":100,"format":"alpha"}}}}`))
	if err != nil {
		t.Fatal(err)
	}

	// get our schema from the schema manager
	schema, err := sm.GetSchema("person")
	if err != nil {
		t.Fatal(err)
	}

	// validate our json data against our schema
	err = schema.Validate([]byte(`{"name": "John"}`))
	if err != nil {
		t.Fatal(err)
	}

	// test bad json
	err = schema.Validate([]byte(`{`))
	if err == nil {
		t.Fatal("expected error")
	}

	// test rule violation
	err = schema.Validate([]byte(`{"name": "1a-"}`))
	if err == nil {
		t.Fatal("expected error")
	}

}

func TestSchemaValidateRequiredProperties(t *testing.T) {

	// create our schema manager
	sm := jsontype.NewSchemaManager()
	if sm == nil {
		t.Fatal("failed to create schema manager")
	}

	// load our schema into the schema manager
	err := sm.LoadSchema([]byte(`{"type":"Person","properties":{"name":{"type":"string"},"age":{"type":"number"}}}`))
	if err != nil {
		t.Fatal(err)
	}

	// get our schema from the schema manager
	schema, err := sm.GetSchema("person")
	if err != nil {
		t.Fatal(err)
	}

	// validate our json data against our schema
	err = schema.Validate([]byte(`{"name": "John"}`))
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestSchemaUndefinedProperties(t *testing.T) {

	// create our schema manager
	sm := jsontype.NewSchemaManager()
	if sm == nil {
		t.Fatal("failed to create schema manager")
	}

	// load our schema into the schema manager
	err := sm.LoadSchema([]byte(`{"type":"Person","properties":{"name":{"type":"string"},"age":{"type":"number"}}}`))
	if err != nil {
		t.Fatal(err)
	}

	// get our schema from the schema manager
	schema, err := sm.GetSchema("person")
	if err != nil {
		t.Fatal(err)
	}

	// validate our json data against our schema
	err = schema.Validate([]byte(`{"name": "John", "age": 30, "address": "123 Main St"}`))
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestSchemaValidatePropertyType(t *testing.T) {

	// create our schema manager
	sm := jsontype.NewSchemaManager()
	if sm == nil {
		t.Fatal("failed to create schema manager")
	}

	// load our schema into the schema manager
	err := sm.LoadSchema([]byte(`{"type":"Person","properties":{"name":{"type":"string"},"age":{"type":"number"}}}`))
	if err != nil {
		t.Fatal(err)
	}

	// get our schema from the schema manager
	schema, err := sm.GetSchema("person")
	if err != nil {
		t.Fatal(err)
	}

	// validate our json data against our schema
	err = schema.Validate([]byte(`{"name": "John", "age": "30"}`))
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestSchemaToString(t *testing.T) {

	// create our schema manager
	sm := jsontype.NewSchemaManager()
	if sm == nil {
		t.Fatal("failed to create schema manager")
	}

	// load our schema into the schema manager
	err := sm.LoadSchema([]byte(`{"type":"Person","properties":{"name":{"type":"string"},"age":{"type":"number"}}}`))
	if err != nil {
		t.Fatal(err)
	}

	// get our schema from the schema manager
	schema, err := sm.GetSchema("person")
	if err != nil {
		t.Fatal(err)
	}

	// validate our json data against our schema
	s := schema.String()
	if reflect.TypeOf(s).Kind() != reflect.String {
		t.Fatal("failed to convert schema to string")
	}

	err = sm.LoadSchema([]byte(s))
	if err != nil {
		t.Fatal(err)
	}
}
