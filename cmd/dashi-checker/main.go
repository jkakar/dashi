package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/jkakar/dashi"
)

func main() {
	manifest := &dashi.Manifest{}
	for _, filename := range os.Args[1:] {
		log.Printf("loading manifest data from %s", filename)
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatalf("error reading file %s: %q", filename, err)
		}
		if err := dashi.Unmarshal(data, manifest); err != nil {
			log.Fatalf("error unmarshaling data: %q", err)
		}
	}
}
