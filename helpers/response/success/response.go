package success

import (
	"reflect"
)

func Response(data interface{}, extra interface{}) map[string]interface{} {

	if reflect.DeepEqual(data, reflect.Zero(reflect.TypeOf(data)).Interface()) {
		data = make(map[string]string)
	}
	if reflect.DeepEqual(extra, reflect.Zero(reflect.TypeOf(extra)).Interface()) {
		extra = make(map[string]string)
	}

	return map[string]interface{}{
		"success": true,
		"data":    data,
		"extra":   extra,
	}
}
