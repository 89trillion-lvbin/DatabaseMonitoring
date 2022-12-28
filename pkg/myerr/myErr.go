package myerr

import (
	"fmt"
	"runtime/debug"
)

// Api handle status
type MyErr struct {
	Code      int
	Message   string
	OriginMsg string
	Level     Level
	Stack     []byte
}

type Level uint32

const (
	ErrorLevel Level = iota
	WarnLevel
	InfoLevel
	DebugLevel
	SuccessLevel
)

// myErr是否相等
func (e *MyErr) Is(err error) bool {
	if myErr, ok := err.(*MyErr); ok {
		return e.Code == myErr.Code
	}
	return false
}

// 附加stack信息
func (e *MyErr) WithStack() *MyErr {
	err := *e
	err.Stack = debug.Stack()
	return &err
}

// 记录原始错误
func (e *MyErr) WithStackInfo(err error) *MyErr {
	return &MyErr{
		Code:      e.Code,
		Message:   e.Message,
		OriginMsg: err.Error(),
		Level:     e.Level,
		Stack:     debug.Stack(),
	}
}

// get error information
func (e *MyErr) Error() string {
	return e.Message
}

func (e *MyErr) Detail() string {
	msg := e.Message
	if len(e.OriginMsg) > 0 {
		msg += e.OriginMsg
	}
	if len(e.Stack) > 0 {
		msg += fmt.Sprintf("\n%s", string(e.Stack))
	}
	return msg
}

var (
	SUCCESS        = &MyErr{Code: 200, Message: "OK", Level: SuccessLevel}
	LACK_OF_CONFIG = &MyErr{Code: 2603, Message: "lack of config", Level: InfoLevel}
)
