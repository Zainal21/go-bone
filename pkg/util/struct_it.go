package util

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

// StructToMap converts a struct to a map using the struct's tags.
// StructToMap uses tags on struct fields to decide which fields to add to the
// returned map.
func StructToMap(src interface{}, tag string) (map[string]interface{}, error) {
	out := map[string]interface{}{}
	v := reflect.ValueOf(src)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		return out, fmt.Errorf("only accepted %s, got %s", reflect.Struct.String(), v.Kind().String())
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)

		//field := reflectValue.Field(i).Interface()
		if !fi.IsExported() {
			continue
		}

		tagsv := strings.Split(fi.Tag.Get(tag), ",")

		if tagsv[0] != "" && fi.PkgPath == "" {

			if isNil(v.Field(i).Interface()) {
				continue
			}

			// skip if omitempty
			if (len(tagsv) > 1 && tagsv[1] == "omitempty") && IsEmptyValue(v.Field(i).Interface()) {
				continue
			}

			if isTime(v.Field(i)) {
				if timeIsZero(v.Field(i)) && (len(tagsv) > 1 && tagsv[1] == "omitempty") {
					continue
				}
			}

			if v.Field(i).Kind() == reflect.Struct {
				continue
			}

			col := tagsv[0]

			if InArray("ne", tagsv) {
				col = fmt.Sprintf("%s !", col)
			}
			// set key value of map interface output
			out[col] = v.Field(i).Interface()
		}

		if tagsv[0] == "" && v.Field(i).Kind() == reflect.Struct {
			x, err := StructToMap(v.Field(i).Interface(), tag)
			if err != nil {
				return out, err
			}

			for y, z := range x {
				out[y] = z
			}
		}
	}

	return out, nil
}

func isNil(i interface{}) bool {
	if i == nil || (reflect.ValueOf(i).Kind() == reflect.Ptr && reflect.ValueOf(i).IsNil()) {
		return true
	}

	return false
}

func isTime(obj reflect.Value) bool {
	_, ok := obj.Interface().(time.Time)
	if ok {
		return ok
	}

	_, ok = obj.Interface().(*time.Time)

	return ok
}

func timeIsZero(obj reflect.Value) bool {
	t, ok := obj.Interface().(time.Time)
	if ok {
		return t.IsZero()
	}

	t2, ok := obj.Interface().(*time.Time)
	if ok {
		return false
	}

	return t2 == nil
}

// ToColumnsValues iterate struct to separate key field and value
func ToColumnsValues(src interface{}, tag string) ([]string, []interface{}, error) {
	var columns []string
	var values []interface{}

	v := reflect.ValueOf(src)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		return nil, nil, fmt.Errorf("only accepted %s, got %s", reflect.Struct.String(), v.Kind().String())
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)

		tagsv := strings.Split(fi.Tag.Get(tag), ",")

		if tagsv[0] != "" && fi.PkgPath == "" {

			if tagsv[0] == "-" {
				continue
			}

			if isNil(v.Field(i).Interface()) {
				continue
			}

			// skip if omitempty
			if (len(tagsv) > 1 && InArray("omitempty", tagsv)) && IsEmptyValue(v.Field(i).Interface()) {
				continue
			}

			if isTime(v.Field(i)) {
				if timeIsZero(v.Field(i)) && (len(tagsv) > 1 && InArray("omitempty", tagsv)) {
					continue
				}
			}

			col := tagsv[0]

			if InArray("ne", tagsv) {
				col = fmt.Sprintf("%s !", col)
			}

			// set value of string slice to value in struct field
			columns = append(columns, col)

			// set value interface of value struct field
			values = append(values, v.Field(i).Interface())

		}
	}

	return columns, values, nil
}
