package configs

import "reflect"

type Fetcher[T any] interface {
	FetchConfig() (T, bool)
}

var fetchers = make(map[string]any)

func RegisterFetcher[T any](fetcher Fetcher[T]) {
	var t T
	fetchers[reflect.TypeOf(t).String()] = fetcher
}

func Fetch[T any]() (T, bool) {
	return getConfig[T]()
}

func getFetcher[T any]() (Fetcher[T], bool) {
	var t T
	if (any)(t) == nil {
		return nil, false
	}

	if f, ok := fetchers[reflect.TypeOf(t).String()]; ok {
		return f.(Fetcher[T]), true
	}
	return nil, false
}

func getConfig[T any]() (T, bool) {
	if f, ok := getFetcher[T](); ok {
		if fetcher, ok := (any)(f).(Fetcher[T]); ok {
			return fetcher.FetchConfig()
		}
	}
	var t T
	return t, false
}
