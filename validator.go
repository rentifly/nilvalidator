package nilvalidator

import (
	"fmt"
	"reflect"
)

const tagKey = "nilvalidator"

func ValidateStructNotNil(v any) error {
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	if val.Kind() != reflect.Struct {
		return fmt.Errorf("ValidateStructNotNil: expected struct, got %s", val.Kind())
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		if tag := fieldType.Tag.Get(tagKey); tag != "notnil" {
			continue
		}

		kind := field.Kind()
		if (kind == reflect.Interface || kind == reflect.Ptr || kind == reflect.Slice ||
			kind == reflect.Map || kind == reflect.Func || kind == reflect.Chan) && field.IsNil() {
			return fmt.Errorf("field '%s' is nil", fieldType.Name)
		}
	}

	return nil
}
