package jsontype

// var v10 *validator.Validate

// func alpha(value string) error {
// 	errs := v10.Var(value, "required,alpha")
// 	if errs != nil {
// 		return fmt.Errorf("must be a valid alpha string but got %v", value)
// 	}
// 	return nil
// }

// func alphanum(value string) error {
// 	errs := v10.Var(value, "required,alphanum")
// 	if errs != nil {
// 		return fmt.Errorf("must be a valid alphanumeric string but got %v", value)
// 	}
// 	return nil
// }

// func alphadash(value string) error {
// 	errs := v10.Var(value, "required,alphanum")
// 	if errs != nil {
// 		return fmt.Errorf("must be a valid alphanumeric string but got %v", value)
// 	}
// 	return nil
// }

// func email(value string) error {
// 	errs := v10.Var(value, "required,email")
// 	if errs != nil {
// 		return fmt.Errorf("must be a valid email address but got %v", value)
// 	}
// 	return nil
// }

// func b64(value string) error {
// 	errs := v10.Var(value, "required,base64")
// 	if errs != nil {
// 		return fmt.Errorf("must be a valid base64 string but got %v", value)
// 	}
// 	return nil
// }

// func hexcolor(value string) error {
// 	errs := v10.Var(value, "required,hexcolor")
// 	if errs != nil {
// 		return fmt.Errorf("must be a valid hexcolor string but got %v", value)
// 	}
// 	return nil
// }

// func ssn(value string) error {
// 	errs := v10.Var(value, "required,ssn")
// 	if errs != nil {
// 		return fmt.Errorf("must be a valid SSN but got %v", value)
// 	}
// 	return nil
// }

// func uuid(value string) error {
// 	errs := v10.Var(value, "required,uuid")
// 	if errs != nil {
// 		return fmt.Errorf("must be a valid UUID but got %v", value)
// 	}
// 	return nil
// }

// func cidr(value string) error {
// 	errs := v10.Var(value, "required,cidr")
// 	if errs != nil {
// 		return fmt.Errorf("must be a valid CIDR but got %v", value)
// 	}
// 	return nil
// }

// func fqdn(value string) error {
// 	errs := v10.Var(value, "required,fqdn")
// 	if errs != nil {
// 		return fmt.Errorf("must be a valid FQDN but got %v", value)
// 	}
// 	return nil
// }

// func ipv4(value string) error {
// 	errs := v10.Var(value, "required,ipv4")
// 	if errs != nil {
// 		return fmt.Errorf("must be a valid IPV4 Address but got %v", value)
// 	}
// 	return nil
// }

// func ipv6(value string) error {
// 	errs := v10.Var(value, "required,ipv6")
// 	if errs != nil {
// 		return fmt.Errorf("must be a valid IPV6 Address but got %v", value)
// 	}
// 	return nil
// }

// func uri(value string) error {
// 	errs := v10.Var(value, "required,uri")
// 	if errs != nil {
// 		return fmt.Errorf("must be a valid URI but got %v", value)
// 	}
// 	return nil
// }

// func url(value string) error {
// 	errs := v10.Var(value, "required,url")
// 	if errs != nil {
// 		return fmt.Errorf("must be a valid URL but got %v", value)
// 	}
// 	return nil
// }
