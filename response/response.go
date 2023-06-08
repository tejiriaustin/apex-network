package response

type Response struct {
	Message string      `json:"message"`
	Body    interface{} `json:"body"`
}

func FormatResponse(statusCode int, message string, body interface{}) Response {
	return format(statusCode, message, body)
}

func format(statusCode int, message string, body interface{}) Response {
	if statusCode == 200 {
		return Response{
			Message: message,
			Body:    body,
		}
	}
	if body == nil {
		return Response{Message: message}
	}
	return Response{Message: message, Body: body}
}
