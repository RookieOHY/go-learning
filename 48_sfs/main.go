package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("p", 8080, "Port to listen on")
	dir := flag.String("d", ".", "Directory to serve")
	flag.Parse()

	startServer(*port, *dir)
}

func startServer(port int, dir string) {
	fmt.Printf("Starting server on port http://localhost:%d serving directory %s\n", port, dir)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, dir+r.URL.Path)
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}
