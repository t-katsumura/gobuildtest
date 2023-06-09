package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	w  = bufio.NewWriterSize(os.Stdout, 1<<12)
	lg = log.New(w, "", log.LstdFlags)
)

/*
okFunc is a HTTP handler function which simply reply with "ok".
This is extremely simple so to endure stress test.
*/
func okFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"OK"}`))
	lg.Println(r.Method, r.URL.Path)
}

func main() {

	addr := ":10001"
	lg.Println("starting server at", addr)
	w.Flush()

	go func() {
		for {
			time.Sleep(5 * time.Second)
			w.Flush()
		}
	}()

	http.Handle("/", http.HandlerFunc(okFunc))
	if err := http.ListenAndServe(addr, nil); err != nil && err != http.ErrServerClosed {
		lg.Println(err)
	}
	w.Flush()

}
