package error

func NewErrorResponse(code int, title string, message string, extra interface{}) map[string]interface{} {

	return map[string]interface{}{
		"success": false,
		"data": map[string]interface{}{
			"code":    code,
			"title":   title,
			"message": message,
			"extra":   extra,
		},
	}
}
