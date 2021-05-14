package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	for _, a := range os.Args[1:] {
		fd, err := os.Open(a)
		if err != nil {
			log.Fatalf("open file failed: %s", err)
		}

		var v interface{}
		err = json.NewDecoder(fd).Decode(&v)
		if err != nil {
			log.Fatalf("decode json failed: %s", err)
		}

		data, err := json.Marshal(v)
		if err != nil {
			log.Fatalf("encode json failed: %s", err)
		}

		output := strings.TrimSuffix(a, filepath.Ext(a)) + ".unfmt.json"
		err = ioutil.WriteFile(output, data, 0666)
		if err != nil {
			log.Fatalf("write file failed: %s", err)
		}
	}
}
