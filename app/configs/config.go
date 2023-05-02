package configs

import (
	"reflect"
	"sync"
)

type Fetcher[T any] interface {
	Fetch() (T, bool)
}

type CachedFetcher[T any] struct {
	fetcher Fetcher[T]
	payload T
	once    sync.Once
}

func RegisterCachedFetcher[T any](fetcher Fetcher[T]) {
	RegisterFetcher[T](&CachedFetcher[T]{fetcher: fetcher})
}

func (f *CachedFetcher[T]) Fetch() (T, bool) {
	success := true
	f.once.Do(func() {
		f.payload, success = f.fetcher.Fetch()
	})
	return f.payload, success
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

func getConfig[T any]() (t T, ok bool) {
	if f, ok := getFetcher[T](); ok {
		if fetcher, ok := (any)(f).(Fetcher[T]); ok {
			return fetcher.Fetch()
		}
	}
	return t, false
}
