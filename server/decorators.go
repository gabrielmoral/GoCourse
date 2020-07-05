package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/pabloos/http/cache"
	"github.com/pabloos/http/greet"
)

func POST(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func Debug(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer h.ServeHTTP(w, r)

		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}

		w.Write([]byte(dump))
	}
}

func Delay(delay time.Duration, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer h.ServeHTTP(w, r)

		time.Sleep(delay)
	}
}

func Cached(h http.HandlerFunc) http.HandlerFunc {
	greetings := &cache.Greetings{
		List: make(map[string][]string),
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var t greet.Greet

		b := bytes.NewBuffer(make([]byte, 0))
		reader := io.TeeReader(r.Body, b)
		r.Body = ioutil.NopCloser(b)
		err := json.NewDecoder(reader).Decode(&t)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		cached, found := greetings.Add(t)

		if found {
			fmt.Fprint(w, cached.Print())
		} else {
			h.ServeHTTP(w, r)
		}
	}
}
