package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/feature"
	"github.com/whosonfirst/go-whosonfirst-spr/whosonfirst"
	"log"
)

func main() {

	flag.Parse()

	for _, path := range flag.Args() {

		f, err := feature.LoadWOFFeatureFromFile(path)

		if err != nil {
			log.Fatal(err)
		}

		s, err := whosonfirst.NewSPRFromFeature(f)

		if err != nil {
			log.Fatal(err)
		}

		body, err := json.Marshal(s)
		fmt.Println(string(body))
	}
}
