package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/jkakar/dashi"
	"github.com/joeshaw/envdecode"
)

type config struct {
	Port int `env:"PORT,default=8080"`
}

func main() {
	cfg := &config{}
	if err := envdecode.StrictDecode(cfg); err != nil {
		log.Fatalf("loading config failed: %q", err)
	}
	manifest := &dashi.Manifest{}
	for _, filename := range os.Args[1:] {
		log.Printf("loading manifest data from %s", filename)
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatalf("error reading file %s: %q", filename, err)
		}
		if err := dashi.Unmarshal(data, manifest); err != nil {
			log.Fatal(err)
		}
	}

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.Handle("/", dashi.NewSearchHandler(manifest))
	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: mux,
	}
	log.Printf("listening on port %d", cfg.Port)
	log.Fatalf("listen and serve failed: %q", s.ListenAndServe())
}
