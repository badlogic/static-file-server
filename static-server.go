package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	fs := http.FileServer(http.Dir(os.Args[1]))
	http.Handle("/", fs)

	log.Print("Listening on http://localhost:" + os.Args[2])
	err := http.ListenAndServe(":" + os.Args[2], nil)
	if err != nil {
		log.Fatal(err)
	}
}