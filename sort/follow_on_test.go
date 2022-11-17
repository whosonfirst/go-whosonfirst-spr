package sort

import (
	"context"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-spr/v2"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestApplyFollowOnSorters(t *testing.T) {

	input := []string{
		"1008184051.geojson", // Poop Emoji Rock
		"85688637.geojson",   // California,
		"420561633.geojson",  // Super Bowl City
		"85922583.geojson",   // San Francisco
		"101736545.geojson",  // Montreal
	}

	expected := []string{
		"85688637",   // California
		"101736545",  // Montreal
		"85922583",   // San Francisco
		"420561633",  // Super Bowl City
		"1008184051", // Poop Emoji Rock
	}

	ctx := context.Background()

	pt_sorter, err := NewSorter(ctx, "placetype://")

	if err != nil {
		t.Fatalf("Failed to create placetype sorter, %v", err)
	}

	nm_sorter, err := NewSorter(ctx, "name://")

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
		t.Fatalf("Failed to sort by placetype, %v", err)
	}

	key_func := func(ctx context.Context, s spr.StandardPlacesResult) (string, error) {
		return s.Placetype(), nil
	}

	sorted_follow_on, err := ApplyFollowOnSorters(ctx, sorted.Results(), key_func, nm_sorter)

	if err != nil {
		t.Fatalf("Failed to apply follow on sorters, %v", err)
	}

	if len(sorted_follow_on) != len(expected) {
		t.Fatalf("Invalid result count: %d", len(sorted.Results()))
	}

	for idx, s := range sorted_follow_on {

		if s.Id() != expected[idx] {
			t.Fatalf("Unexpected sort result at offset %d. Expected '%s' but got '%s'", idx, expected[idx], s.Id())
		}

		fmt.Sprintf("%d %s %s\n", idx, s.Name(), expected[idx])
	}
}
