package errcode

import (
	"fmt"
	"github.com/jinzhu/copier"
	"sync"
)

type Err interface {
	ECode() int
	Error() string
	WithDetails(...string) Err
}

var globalErr map[int]Err
var once sync.Once

func NewErr(code int, msg string) Err {
	//唯一初始化全局错误码
	once.Do(func() {
		globalErr = make(map[int]Err)
	},
	)

	if _, ok := globalErr[code]; ok {
		panic("错误码已经存在")
	}
	globalErr[code] = &MyErr{
		Code: code,
		Msg:  msg,
	}
	return globalErr[code]
}

type MyErr struct {
	Code    int      `json:"status_code"`
	Msg     string   `json:"status_msg"`
	Details []string `json:"-"`
}

func (m *MyErr) ECode() int {
	return m.Code
}

func (m *MyErr) Error() string {
	return fmt.Sprintf("错误码：%d  错误信息: %s  详细信息：%v", m.Code, m.Msg, m.Details)
}

func (m *MyErr) WithDetails(details ...string) Err {
	newError := &MyErr{}
	copier.Copy(newError, m)
	newError.Details = append(newError.Details, details...)
	return newError
}
