package main

import (
	"context"
	"flag"
	"io"
	"log"
	"log/slog"

	"github.com/whosonfirst/go-whosonfirst-iterate/v3"
	"github.com/whosonfirst/go-whosonfirst-spr/v2"
	"github.com/whosonfirst/go-whosonfirst-uri"
)

func main() {

	ctx := context.Background()

	var iterator_uri string

	flag.StringVar(&iterator_uri, "iterator-uri", "repo://", "")

	flag.Parse()

	iterator_sources := flag.Args()

	iter, err := iterate.NewIterator(ctx, iterator_uri)

	if err != nil {
		log.Fatalf("Failed to create new iterator, %v", err)
	}

	for rec, err := range iter.Iterate(ctx, iterator_sources...) {

		if err != nil {
			log.Fatalf("Failed to iterate sources, %v", err)
		}

		defer rec.Body.Close()

		_, uri_args, err := uri.ParseURI(rec.Path)

		if err != nil {
			log.Fatalf("Failed to parse URI for %s, %v", rec.Path, err)
		}

		if uri_args.IsAlternate {
			continue
		}

		body, err := io.ReadAll(rec.Body)

		if err != nil {
			log.Fatalf("Failed to read body for %s, %v", rec.Path, err)
		}

		_, err = spr.WhosOnFirstSPR(body)

		if err != nil {
			slog.Warn("Failed to derive SPR", "path", rec.Path, "error", err)
		}
	}
}
