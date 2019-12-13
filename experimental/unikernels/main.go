// CGO_ENABLED=0 GOOS=linux go build -a -tags netgo  -ldflags '-extldflags "-static"'
// ops run main
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	log.Println("Listening...on 8080")
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
