package response

import (
	"github.com/gin-gonic/gin"
	"iws/common/ccerr"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//SendResponse 共通返回结果方法
func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := ccerr.DecodeErr(err)
	// always return http.StatusOK
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
