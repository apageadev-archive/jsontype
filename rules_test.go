package jsontype_test

import (
	"testing"

	"github.com/apageadev/jsontype"
)

func TestEvalMin(t *testing.T) {

	// test that 2 is greater than 1
	err := jsontype.Evaluate("fake", "min", 1, 2)
	if err != nil {
		t.Fatal(err)
	}

	// test that 1 is not greater than 2
	err = jsontype.Evaluate("fake", "min", 2, 1)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestEvalMax(t *testing.T) {

	// test that 2 is less than 3
	err := jsontype.Evaluate("fake", "max", 3, 2)
	if err != nil {
		t.Fatal(err)
	}

	// test that 3 is not less than 2
	err = jsontype.Evaluate("fake", "max", 2, 3)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestEvalMinLength(t *testing.T) {

	// test bad arg type
	err := jsontype.Evaluate("fake", "min_length", "abc", "abc")
	if err == nil {
		t.Fatal("expected error")
	}

	// test that "abc" is greater than 2
	err = jsontype.Evaluate("fake", "min_length", 2.0, "abc")
	if err != nil {
		t.Fatal(err)
	}

	err = jsontype.Evaluate("fake", "min_length", 2.0, "a")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestEvalMaxLength(t *testing.T) {
	// test bad arg type
	err := jsontype.Evaluate("fake", "max_length", "abc", "abc")
	if err == nil {
		t.Fatal("expected error")
	}

	// test that "abc" is less than 4
	err = jsontype.Evaluate("fake", "max_length", 4.0, "abc")
	if err != nil {
		t.Fatal(err)
	}

	err = jsontype.Evaluate("fake", "max_length", 4.0, "abcde")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestEvalOneOf(t *testing.T) {
	// test bad arg type
	err := jsontype.Evaluate("fake", "oneof", "abc", "abc")
	if err == nil {
		t.Fatal("expected error")
	}

	// test that "abc" is in the list
	err = jsontype.Evaluate("fake", "oneof", []interface{}{"abc", "def"}, "abc")
	if err != nil {
		t.Fatal(err)
	}

	// test that "ghi" is not in the list
	err = jsontype.Evaluate("fake", "oneof", []interface{}{"abc", "def"}, "ghi")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestEvalAnyOf(t *testing.T) {
	// test bad arg type
	err := jsontype.Evaluate("fake", "anyof", "abc", "abc")
	if err == nil {
		t.Fatal("expected error")
	}

	// test bad value type
	err = jsontype.Evaluate("fake", "noneof", []interface{}{"abc", "def"}, "abc")
	if err == nil {
		t.Fatal("expected error")
	}

	// test that "abc" is in the list
	err = jsontype.Evaluate("fake", "anyof", []interface{}{"abc", "def"}, []interface{}{"abc"})
	if err != nil {
		t.Fatal(err)
	}

	// test that "ghi" is not in the list
	err = jsontype.Evaluate("fake", "anyof", []interface{}{"abc", "def"}, []interface{}{"ghi"})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestEvalAllOf(t *testing.T) {
	// test bad arg type
	err := jsontype.Evaluate("fake", "allof", "abc", "abc")
	if err == nil {
		t.Fatal("expected error")
	}

	// test bad value type
	err = jsontype.Evaluate("fake", "noneof", 123, "abc")
	if err == nil {
		t.Fatal("expected error")
	}

	// test that "abc" is in the list
	err = jsontype.Evaluate("fake", "allof", []interface{}{"abc", "def"}, []interface{}{"abc", "def"})
	if err != nil {
		t.Fatal(err)
	}

	// test that "ghi" is not in the list
	err = jsontype.Evaluate("fake", "allof", []interface{}{"abc", "def"}, []interface{}{"ghi"})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestEvalNoneOf(t *testing.T) {
	// test bad arg type
	err := jsontype.Evaluate("fake", "noneof", "abc", "abc")
	if err == nil {
		t.Fatal("expected error")
	}

	// test bad value type
	err = jsontype.Evaluate("fake", "noneof", []interface{}{"abc", "def"}, "abc")
	if err == nil {
		t.Fatal("expected error")
	}

	// test that "abc" is in the list
	err = jsontype.Evaluate("fake", "noneof", []interface{}{"abc", "def"}, []interface{}{"ghi"})
	if err != nil {
		t.Fatal(err)
	}

	// test that "ghi" is not in the list
	err = jsontype.Evaluate("fake", "noneof", []interface{}{"abc", "def"}, []interface{}{"abc"})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestEvalRegex(t *testing.T) {
	// test bad arg type
	err := jsontype.Evaluate("fake", "regex", 123, "abc")
	if err == nil {
		t.Fatal("expected error")
	}

	// test bad value type
	err = jsontype.Evaluate("fake", "regex", "abc", 123)
	if err == nil {
		t.Fatal("expected error")
	}

	// test that "abc" matches the regex
	err = jsontype.Evaluate("fake", "regex", "abc", "abc")
	if err != nil {
		t.Fatal(err)
	}

	// test that "123" does not match the regex
	err = jsontype.Evaluate("fake", "regex", "abc", "123")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestEvalContaains(t *testing.T) {
	// test that "abc" contains "a"
	err := jsontype.Evaluate("fake", "contains", "a", "abc")
	if err != nil {
		t.Fatal(err)
	}

	// test that "abc" does not contain "d"
	err = jsontype.Evaluate("fake", "contains", "d", "abc")
	if err == nil {
		t.Fatal("expected error")
	}

	// test that slices work
	err = jsontype.Evaluate("fake", "contains", "a", []interface{}{"a", "b", "c"})
	if err != nil {
		t.Fatal(err)
	}

	// test that maps work
	err = jsontype.Evaluate("fake", "contains", "a", map[string]interface{}{"a": 1, "b": 2, "c": 3})
	if err != nil {
		t.Fatal(err)
	}
}

func TestEvalStartsWith(t *testing.T) {

	// test bad arg type
	err := jsontype.Evaluate("fake", "startswith", 123, "abc")
	if err == nil {
		t.Fatal("expected error")
	}

	// test bad value type
	err = jsontype.Evaluate("fake", "startswith", "abc", 123)
	if err == nil {
		t.Fatal("expected error")
	}

	// test that "abc" starts with "a"
	err = jsontype.Evaluate("fake", "startswith", "a", "abc")
	if err != nil {
		t.Fatal(err)
	}

	// test that "abc" does not start with "d"
	err = jsontype.Evaluate("fake", "startswith", "d", "abc")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestEvalFormatAlpha(t *testing.T) {
	// test that "abc" is alpha
	err := jsontype.Evaluate("fake", "format", "alpha", "abc")
	if err != nil {
		t.Fatal(err)
	}

	// test that "123" is not alpha
	err = jsontype.Evaluate("fake", "format", "alpha", "123")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestEvalFormatBadInput(t *testing.T) {
	// test bad arg type
	err := jsontype.Evaluate("fake", "format", 123, "abc")
	if err == nil {
		t.Fatal("expected error")
	}

	// test bad value type
	err = jsontype.Evaluate("fake", "format", "alpha", 123)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestEvalFormatAlphaNum(t *testing.T) {

	// test bad arg type
	err := jsontype.Evaluate("fake", "format", "alphanum", 123)
	if err == nil {
		t.Fatal("expected error")
	}

	// test that "abc123" is alphanum
	err = jsontype.Evaluate("fake", "format", "alphanum", "abc123")
	if err != nil {
		t.Fatal(err)
	}

	// test that "123" is alphanum
	err = jsontype.Evaluate("fake", "format", "alphanum", "123")
	if err != nil {
		t.Fatal(err)
	}

	// test bad value type
	err = jsontype.Evaluate("fake", "format", "alphanum", "abc-123")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestEvalFormatAlphaDash(t *testing.T) {

	// test bad arg type
	err := jsontype.Evaluate("fake", "format", "alphadash", 123)
	if err == nil {
		t.Fatal("expected error")
	}

	// test that "abc123" is alphanum
	err = jsontype.Evaluate("fake", "format", "alphadash", "abc123")
	if err != nil {
		t.Fatal(err)
	}

	// test that "123" is alphanum
	err = jsontype.Evaluate("fake", "format", "alphadash", "123")
	if err != nil {
		t.Fatal(err)
	}

	// test that "abc-123" is alphadash
	err = jsontype.Evaluate("fake", "format", "alphadash", "abc-123")
	if err != nil {
		t.Fatal(err)
	}

	err = jsontype.Evaluate("fake", "format", "alphadash", "abc-123_456")
	if err != nil {
		t.Fatal(err)
	}

	err = jsontype.Evaluate("fake", "format", "alphadash", "abc-123*456[]?{}")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestEvalFormatEmail(t *testing.T) {

	// test bad arg type
	err := jsontype.Evaluate("fake", "format", "email", 123)
	if err == nil {
		t.Fatal("expected error")
	}

	// test that "abc@gmail.com" is email
	err = jsontype.Evaluate("fake", "format", "email", "abc@gmail.com")
	if err != nil {
		t.Fatal(err)
	}

	// test that "abc" is not email
	err = jsontype.Evaluate("fake", "format", "email", "abc")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestEvalFormatBase64(t *testing.T) {

	// test bad arg type
	err := jsontype.Evaluate("fake", "format", "base64", 123)
	if err == nil {
		t.Fatal("expected error")
	}

	// test that "YWJj" is base64 encoded
	err = jsontype.Evaluate("fake", "format", "base64", "YWJj")
	if err != nil {
		t.Fatal(err)
	}

	// test that "abc" is not base64
	err = jsontype.Evaluate("fake", "format", "base64", "abc")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestEvalFormatHexColor(t *testing.T) {

	// test bad arg type
	err := jsontype.Evaluate("fake", "format", "hexcolor", 123)
	if err == nil {
		t.Fatal("expected error")
	}

	// test that "#ffffff" is hexcolor
	err = jsontype.Evaluate("fake", "format", "hexcolor", "#ffffff")
	if err != nil {
		t.Fatal(err)
	}

	// test that "333" is hexcolor
	err = jsontype.Evaluate("fake", "format", "hexcolor", "333")
	if err != nil {
		t.Fatal(err)
	}

	// test invalid hex color
	err = jsontype.Evaluate("fake", "format", "hexcolor", "Z34FF9")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestEvalFormatHexadecimal(t *testing.T) {

	// test bad arg type
	err := jsontype.Evaluate("fake", "format", "hexadecimal", 123)
	if err == nil {
		t.Fatal("expected error")
	}

	// test that F00B42 is hexidecimal
	err = jsontype.Evaluate("fake", "format", "hexadecimal", "F00B42")
	if err != nil {
		t.Fatal(err)
	}

	// test invalid hex
	err = jsontype.Evaluate("fake", "format", "hexadecimal", "Z34FF9")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestBadRule(t *testing.T) {
	err := jsontype.Evaluate("fake", "badrule", "abc", "abc")
	if err == nil {
		t.Fatal("expected error")
	}
}
