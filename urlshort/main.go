package main

import (
	"fmt"
	"log"
	"net/http"
)

// func main() {
// 	mux := defaultMux()

// 	// Build the MapHandler using the mux as the fallback
// 	pathsToUrls := map[string]string{
// 		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
// 		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
// 	}
// 	mapHandler := handler.MapHandler(pathsToUrls, mux)

// 	// Build the YAMLHandler using the mapHandler as the
// 	// fallback
// 	// yaml := `
// 	// - path: /urlshort
// 	//   url: https://github.com/gophercises/urlshort
// 	// - path: /urlshort-final
// 	//   url: https://github.com/gophercises/urlshort/tree/solution
// 	// `
// 	// yamlHandler, err := handler.YAMLHandler([]byte(yaml), mapHandler)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	fmt.Println("Starting the server on :8080")
// 	// http.ListenAndServe(":8080", yamlHandler)
// 	http.ListenAndServe(":8080", mapHandler)
// }

// func defaultMux() *http.ServeMux {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", hello)
// 	return mux
// }

// func hello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello, world!")
// }

type pounds float32

func (p pounds) String() string {
	return fmt.Sprintf("£%.2f", p)
}

type database map[string]pounds

func (d database) foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "foo: %s\n", d["foo"])
}

func (d database) bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "bar: %s\n", d["bar"])
}

func (d database) baz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "baz: %s\n", d["baz"])
}

func main() {
	db := database{
		"foo": 1,
		"bar": 2,
		"baz": 3,
	}

	http.HandleFunc("/foo", db.foo)
	http.HandleFunc("/bar", db.bar)
	http.HandleFunc("/baz", db.baz)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
