package utils

func Response(success bool, msg string, data interface{}, err error) map[string]interface{} {
	resMap := map[string]interface{}{
		"success": success,
		"msg":     msg,
		"data":    data,
	}

	if err != nil {
		resMap["error"] = err.Error()
	}

	return resMap
}
