package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Adapter func(http.Handler, *interface{}) http.Handler

func Adapt(handler interface{}, adapters ...Adapter) (h http.Handler) {
	var response interface{}
	switch handler := handler.(type) {
	case http.Handler:
		h = handler
	case func(http.ResponseWriter, *http.Request):
		h = http.HandlerFunc(handler)
	case func(*http.Request) interface{}:
		h = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			response = handler(r)
		})
	default:
		log.Fatal("Invalid Adapt Handler", handler)
	}

	for _, adapter := range adapters {
		h = adapter(h, &response)
	}

	return h
}

func Logging(l *log.Logger) Adapter {
	return func(h http.Handler, response *interface{}) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			l.Println("http:", r.Method, r.URL.Path, r.UserAgent())
			h.ServeHTTP(w, r)
		})
	}
}

func UseMethod(name string) Adapter {
	return func(h http.Handler, response *interface{}) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != name {
				http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			} else {
				h.ServeHTTP(w, r)
			}
		})
	}
}

func API(debug bool) Adapter {
	return func(h http.Handler, response *interface{}) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
			w.Header().Set("Content-Type", "application/json")
			payload := make(map[string]interface{})
			if e, ok := (*response).(error); ok {
				if debug {
					fmt.Println("handler returned error", e.Error())
				}
				payload["error"] = e.Error()
			} else if s, ok := (*response).(fmt.Stringer); ok {
				payload["data"] = s.String()
			} else {
				payload["data"] = response
			}
			json.NewEncoder(w).Encode(payload)
		})
	}
}
