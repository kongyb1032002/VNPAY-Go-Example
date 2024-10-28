package mapper

import (
	"errors"
	"reflect"
)

func Map(src interface{}, dst interface{}) error {
	srcVal := reflect.ValueOf(src)
	dstVal := reflect.ValueOf(dst)
	if srcVal.Kind() != reflect.Ptr || dstVal.Kind() != reflect.Ptr {
		return errors.New("src and dst must be pointers")
	}
	srcVal = srcVal.Elem()
	dstVal = dstVal.Elem()

	if srcVal.Kind() != reflect.Struct && srcVal.Kind() != reflect.Slice {
		return errors.New("src must be a struct or a slice")
	}

	if srcVal.Kind() == reflect.Slice {
		if dstVal.Kind() != reflect.Slice {
			return errors.New("dst must be a slice if src is a slice")
		}

		for i := 0; i < srcVal.Len(); i++ {
			srcElem := srcVal.Index(i)
			dstElem := reflect.New(dstVal.Type().Elem()).Elem()
			if err := Map(srcElem.Addr().Interface(), dstElem.Addr().Interface()); err != nil {
				return err
			}
			dstVal.Set(reflect.Append(dstVal, dstElem))
		}
		return nil
	}

	if srcVal.Kind() == reflect.Struct && dstVal.Kind() != reflect.Struct {
		return errors.New("src and dst must be structs")
	}

	srcType := srcVal.Type()
	dstType := dstVal.Type()

	for i := 0; i < srcType.NumField(); i++ {
		srcField := srcType.Field(i)
		srcFieldValue := srcVal.Field(i)

		dstField, found := dstType.FieldByName(srcField.Name)
		if !found {
			continue
		}

		dstFieldValue := dstVal.FieldByName(dstField.Name)
		if srcField.Type != dstField.Type {
			if srcField.Type.Kind() == reflect.Slice && dstField.Type.Kind() == reflect.Slice {
				if err := Map(srcFieldValue.Addr().Interface(), dstFieldValue.Addr().Interface()); err != nil {
					return err
				}
			}
			continue
		}

		dstFieldValue.Set(srcFieldValue)
	}
	return nil
}
