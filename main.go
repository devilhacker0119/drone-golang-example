package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

// HelloWorld for hello world
func HelloWorld() string {
	return "Hello World, golang workshop!"
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s, I love %s!", HelloWorld(), r.URL.Path[1:])
}

func main() {
	var port string
	flag.StringVar(&port, "port", "8080", "server port")
	flag.StringVar(&port, "p", "8080", "server port")
	if p, ok := os.LookupEnv("PORT"); ok {
		port = p
	}
	flag.Parse()
	http.HandleFunc("/", handler)
	log.Println("http server run on " + port + " port")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
