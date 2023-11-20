package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ServiceWeaver/weaver"
	"github.com/charliego3/shandler"
)

func main() {
	slog.SetDefault(slog.New(shandler.NewTextHandler(shandler.WithCaller())))
	if err := weaver.Run(context.Background(), serve); err != nil {
		slog.Error("application running got an error", "err", err)
	}
}

type app struct {
	weaver.Implements[weaver.Main]
	reverser weaver.Ref[Reverser]
	endpoint weaver.Listener
}

func serve(ctx context.Context, app *app) error {
	http.HandleFunc("/reverse", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		var reverser Reverser = app.reverser.Get()
		reversed, err := reverser.Reverse(ctx, key)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, `{"code":50001,"message":"%s"}`, err)
			return
		}

		w.WriteHeader(200)
		fmt.Fprintf(w, `{"code": 20000, "message": "success", "payload": "%s"}`, reversed)
		app.Logger(ctx).Info("request is completed")
	})
	return http.Serve(app.endpoint, nil)
}

type Reverser interface {
	Reverse(context.Context, string) (string, error)
}

// Implementation of the Reverser component.
type reverser struct {
	weaver.Implements[Reverser]
}

func (r *reverser) Reverse(_ context.Context, s string) (string, error) {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}
	return string(runes), nil
}
