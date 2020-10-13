package success

import (
	"reflect"
)

func Response(data interface{}) map[string]interface{} {

	if reflect.DeepEqual(data, reflect.Zero(reflect.TypeOf(data)).Interface()) {
		data = make(map[string]string)
	}

	return map[string]interface{}{
		"success": true,
		"data":    data,
	}
}
