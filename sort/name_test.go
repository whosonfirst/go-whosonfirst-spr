package sort

import (
	"context"
	"github.com/whosonfirst/go-whosonfirst-spr/v2"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestSortByName(t *testing.T) {

	input := []string{
		"1008184051.geojson", // Poop Emoji Rock
		"85688637.geojson",   // California,
		"420561633.geojson",  // Super Bowl City
		"85922583.geojson",   // San Francisco
	}

	expected := []string{
		"85688637",   // California
		"1008184051", // Poop Emoji Rock
		"85922583",   // San Francisco
		"420561633",  // Super Bowl City
	}

	ctx := context.Background()

	pt_sorter, err := NewSorter(ctx, "name://")

	if err != nil {
		t.Fatalf("Failed to create name sorter, %v", err)
	}

	path_fixtures, err := filepath.Abs("../fixtures")

	if err != nil {
		t.Fatalf("Failed to derive path for fixtures, %v", err)
	}

	count := len(input)

	results := make([]spr.StandardPlacesResult, count)

	for idx, path := range input {

		path_feature := filepath.Join(path_fixtures, path)

		r, err := os.Open(path_feature)

		if err != nil {
			t.Fatalf("Failed to open %s, %v", path_feature, err)
		}

		defer r.Close()

		body, err := io.ReadAll(r)

		if err != nil {
			t.Fatalf("Failed to read %s, %v", path_feature, err)
		}

		s, err := spr.WhosOnFirstSPR(body)

		if err != nil {
			t.Fatalf("Failed to derive SPR for %s, %v", path_feature, err)
		}

		results[idx] = s
	}

	r := &testableStandardPlacesResults{
		results: results,
	}

	sorted, err := pt_sorter.Sort(ctx, r)

	if err != nil {
		t.Fatalf("Failed to sort by name, %v", err)
	}

	for idx, s := range sorted.Results() {

		if s.Id() != expected[idx] {
			t.Fatalf("Unexpected sort result at offset %d. Expected '%s' but got '%s'", idx, expected[idx], s.Id())
		}
	}
}
