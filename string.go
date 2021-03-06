package jsonvalidator

import (
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/tidwall/gjson"
)

//IsString Checks if a given value is a string
func IsString(value *gjson.Result) bool {
	_, ok := value.Value().(string)
	return ok
}

//CollapseWhiteSpaces Remove all preceeding whitespaces from a given string
func CollapseWhiteSpaces(value string) string {
	strs := make([]string, 0)

	for _, s := range strings.Split(value, " ") {
		tmpStr := strings.TrimSpace(s)

		if 0 != len(tmpStr) {
			strs = append(strs, tmpStr)
		}
	}
	return strings.Join(strs, " ")
}

//TypeString Returns a new instance of String validation constraint
func TypeString(message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if IsEmpty(value) || IsString(value) {
			return
		}

		violations.Add(field, message)
	})
}

//MinLength Creates a new MinLength validation constraint
func MinLength(length int, message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if IsEmpty(value) || !IsString(value) {
			return
		}

		if len(value.String()) < length {
			violations.Add(field, message)
		}
	})
}

//MaxLength Creates a new MaxLength validation constraint
func MaxLength(length int, message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if IsEmpty(value) || !IsString(value) {
			return
		}

		if len(value.String()) > length {
			violations.Add(field, message)
		}
	})
}

//Length Creates a constraint for validating length of a string
func Length(length int, message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if IsEmpty(value) || !IsString(value) {
			return
		}

		if len(value.String()) != length {
			violations.Add(field, message)
		}
	})
}

//IPV4 Creates a constraint for validating IP address
func IPV4(message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if IsEmpty(value) {
			return
		}

		if !IsString(value) || !govalidator.IsIPv4(value.String()) {
			violations.Add(field, message)
		}
	})
}

//IPV6 Creates a constraint for validating IP address
func IPV6(message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if IsEmpty(value) {
			return
		}

		if !IsString(value) || !govalidator.IsIPv6(value.String()) {
			violations.Add(field, message)
		}
	})
}

//Latitude Creates a constraint for validating latitudes
func Latitude(message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if IsEmpty(value) {
			return
		}

		if !IsString(value) || !govalidator.IsLatitude(value.String()) {
			violations.Add(field, message)
		}
	})
}

//Longitude Creates a constraint for validating Longitude
func Longitude(message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if IsEmpty(value) {
			return
		}

		if !IsString(value) || !govalidator.IsLongitude(value.String()) {
			violations.Add(field, message)
		}
	})
}

//Port Creates a constraint for validating Longitude
func Port(message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if IsEmpty(value) {
			return
		}

		if !IsString(value) || !govalidator.IsPort(value.String()) {
			violations.Add(field, message)
		}
	})
}

//Alpha Creates a constraint for validating alpha characters
func Alpha(message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if IsEmpty(value) {
			return
		}

		if !IsString(value) || !govalidator.IsAlpha(value.String()) {
			violations.Add(field, message)
		}
	})
}

//AlphaNumberic Creates a constraint for validating alpha numeric characters
func AlphaNumberic(message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if IsEmpty(value) {
			return
		}

		if !IsString(value) || !govalidator.IsAlphanumeric(value.String()) {
			violations.Add(field, message)
		}
	})
}

//Lowercase Creates a constraint for validating lowercase characters
func Lowercase(message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if IsEmpty(value) {
			return
		}

		if !IsString(value) || !govalidator.IsLowerCase(value.String()) {
			violations.Add(field, message)
		}
	})
}

//Uppercase Creates a constraint for validating uppercase characters
func Uppercase(message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if IsEmpty(value) {
			return
		}

		if !IsString(value) || !govalidator.IsUpperCase(value.String()) {
			violations.Add(field, message)
		}
	})
}

//ASCII Creates a constraint for validating ASCII characters
func ASCII(message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if IsEmpty(value) {
			return
		}

		if !IsString(value) || !govalidator.IsASCII(value.String()) {
			violations.Add(field, message)
		}
	})
}

//Phone Creates a constraint for validating phone numbers
func Phone(message string) *Rule {
	return Pattern(`^[+]?([\d]{0,3})?[\(\.\-\s]?(([\d]{1,3})[\)\.\-\s]*)?(([\d]{3,5})[\.\-\s]?([\d]{4})|([\d]{2}[\.\-\s]?){4})$`, message)
}

//URL Creates a constraint for validating URL
func URL(message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if IsEmpty(value) {
			return
		}

		if !IsString(value) || !govalidator.IsURL(value.String()) {
			violations.Add(field, message)
		}
	})
}

//Subdomain Creates a constraint for validating Subdomain
func Subdomain(message string) *Rule {
	return Pattern(`^[A-Za-z0-9](?:[A-Za-z0-9\-.]{0,61}[A-Za-z0-9])?$`, message)
}

//Pattern Creates a constraint for validating a string againt a given pattern
func Pattern(pattern, message string) *Rule {
	return NewRule(func(field string, value *gjson.Result, parent *gjson.Result, source *gjson.Result, violations *Violations, validator *Validator) {
		if IsEmpty(value) {
			return
		}

		if !IsString(value) || !govalidator.Matches(value.String(), pattern) {
			violations.Add(field, message)
		}
	})
}
