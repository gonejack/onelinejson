package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	for _, name := range os.Args[1:] {
		f, e := os.Open(name)
		if e != nil {
			log.Fatalf("open file failed: %s", e)
		}

		var v interface{}
		e = json.NewDecoder(f).Decode(&v)
		if e != nil {
			log.Fatalf("decode json failed: %s", e)
		}
		f.Close()

		o := strings.TrimSuffix(name, filepath.Ext(name)) + ".oneline.json"
		w, e := os.Create(o)
		if e != nil {
			log.Fatalf("cannot create file %s: %s", o, e)
		}

		e = json.NewEncoder(w).Encode(v)
		if e != nil {
			log.Fatalf("encode json failed: %s", e)
		}
		w.Close()
	}
}
