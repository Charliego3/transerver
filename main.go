package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/ServiceWeaver/weaver"
)

func main() {
	if err := weaver.Run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

type app struct {
	weaver.Implements[weaver.Main]
	reverser weaver.Ref[Reverser]
}

func (app *app) Init(context.Context) error {
	log.Println("app inited ....")
	return nil
}

func (app *app) Main(ctx context.Context) error {
	// Get a network listener on address "localhost:12345".
	opts := weaver.ListenerOptions{LocalAddress: "localhost:12345"}
	lis, err := app.Listener("transerver", opts)
	if err != nil {
		return err
	}
	fmt.Printf("transerver listener available on %v\n", lis)

	// Serve the /hello endpoint.
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "World"
		}
		reversed, err := app.reverser.Get().Reverse(ctx, name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, _ = fmt.Fprintf(w, "Hello, %s!\n", reversed)
	})
	return http.Serve(lis, nil)
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
