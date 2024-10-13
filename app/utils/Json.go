package utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	TraceId string `json:"trace_id"`
}

type Response struct {
	Data   interface{} `json:"data"`
	Error  *Error      `json:"error"`
	Status bool        `json:"status"`
}

// JSONResponse responses a http response with json data.
func JSONResponse(c *gin.Context, code int, data interface{}, err error) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(code)

	response := Response{
		Data: data,
	}
	if err != nil {
		e := new(Error)
		e.Message = err.Error()

		e.Code = "InternalServerError"

		response.Error = e
		response.Status = false
	} else {
		response.Status = true
		response.Error = nil
	}

	respBytes, _ := json.Marshal(response)

	c.Writer.Write(respBytes)
	c.Abort()
}
