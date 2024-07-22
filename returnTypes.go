package kreturntypes

import (
	"time"

	"github.com/gin-gonic/gin"
)

type ReturnTypes struct {
	Code int
	Message string
	Timestamp time.Time
	Data any
}

func GenResultWithData(target_code int, msg string, data any) ReturnTypes {
	return ReturnTypes{
		Code: target_code,
		Message: msg,
		Timestamp: time.Now(),
		Data: data,
	}
}

func GenResultWithoutData(target_code int, msg string) ReturnTypes {
	return ReturnTypes{
		Code: target_code,
		Message: msg,
		Timestamp: time.Now(),
		Data: nil,
	}
}

func GenGinHWithData(target_code int, msg string, data any) gin.H {
	return gin.H{
		"code": target_code,
		"message": msg,
		"timestamp": time.Now(),
		"data": data,
	}
}

func GenGinHWithoutData(target_code int, msg string) gin.H {
	return gin.H{
		"code": target_code,
		"message": msg,
		"timestamp": time.Now(),
	}
}