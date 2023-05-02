package configs

import (
	"fmt"
	"reflect"
	"sync"
)

type Fetcher[T any] interface {
	Fetch() (T, error)
}

type CachedFetcher[T any] struct {
	fetcher Fetcher[T]
	payload T
	once    sync.Once
}

func RegisterCachedFetcher[T any](fetcher Fetcher[T]) {
	RegisterFetcher[T](&CachedFetcher[T]{fetcher: fetcher})
}

func (f *CachedFetcher[T]) Fetch() (T, error) {
	var err error
	f.once.Do(func() {
		f.payload, err = f.fetcher.Fetch()
	})
	return f.payload, err
}

var fetchers = make(map[string]any)

func RegisterFetcher[T any](fetcher Fetcher[T]) {
	var t T
	fetchers[reflect.TypeOf(t).String()] = fetcher
}

func Fetch[T any]() (T, error) {
	return getConfig[T]()
}

func Must[T any]() T {
	ins, _ := Fetch[T]()
	return ins
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

func getConfig[T any]() (t T, ok error) {
	if f, ok := getFetcher[T](); ok {
		if fetcher, ok := (any)(f).(Fetcher[T]); ok {
			return fetcher.Fetch()
		}
	}
	return t, fmt.Errorf("config not found: %s", reflect.TypeOf(t).String())
}
