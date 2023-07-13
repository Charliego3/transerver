package app

import "reflect"

func Default[T any](source T, defval T) T {
	return DefaultFunc(source, func() T {
		return defval
	})
}

func DefaultFunc[T any](source T, creator func() T) T {
	v := reflect.ValueOf(source)
	if v.IsNil() {
		return creator()
	}
	return source
}
