package gate

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

var (
	// errInvalidHandlerType : handler 不是 func
	errInvalidHandlerType = errors.New("handler should be function type")
	// errInvalidInputParamsNum : handler 函数入参数数量错误
	errInvalidInputParamsNum = errors.New("handler function require 1 or 2 input parameters")
	// errInvalidOutputParamsNum : handler 函数返回参数数数量错误
	errInvalidOutputParamsNum = errors.New("handler function require 1 output parameters")
)

func Wrap(f interface{}) gin.HandlerFunc {
	t := reflect.TypeOf(f)

	// 检查 handler 的 Type
	if t.Kind() != reflect.Func {
		panic(errInvalidHandlerType)
	}

	// 检查 handler 的 input num
	numIn := t.NumIn()
	if numIn < 1 || numIn > 2 {
		panic(errInvalidInputParamsNum)
	}

	// 检查 handler 的 output num
	numOut := t.NumOut()
	if numOut != 1 {
		panic(errInvalidOutputParamsNum)
	}

	return func(c *gin.Context) {
		// handler 入参
		// 第一个参数为 context.Context, 传递上下文
		inValues := []reflect.Value{
			reflect.ValueOf(c),
		}

		// 如果需要
		if numIn == 2 {
			req := newReqInstance(t.In(1))
			if err := c.Bind(req); err != nil {
				fmt.Println(err)
				// TODO: err handle
			}

			inValues = append(inValues, reflect.ValueOf(req))
		}

		ret := reflect.ValueOf(f).Call(inValues)
		c.JSON(http.StatusOK, ret[0].Interface())
	}
}

func newReqInstance(t reflect.Type) interface{} {
	switch t.Kind() {
	case reflect.Ptr, reflect.Interface:
		return newReqInstance(t.Elem())
	default:
		return reflect.New(t).Interface()
	}
}
