package jsontype_test

import (
	"testing"

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
	err := sm.LoadSchema([]byte(`{"type": "Person", "properties": {"name": {"type": "string"}}}`))
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
}
