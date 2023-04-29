package stringer

import (
	"encoding/json"
	"reflect"
)

func ToString(object any) string {
	b, err := json.Marshal(object)
	if err != nil {
		return err.Error()
	}

	return string(b)
}

func TypeOf(object any) string {
	return reflect.TypeOf(object).String()
}
