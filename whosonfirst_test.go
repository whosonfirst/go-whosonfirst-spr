package spr

import (
	_ "fmt"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestWhosOnFirstSPR(t *testing.T) {

	path_fixtures, err := filepath.Abs("fixtures")

	if err != nil {
		t.Fatalf("Failed to derive path for fixtures, %v", err)
	}

	path_feature := filepath.Join(path_fixtures, "101736545.geojson")

	r, err := os.Open(path_feature)

	if err != nil {
		t.Fatalf("Failed to open %s, %v", path_feature, err)
	}

	defer r.Close()

	body, err := io.ReadAll(r)

	if err != nil {
		t.Fatalf("Failed to read %s, %v", path_feature, err)
	}

	s, err := WhosOnFirstSPR(body)

	if err != nil {
		t.Fatalf("Failed to derive SPR for %s, %v", path_feature, err)
	}

	if s.Id() != "101736545" {
		t.Fatalf("Invalid ID")
	}

	if s.Name() != "Montreal" {
		t.Fatalf("Invalid name")
	}

	if s.Placetype() != "locality" {
		t.Fatalf("Invalid locality")
	}
	
}
