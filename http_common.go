package http_common

import (
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
	"net/http"
)

type RequestMeta struct {
	RequestID string `json:"requestID"` // 首字母小写，因为swagger生成的API首字母是小写的。
}

type ResponseMeta struct {
	RequestID string `json:"requestID"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
}

func AbortWithError(err error, code int, requestID string, c *gin.Context) {
	klog.Error(err)
	respMeta := ResponseMeta{
		RequestID: requestID,
		Code:      code,
		Message:   err.Error(),
	}
	c.JSON(http.StatusOK, respMeta)
}

type List struct {
}

type ListResponse struct {
	ResponseMetaData *ResponseMeta
	Result           *List
}
