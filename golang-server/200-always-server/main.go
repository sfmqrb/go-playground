//from following git repo:
//https://github.com/caarlos0/sinkhole

package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var addr = ":1809"
	if p := os.Getenv("PORT"); p != "" {
		addr = ":" + p
	}

	mux := http.NewServeMux()
	mux.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			sleep := r.URL.Query().Get("sleep")
			if sleep != "" {
				d, err := time.ParseDuration(sleep)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				time.Sleep(d)
			}

			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					return
				}
			}(r.Body)
			log.Println(r.URL)
			_, err := fmt.Fprintln(w, "ok")
			if err != nil {
				return
			}
		},
	)

	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	done := make(chan os.Signal, 1)

	signal.Notify(
		done,
		os.Interrupt,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Println("server started on " + addr)
	<-done
	log.Println("server stop requested")

	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown failed:%+v", err)
	}
	log.Print("server exit properly")
}
