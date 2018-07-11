package main

import (
	"log"
	"os"
	"net/http"
)

type Importer interface {
	Import() (*Stats, error)
}

var importer Importer

func init() {
	if len(os.Args) < 2 {
		log.Fatal("You do not set proxy container name. Usage: ./statist proxy_container_name")
	}
	importer = &Docker{os.Args[1]}
}

func main() {
	log.Println("Statist starting")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Statist handle requst starting")
		s, err := importer.Import()
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("proxy\n"))
			return
		}
		w.WriteHeader(200)
		w.Write([]byte(s.ToInfluxFormat() + "\n"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
