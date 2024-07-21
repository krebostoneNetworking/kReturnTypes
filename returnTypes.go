package kreturntypes

import (
	"time"
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