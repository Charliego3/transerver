package main

import (
	"net/http"
	"rock"
	"rock/hapi"
)

func main() {
	app := rock.NewApp(rock.WithAddress(":9090"), rock.WithHTTP())
	app.ConfigureHTTP(func(hs *hapi.Server) {
		hs.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"success": true}`))
		})
	})
	app.MustRun()
}