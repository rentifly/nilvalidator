package nilvalidator

import (
	"fmt"
	"reflect"
)

const tagKey = "nilvalidator"

func ValidateStructNotNil(v any) error {
	val := reflect.ValueOf(v)

	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return fmt.Errorf("ValidateStructNotNil: nil pointer")
		}
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return fmt.Errorf("ValidateStructNotNil: expected struct, got %s", val.Kind())
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		if fieldType.Tag.Get(tagKey) != "notnil" {
			continue
		}

		switch field.Kind() {
		case reflect.Interface, reflect.Ptr, reflect.Map, reflect.Slice, reflect.Func, reflect.Chan:
			if field.IsNil() {
				return fmt.Errorf("field '%s' is nil", fieldType.Name)
			}
		default:
			continue
		}
	}

	return nil
}
