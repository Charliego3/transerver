package utils

import (
	"fmt"
	"reflect"
)

type container struct {
	objs map[string]any
}

var c = &container{
	make(map[string]any),
}

func Singleton[T any](t T) {
	name := fmt.Sprintf("%T", t)
	c.objs[name] = t
}

func SingletonFn[T any](fn func() (T, error)) {

}

func Obj[T any]() (T, error) {
	var t T
	// name := fmt.Sprintf("%T", t)
	// if o, ok := c.objs[name]; ok {
	// 	switch obj := o.(type) {
	// 	case *T:
	// 		// *t = *obj
	// 	}
	// }
	return t, nil
}

func Call[T any](fn func(T) error) {
	tf := reflect.TypeOf(fn)
	if tf.Kind() != reflect.Func {
		return
	}

	tf.NumIn()
}
