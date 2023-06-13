package helper

func FailedResponse(message string) map[string]any {
	return map[string]any{
		"status":  "failed",
		"message": message,
	}
}

func SuccessResponse(message string) map[string]interface{} {
	return map[string]interface{}{
		"status":  "success",
		"message": message,
	}
}

func SuccessWithDataResponse(message string, data any) map[string]any {
	return map[string]any{
		"status":  "success",
		"message": message,
		"data":    data,
	}
}
