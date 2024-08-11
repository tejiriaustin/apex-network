package response

import "github.com/gin-gonic/gin"

type Response struct {
	Message string      `json:"message,omitempty"`
	Body    interface{} `json:"body,omitempty"`
}

func FormatResponse(c *gin.Context, statusCode int, message string, body interface{}) {
	format(c, statusCode, message, body)
}

func format(c *gin.Context, statusCode int, message string, body interface{}) {
	var resp Response
	resp.Message = message
	if body != nil {
		resp.Body = body
	}

	c.JSON(statusCode, resp)
}
