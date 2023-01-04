package main

import (
	"log"
	"net/http"
	"os"
	"github.com/pkg/browser"
)

func main() {
	fs := http.FileServer(http.Dir(os.Args[1]))
	http.Handle("/", fs)

	log.Print("Listening on http://localhost:" + os.Args[2])
	if (len(os.Args) == 3) {
		browser.OpenURL("http://localhost:" + os.Args[2])
	}
	if (len(os.Args) == 4) {
		browser.OpenURL("http://localhost:" + os.Args[2] + "/" + os.Args[3])
	}
	err := http.ListenAndServe(":"+os.Args[2], nil)
	if err != nil {
		log.Fatal(err)
	}
}
