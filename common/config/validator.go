package config

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	enumspb "go.temporal.io/api/enums/v1"
)

func newValidator() *validator.Validate {
	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterValidation("searchAttributeType", validateSearchAttributeType)
	return validate
}

func validateSearchAttributeType(fl validator.FieldLevel) bool {
	field := fl.Field()
	kind := field.Kind()
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value := int32(field.Int())
		_, ok := enumspb.IndexedValueType_name[value]
		return ok && value > 0

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		value := int32(field.Uint())
		_, ok := enumspb.IndexedValueType_name[value]
		return ok && value > 0

	case reflect.String:
		name := field.String()
		value1, ok1 := enumspb.IndexedValueType_value[name]
		value2, ok2 := enumspb.IndexedValueType_shorthandValue[name]
		return (ok1 && value1 > 0) || (ok2 && value2 > 0)

	default:
		return false
	}
}
