package jsontype_test

import (
	"encoding/json"
	"testing"

	"github.com/apageadev/jsontype"
)

func TestStringType(t *testing.T) {
	if !jsontype.IsString("hello123") {
		t.Fatal("expected string")
	}
}

func TestNumberType(t *testing.T) {
	if !jsontype.IsNumber(1) {
		t.Fatal("expected number")
	}

	if !jsontype.IsNumber(1.0) {
		t.Fatal("expected number")
	}

	if !jsontype.IsNumber(1.0e10) {
		t.Fatal("expected number")
	}

	if !jsontype.IsNumber(1.0e-10) {
		t.Fatal("expected number")
	}
}

func TestBoolType(t *testing.T) {
	if !jsontype.IsBool(true) {
		t.Fatal("expected bool")
	}

	if !jsontype.IsBool(false) {
		t.Fatal("expected bool")
	}
}

func TestNullType(t *testing.T) {
	if !jsontype.IsNull(nil) {
		t.Fatal("expected null")
	}

	data := []byte(`{"field": null}`)
	var obj map[string]interface{}
	err := json.Unmarshal(data, &obj)
	if err != nil {
		t.Fatal(err)
	}

	if !jsontype.IsNull(obj["field"]) {
		t.Fatal("expected null")
	}
}

func TestObjectType(t *testing.T) {
	data := []byte(`{"field": "value"}`)
	var obj map[string]interface{}
	err := json.Unmarshal(data, &obj)
	if err != nil {
		t.Fatal(err)
	}

	if !jsontype.IsObject(obj) {
		t.Fatal("expected object")
	}
}

func TestArrayType(t *testing.T) {
	data := []byte(`[1,2,3]`)
	var arr []interface{}
	err := json.Unmarshal(data, &arr)
	if err != nil {
		t.Fatal(err)
	}

	if !jsontype.IsArray(arr) {
		t.Fatal("expected array")
	}

	data = []byte(`[1,2,"3"]`)
	err = json.Unmarshal(data, &arr)
	if err != nil {
		t.Fatal(err)
	}

	if jsontype.IsArray(arr) {
		t.Fatal("expected list")
	}

	// test non-slice type
	if jsontype.IsArray(1) {
		t.Fatal("expected list")
	}
}

func TestListType(t *testing.T) {
	data := []byte(`[1,2,3]`)
	var arr []interface{}
	err := json.Unmarshal(data, &arr)
	if err != nil {
		t.Fatal(err)
	}

	if !jsontype.IsList(arr) {
		t.Fatal("expected list")
	}

	data = []byte(`[1,2,"3"]`)
	err = json.Unmarshal(data, &arr)
	if err != nil {
		t.Fatal(err)
	}

	if !jsontype.IsList(arr) {
		t.Fatal("expected list")
	}
}

func TestListTypeWithObject(t *testing.T) {
	data := []byte(`[1,2,{"field": "value"}]`)
	var arr []interface{}
	err := json.Unmarshal(data, &arr)
	if err != nil {
		t.Fatal(err)
	}

	if !jsontype.IsList(arr) {
		t.Fatal("expected list")
	}
}

func TestIsType(t *testing.T) {
	str := "hello"
	if !jsontype.IsType(str, "string") {
		t.Fatal("expected string")
	}

	num := 1
	if !jsontype.IsType(num, "number") {
		t.Fatal("expected number")
	}

	b := true
	if !jsontype.IsType(b, "bool") {
		t.Fatal("expected bool")
	}

	if !jsontype.IsType(nil, "null") {
		t.Fatal("expected null")
	}

	data := []byte(`{"field": "value"}`)
	var obj map[string]interface{}
	err := json.Unmarshal(data, &obj)
	if err != nil {
		t.Fatal(err)
	}

	if !jsontype.IsType(obj, "object") {
		t.Fatal("expected object")
	}

	data = []byte(`[1,2,3]`)
	var arr []interface{}
	err = json.Unmarshal(data, &arr)
	if err != nil {
		t.Fatal(err)
	}

	if !jsontype.IsType(arr, "array") {
		t.Fatal("expected array")
	}

	data = []byte(`[1,2,"3"]`)
	err = json.Unmarshal(data, &arr)
	if err != nil {
		t.Fatal(err)
	}

	if !jsontype.IsType(arr, "list") {
		t.Fatal("expected list")
	}

	data = []byte(`[1,2,{"field": "value"}]`)
	err = json.Unmarshal(data, &arr)
	if err != nil {
		t.Fatal(err)
	}

	if !jsontype.IsType(arr, "list") {
		t.Fatal("expected list")
	}

	if jsontype.IsType(arr, "array") {
		t.Fatal("expected list")
	}

	if jsontype.IsType(arr, "object") {
		t.Fatal("expected list")
	}

	if jsontype.IsType(nil, "badtype") != false {
		t.Fatal("expected false")
	}
}
