/*
Copyright 2018 morimolymoly

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Shitty Routing Server")

	http.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/redirect/" {
		}
		http.Redirect(w, r, "http://google.com/", http.StatusFound)
		fmt.Printf("%#v", w.Header()["Location"])
	})

	http.HandleFunc("/shitty", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<html><body><h1>Hello, World</h1><p>%#v</p></body></html>", w.Header())
	})

	http.HandleFunc("/shitty/redirect", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "google.com", http.StatusFound)
		fmt.Printf("%v", w.Header()["Location"])
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
