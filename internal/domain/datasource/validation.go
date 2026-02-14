package datasource

import (
	"errors"
	"fmt"
	"strings"
)

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Validate datasource
// - sourceType must be supported
// - configuration params must be valid for the given sourceType
func ValidateInput(sourceType string, configuration map[string]interface{}) ([]FieldError, error) {
	var formDef []map[string]interface{}
	for _, datasource := range SupportedDatasources {
		if datasource["sourceType"] == sourceType {
			formDef = datasource["form"].([]map[string]interface{})
			break
		}
	}

	if formDef == nil {
		return nil, errors.New("Unsupported datasource.")
	}

	var errs []FieldError

	for _, field := range formDef {
		title := field["title"].(string)
		fieldType, _ := field["type"].(string)
		required, _ := field["required"].(bool)

		if fieldType == "const" {
			continue
		}

		value, exists := configuration[title]

		if required && (!exists || value == nil || value == "") {
			errs = append(errs, FieldError{Field: title, Message: fmt.Sprintf("%s is required.", title)})
			continue
		}

		if !exists || value == nil {
			continue
		}

		switch fieldType {
		case "string":
			if _, ok := value.(string); !ok {
				errs = append(errs, FieldError{Field: title, Message: fmt.Sprintf("%s must be a string.", title)})
			}

		case "integer":
			num, ok := toFloat64(value)
			if !ok {
				errs = append(errs, FieldError{Field: title, Message: fmt.Sprintf("%s must be an integer.", title)})
			} else {
				if min, hasMin := field["min"]; hasMin {
					if num < toFloat64Must(min) {
						errs = append(errs, FieldError{Field: title, Message: fmt.Sprintf("%s must be at least %v.", title, min)})
					}
				}
				if max, hasMax := field["max"]; hasMax {
					if num > toFloat64Must(max) {
						errs = append(errs, FieldError{Field: title, Message: fmt.Sprintf("%s must be at most %v.", title, max)})
					}
				}
			}

		case "boolean":
			if _, ok := value.(bool); !ok {
				errs = append(errs, FieldError{Field: title, Message: fmt.Sprintf("%s must be a boolean.", title)})
			}

		case "array":
			arr, ok := value.([]interface{})
			if !ok {
				errs = append(errs, FieldError{Field: title, Message: fmt.Sprintf("%s must be an array.", title)})
			} else if minLen, has := field["minLength"]; has {
				if float64(len(arr)) < toFloat64Must(minLen) {
					errs = append(errs, FieldError{Field: title, Message: fmt.Sprintf("%s must have at least %v items.", title, minLen)})
				}
			}

		case "object":
			obj, ok := value.(map[string]interface{})
			if !ok {
				errs = append(errs, FieldError{Field: title, Message: fmt.Sprintf("%s must be an object.", title)})
			} else if oneOf, has := field["oneOf"]; has {
				errs = append(errs, validateOneOf(title, obj, oneOf.([]map[string]interface{}))...)
			}
		}
	}

	if len(errs) > 0 {
		return errs, nil
	}

	return nil, nil
}

func validateOneOf(parent string, value map[string]interface{}, options []map[string]interface{}) []FieldError {
	var errs []FieldError

	mode, hasMode := value["mode"].(string)
	if !hasMode {
		return []FieldError{{Field: parent, Message: fmt.Sprintf("%s requires a 'mode' selection.", parent)}}
	}

	var matched map[string]interface{}
	for _, opt := range options {
		if opt["value"] == mode {
			matched = opt
			break
		}
	}

	if matched == nil {
		allowed := make([]string, 0, len(options))
		for _, opt := range options {
			allowed = append(allowed, opt["value"].(string))
		}
		return []FieldError{{Field: parent, Message: fmt.Sprintf("%s mode must be one of: %s.", parent, strings.Join(allowed, ", "))}}
	}

	if fields, has := matched["fields"].([]map[string]interface{}); has {
		for _, f := range fields {
			fTitle := f["title"].(string)
			fRequired, _ := f["required"].(bool)
			fVal, fExists := value[fTitle]

			if fRequired && (!fExists || fVal == nil || fVal == "") {
				errs = append(errs, FieldError{
					Field:   parent + "." + fTitle,
					Message: fmt.Sprintf("%s is required when mode is '%s'.", fTitle, mode),
				})
			}
		}
	}

	return errs
}

func toFloat64(v interface{}) (float64, bool) {
	switch n := v.(type) {
	case float64:
		return n, true
	case int:
		return float64(n), true
	case int64:
		return float64(n), true
	default:
		return 0, false
	}
}

func toFloat64Must(v interface{}) float64 {
	f, _ := toFloat64(v)
	return f
}