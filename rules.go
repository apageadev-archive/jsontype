package jsontype

import (
	"fmt"

	"github.com/goccy/go-reflect"
	"github.com/gookit/validate"
)

func Evaluate(property, ruleType string, ruleArg, value interface{}) error {
	switch ruleType {
	case "min":
		if !validate.Min(value, ruleArg) {
			return fmt.Errorf("%s must be greater than %v but got %v", property, ruleArg, value)
		}
	case "max":
		if !validate.Max(value, ruleArg) {
			return fmt.Errorf("%s must be less than %v but got %v", property, ruleArg, value)
		}
	case "min_length":
		// ensure our ruleArg is an int
		min, ok := ruleArg.(float64)
		if !ok {
			ruleArgType := reflect.TypeOf(ruleArg)
			return fmt.Errorf("min_length rule must be an number but got %v", ruleArgType)
		}
		if !validate.MinLength(value, int(min)) {
			return fmt.Errorf("%s length must be greater than %v but got %v", property, ruleArg, value)
		}

	case "max_length":
		max, ok := ruleArg.(float64)
		if !ok {
			ruleArgType := reflect.TypeOf(ruleArg)
			return fmt.Errorf("max_length rule must be an number but got %v", ruleArgType)
		}
		if !validate.MaxLength(value, int(max)) {
			return fmt.Errorf("%s length must be less than %v but got %v", property, ruleArg, value)
		}
	case "oneof":
		options, ok := ruleArg.([]interface{})
		if !ok {
			return fmt.Errorf("oneof rule must be an array but got %v", ruleArg)
		}
		for _, option := range options {
			if option == value {
				return nil
			}
		}
		return fmt.Errorf("%s must be one of %v but got %v", property, options, value)

	case "noneof":
		options, ok := ruleArg.([]interface{})
		if !ok {
			ruleArgType := reflect.TypeOf(ruleArg)
			return fmt.Errorf("noneof rule must be an array but got %v", ruleArgType)
		}
		values, ok := value.([]interface{})
		if !ok {
			valueType := reflect.TypeOf(value)
			return fmt.Errorf("%s value must be an array but got %v", property, valueType)
		}

		for _, option := range options {
			for _, val := range values {
				if reflect.DeepEqual(option, val) {
					return fmt.Errorf("must not be one of %v but got %v", options, value)
				}
			}
		}
		return nil

	case "allof":
		options, ok := ruleArg.([]interface{})
		if !ok {
			return fmt.Errorf("allof rule must be an array but got %v", ruleArg)
		}
		values, ok := value.([]interface{})
		if !ok {
			return fmt.Errorf("allof value must be an array but got %v", value)
		}
		for _, val := range values {
			found := false
			for _, option := range options {
				if option == val {
					found = true
				}
			}
			if !found {
				return fmt.Errorf("%s must be all of %v but got %v", property, options, value)
			}
		}
		return nil

	case "anyof":
		options, ok := ruleArg.([]interface{})
		if !ok {
			return fmt.Errorf("anyof rule must be an array but got %v", ruleArg)
		}
		values, ok := value.([]interface{})
		if !ok {
			valueType := reflect.TypeOf(value)
			return fmt.Errorf("%s value must be an array but got %v", property, valueType)
		}
		for _, val := range values {
			found := false
			for _, option := range options {
				if option == val {
					found = true
				}
			}
			if found {
				return nil
			}
		}
		return fmt.Errorf("%s must be any of %v but got %v", property, options, value)

	case "regex":
		regex, ok := ruleArg.(string)
		if !ok {
			ruleArgType := reflect.TypeOf(ruleArg)
			return fmt.Errorf("regex rule must be a string but got %v", ruleArgType)
		}

		vstr, ok := value.(string)
		if !ok {
			valueType := reflect.TypeOf(value)
			return fmt.Errorf("%s must be a string but got %v", property, valueType)
		}

		if !validate.Regexp(vstr, regex) {
			return fmt.Errorf("%s must match %v but got %v", property, regex, value)
		}

	case "contains":
		if !validate.Contains(value, ruleArg) {
			return fmt.Errorf("%s must contain %v but got %v", property, ruleArg, value)
		}

	case "startswith":
		substr, ok := ruleArg.(string)
		if !ok {
			return fmt.Errorf("startswith rule must be a string but got %v", ruleArg)
		}

		str, ok := value.(string)
		if !ok {
			return fmt.Errorf("%s must be a string but got %v", property, value)
		}

		if !validate.StartsWith(str, substr) {
			return fmt.Errorf("%s must start with %v but got %v", property, substr, value)
		}

	// TODO: should we support multiple formats for a single property?
	case "format":
		format, ok := ruleArg.(string)
		if !ok {
			return fmt.Errorf("format rule must be a string but got %v", ruleArg)
		}
		switch format {
		case "alpha":
			v, ok := value.(string)
			if !ok {
				return fmt.Errorf("%s must be a string but got %v", property, value)
			}
			if !validate.IsAlpha(v) {
				return fmt.Errorf("%s must be alpha but got %v", property, value)
			}

		case "alphanum":
			v, ok := value.(string)
			if !ok {
				return fmt.Errorf("%s must be a string but got %v", property, value)
			}
			if !validate.IsAlphaNum(v) {
				return fmt.Errorf("%s must be alpha numeric but got %v", property, value)
			}

		case "alphadash":
			v, ok := value.(string)
			if !ok {
				return fmt.Errorf("%s must be a string but got %v", property, value)
			}
			if !validate.IsAlphaDash(v) {
				return fmt.Errorf("%s must be alpha dash but got %v", property, value)
			}

		case "email":
			v, ok := value.(string)
			if !ok {
				return fmt.Errorf("%s must be a string but got %v", property, value)
			}
			if !validate.IsEmail(v) {
				return fmt.Errorf("%s must be email but got %v", property, value)
			}

		case "base64":
			v, ok := value.(string)
			if !ok {
				return fmt.Errorf("%s must be a string but got %v", property, value)
			}
			if !validate.IsBase64(v) {
				return fmt.Errorf("%s must be base64 but got %v", property, value)
			}

		case "hexcolor":
			v, ok := value.(string)
			if !ok {
				return fmt.Errorf("%s must be a string but got %v", property, value)
			}
			if !validate.IsHexColor(v) {
				return fmt.Errorf("%s must be hex color but got %v", property, value)
			}

		case "hexadecimal":
			v, ok := value.(string)
			if !ok {
				return fmt.Errorf("%s must be a string but got %v", property, value)
			}
			if !validate.IsHexadecimal(v) {
				return fmt.Errorf("%s must be hexadecimal but got %v", property, value)
			}

		case "json":
			v, ok := value.(string)
			if !ok {
				return fmt.Errorf("%s must be a string but got %v", property, value)
			}
			if !validate.IsJSON(v) {
				return fmt.Errorf("%s must be JSON but got %v", property, value)
			}

		case "rgbcolor":
			v, ok := value.(string)
			if !ok {
				return fmt.Errorf("%s must be a string but got %v", property, value)
			}
			if !validate.IsRGBColor(v) {
				return fmt.Errorf("%s must be RGB color but got %v", property, value)
			}

		case "url":
			v, ok := value.(string)
			if !ok {
				return fmt.Errorf("%s must be a string but got %v", property, value)
			}
			if !validate.IsURL(v) {
				return fmt.Errorf("%s must be URL but got %v", property, value)
			}

		case "fullurl":
			v, ok := value.(string)
			if !ok {
				return fmt.Errorf("%s must be a string but got %v", property, value)
			}
			if !validate.IsFullURL(v) {
				return fmt.Errorf("%s must be URL but got %v", property, value)
			}

		case "ip":
			v, ok := value.(string)
			if !ok {
				return fmt.Errorf("%s must be a string but got %v", property, value)
			}
			if !validate.IsIP(v) {
				return fmt.Errorf("%s must be IP but got %v", property, value)
			}

		case "ipv4":
			v, ok := value.(string)
			if !ok {
				return fmt.Errorf("%s must be a string but got %v", property, value)
			}
			if !validate.IsIPv4(v) {
				return fmt.Errorf("%s must be IPv4 but got %v", property, value)
			}

		case "ipv6":
			v, ok := value.(string)
			if !ok {
				return fmt.Errorf("%s must be a string but got %v", property, value)
			}
			if !validate.IsIPv6(v) {
				return fmt.Errorf("%s must be IPv6 but got %v", property, value)
			}

		case "cidr":
			v, ok := value.(string)
			if !ok {
				return fmt.Errorf("%s must be a string but got %v", property, value)
			}
			if !validate.IsCIDR(v) {
				return fmt.Errorf("%s must be CIDR but got %v", property, value)
			}

		case "cidrv4":
			v, ok := value.(string)
			if !ok {
				return fmt.Errorf("%s must be a string but got %v", property, value)
			}
			if !validate.IsCIDRv4(v) {
				return fmt.Errorf("%s must be CIDRv4 but got %v", property, value)
			}

		case "cidrv6":
			v, ok := value.(string)
			if !ok {
				return fmt.Errorf("%s must be a string but got %v", property, value)
			}
			if !validate.IsCIDRv6(v) {
				return fmt.Errorf("%s must be CIDRv6 but got %v", property, value)
			}

		case "uuid":
			v, ok := value.(string)
			if !ok {
				return fmt.Errorf("%s must be a string but got %v", property, value)
			}
			if !validate.IsUUID(v) {
				return fmt.Errorf("%s must be UUID but got %v", property, value)
			}

		case "filepath":
			v, ok := value.(string)
			if !ok {
				return fmt.Errorf("%s must be a string but got %v", property, value)
			}
			if !validate.IsFilePath(v) {
				return fmt.Errorf("%s must be file path but got %v", property, value)
			}
		}

	default:
		return fmt.Errorf("unknown rule %s", ruleType)
	}
	return nil
}
