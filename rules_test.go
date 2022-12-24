package jsontype_test

import (
	"io/ioutil"
	"os"
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

func TestEvalFormatJSON(t *testing.T) {

	// test bad arg type
	err := jsontype.Evaluate("fake", "format", "json", 123)
	if err == nil {
		t.Fatal("expected error")
	}

	// test that "{}" is json
	err = jsontype.Evaluate("fake", "format", "json", "{}")
	if err != nil {
		t.Fatal(err)
	}

	// test that "{}" is json
	err = jsontype.Evaluate("fake", "format", "json", `{"a":1,"b":2,"c":3}`)
	if err != nil {
		t.Fatal(err)
	}

	// test that "abc" is not json
	err = jsontype.Evaluate("fake", "format", "json", "abc")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestEvalFormatRGBColor(t *testing.T) {

	// test bad arg type
	err := jsontype.Evaluate("fake", "format", "rgbcolor", 123)
	if err == nil {
		t.Fatal("expected error")
	}

	// test that "rgb(255,255,255)" is rgbcolor
	err = jsontype.Evaluate("fake", "format", "rgbcolor", "rgb(255,255,255)")
	if err != nil {
		t.Fatal(err)
	}

	// test that "rgb(255, 255, 255)" is rgbcolor
	err = jsontype.Evaluate("fake", "format", "rgbcolor", "rgb(255, 255, 255)")
	if err != nil {
		t.Fatal(err)
	}
}

func TestEvalFormatURL(t *testing.T) {

	// test bad arg type
	err := jsontype.Evaluate("fake", "format", "url", 123)
	if err == nil {
		t.Fatal("expected error")
	}

	// test that "http://example.com" is url
	err = jsontype.Evaluate("fake", "format", "url", "http://example.com")
	if err != nil {
		t.Fatal(err)
	}

	// test that "example.com" is a url
	err = jsontype.Evaluate("fake", "format", "url", "example.com")
	if err != nil {
		t.Fatal(err)
	}
}

func TestEvalFormatFullURL(t *testing.T) {

	// test bad arg type
	err := jsontype.Evaluate("fake", "format", "fullurl", 123)
	if err == nil {
		t.Fatal("expected error")
	}

	// test that "http://example.com" is fullurl
	err = jsontype.Evaluate("fake", "format", "fullurl", "http://example.com")
	if err != nil {
		t.Fatal(err)
	}

	// test that "example.com" is not a fullurl
	err = jsontype.Evaluate("fake", "format", "fullurl", "example.com")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestEvalFormatIP(t *testing.T) {

	// test bad arg type
	err := jsontype.Evaluate("fake", "format", "ip", 123)
	if err == nil {
		t.Fatal("expected error")
	}

	// test that "192.168.1.1" is ip
	err = jsontype.Evaluate("fake", "format", "ip", "192.168.1.1")
	if err != nil {
		t.Fatal(err)
	}

	// test that "0.1.2" is not ip
	err = jsontype.Evaluate("fake", "format", "ip", "0.1.2")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestEvalFormatIPv4(t *testing.T) {

	// test bad arg type
	err := jsontype.Evaluate("fake", "format", "ipv4", 123)
	if err == nil {
		t.Fatal("expected error")
	}

	// test that "192.168.1.123" is ipv4
	err = jsontype.Evaluate("fake", "format", "ipv4", "192.168.1.123")
	if err != nil {
		t.Fatal(err)
	}

	// test that "0.1.2" is not ipv4
	err = jsontype.Evaluate("fake", "format", "ipv4", "0.1.2")
	if err == nil {
		t.Fatal("expected error")
	}

	// test that "2001:db8:85a3:8d3:1319:8a2e:370:7348" is not ipv4
	err = jsontype.Evaluate("fake", "format", "ipv4", "2001:db8:85a3:8d3:1319:8a2e:370:7348")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestEvalFormatIPv6(t *testing.T) {

	// test bad arg type
	err := jsontype.Evaluate("fake", "format", "ipv6", 123)
	if err == nil {
		t.Fatal("expected error")
	}

	// test that "2001:db8:85a3:8d3:1319:8a2e:370:7348" is ipv6
	err = jsontype.Evaluate("fake", "format", "ipv6", "2001:db8:85a3:8d3:1319:8a2e:370:7348")
	if err != nil {
		t.Fatal(err)
	}

	// test that "0.1.2" is not ipv6
	err = jsontype.Evaluate("fake", "format", "ipv6", "0.1.2")
	if err == nil {
		t.Fatal("expected error")
	}

	// test that "192.168.1.2" is not ipv6
	err = jsontype.Evaluate("fake", "format", "ipv6", "192.168.1.2")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestEvalFormatCIDR(t *testing.T) {

	// test bad arg type
	err := jsontype.Evaluate("fake", "format", "cidr", 123)
	if err == nil {
		t.Fatal("expected error")
	}

	// test that "192.168. 129.23/17" is cidr
	err = jsontype.Evaluate("fake", "format", "cidr", "192.168.129.23/17")
	if err != nil {
		t.Fatal(err)
	}

	// test that "0.1.2" is not cidr
	err = jsontype.Evaluate("fake", "format", "cidr", "0.1.2")
	if err == nil {
		t.Fatal("expected error")
	}

	// test that "2002::1234:abcd:ffff:c0a8:101/64" is cidr
	err = jsontype.Evaluate("fake", "format", "cidr", "2002::1234:abcd:ffff:c0a8:101/64")
	if err != nil {
		t.Fatal(err)
	}

	// test that "2002::1234:abcd:ffff:c0a8:101" is not cidr
	err = jsontype.Evaluate("fake", "format", "cidr", "2002::1234:abcd:ffff:c0a8:101")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestEvalFormatCIDRV4(t *testing.T) {

	// test bad arg type
	err := jsontype.Evaluate("fake", "format", "cidrv4", 123)
	if err == nil {
		t.Fatal("expected error")
	}

	// test that "192.168. 129.23/17" is cidrv4
	err = jsontype.Evaluate("fake", "format", "cidrv4", "192.168.129.23/17")
	if err != nil {
		t.Fatal(err)
	}

	// test that "0.1.2" is not cidrv4
	err = jsontype.Evaluate("fake", "format", "cidrv4", "0.1.2")
	if err == nil {
		t.Fatal("expected error")
	}

	// test that "2002::1234:abcd:ffff:c0a8:101/64" is not cidrv4
	err = jsontype.Evaluate("fake", "format", "cidrv4", "2002::1234:abcd:ffff:c0a8:101/64")
	if err == nil {
		t.Fatal("expected error")
	}

	// test that "2002::1234:abcd:ffff:c0a8:101" is not cidrv4
	err = jsontype.Evaluate("fake", "format", "cidrv4", "2002::1234:abcd:ffff:c0a8:101")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestEvalFormatCIDRV6(t *testing.T) {

	// test bad arg type
	err := jsontype.Evaluate("fake", "format", "cidrv6", 123)
	if err == nil {
		t.Fatal("expected error")
	}

	// test that "192.168.129.23/17" is not cidrv6
	err = jsontype.Evaluate("fake", "format", "cidrv6", "192.168.129.23/17")
	if err == nil {
		t.Fatal("expected error")
	}

	// test that "0.1.2" is not cidrv6
	err = jsontype.Evaluate("fake", "format", "cidrv6", "0.1.2")
	if err == nil {
		t.Fatal("expected error")
	}

	// test that "2002::1234:abcd:ffff:c0a8:101/64" is cidrv6
	err = jsontype.Evaluate("fake", "format", "cidrv6", "2002::1234:abcd:ffff:c0a8:101/64")
	if err != nil {
		t.Fatal(err)
	}

	// test that "2002::1234:abcd:ffff:c0a8:101" is not cidrv6
	err = jsontype.Evaluate("fake", "format", "cidrv6", "2002::1234:abcd:ffff:c0a8:101")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestEvalFormatUUID(t *testing.T) {

	// test bad arg type
	err := jsontype.Evaluate("fake", "format", "uuid", 123)
	if err == nil {
		t.Fatal("expected error")
	}

	// test that "f81d4fae-7dec-11d0-a765-00a0c91e6bf6" is uuid
	err = jsontype.Evaluate("fake", "format", "uuid", "f81d4fae-7dec-11d0-a765-00a0c91e6bf6")
	if err != nil {
		t.Fatal(err)
	}

	// test that "0.1.2" is not uuid
	err = jsontype.Evaluate("fake", "format", "uuid", "0.1.2")
	if err == nil {
		t.Fatal("expected error")
	}

	// test that "2002::1234:abcd:ffff:c0a8:101/64" is not uuid
	err = jsontype.Evaluate("fake", "format", "uuid", "2002::1234:abcd:ffff:c0a8:101/64")
	if err == nil {
		t.Fatal("expected error")
	}

	// test uuid v3
	err = jsontype.Evaluate("fake", "format", "uuid", "a3bb189e-8bf9-3888-9912-ace4e6543002")
	if err != nil {
		t.Fatal(err)
	}

	// test uuid v4
	err = jsontype.Evaluate("fake", "format", "uuid", "c6b932b1-fb82-403c-9c8b-ae1816291648")
	if err != nil {
		t.Fatal(err)
	}

	// test uuid v5
	err = jsontype.Evaluate("fake", "format", "uuid", "a6edc906-2f9f-5fb2-a373-efac406f0ef2")
	if err != nil {
		t.Fatal(err)
	}
}

func TestEvalFormatFilePath(t *testing.T) {

	// test bad arg type
	err := jsontype.Evaluate("fake", "format", "filepath", 123)
	if err == nil {
		t.Fatal("expected error")
	}

	// create a file to test
	err = ioutil.WriteFile("/tmp/test.txt", []byte("test"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// test that "/tmp/test" is filepath
	err = jsontype.Evaluate("fake", "format", "filepath", "/tmp/test.txt")
	if err != nil {
		t.Fatal(err)
	}

	// test that "0.1.2" is not filepath
	err = jsontype.Evaluate("fake", "format", "filepath", "0.1.2")
	if err == nil {
		t.Fatal("expected error")
	}

	// test that "2002::1234:abcd:ffff:c0a8:101/64" is not filepath
	err = jsontype.Evaluate("fake", "format", "filepath", "2002::1234:abcd:ffff:c0a8:101/64")
	if err == nil {
		t.Fatal("expected error")
	}

	// test that "/tmp/test" is not a filepath
	err = jsontype.Evaluate("fake", "format", "filepath", "/tmp/test")
	if err == nil {
		t.Fatal("expected error")
	}

	// delete the test file
	err = os.Remove("/tmp/test.txt")
	if err != nil {
		t.Fatal(err)
	}
}

func TestEvalFormatISBN10(t *testing.T) {

	// test bad arg type
	err := jsontype.Evaluate("fake", "format", "isbn10", 123)
	if err == nil {
		t.Fatal("expected error")
	}

	// test that "0-545-01022-5" is isbn10
	err = jsontype.Evaluate("fake", "format", "isbn10", "0-545-01022-5")
	if err != nil {
		t.Fatal(err)
	}

	// test that "0.1.2" is not isbn10
	err = jsontype.Evaluate("fake", "format", "isbn10", "0.1.2")
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
