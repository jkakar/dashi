package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joeshaw/envdecode"
)

type config struct {
	Port int `env:"PORT,default=8080"`
}

func http404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func main() {
	cfg := &config{}
	if err := envdecode.StrictDecode(cfg); err != nil {
		log.Fatalf("loading config failed: %q", err)
	}
	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: http.HandlerFunc(http404),
	}
	log.Printf("listening on port %d", cfg.Port)
	log.Fatalf("listen and serve failed: %q", s.ListenAndServe())
}
