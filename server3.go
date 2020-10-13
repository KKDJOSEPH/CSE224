// Server1 is a minimal "echo" server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	http.HandleFunc("/gendata", gendataHandler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))

}

// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func gendataHandler(w http.ResponseWriter, r *http.Request){
	numBytes, err :=r.URL.Query()["numBytes"]
	if !err || numBytes==nil{
		log.Println("There's no numBytes")
		return
	}

	length, _ :=strconv.Atoi(numBytes[0])
	str :=strings.Repeat("-", length)
	fmt.Fprintf(w, "%s", str)
}

