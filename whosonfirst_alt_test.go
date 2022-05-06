package spr

import (
	_ "fmt"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestWhosOnFirstAltSPR(t *testing.T) {

	path_fixtures, err := filepath.Abs("fixtures")

	if err != nil {
		t.Fatalf("Failed to derive path for fixtures, %v", err)
	}

	path_feature := filepath.Join(path_fixtures, "102147495-alt-quattroshapes.geojson")

	r, err := os.Open(path_feature)

	if err != nil {
		t.Fatalf("Failed to open %s, %v", path_feature, err)
	}

	defer r.Close()

	body, err := io.ReadAll(r)

	if err != nil {
		t.Fatalf("Failed to read %s, %v", path_feature, err)
	}

	s, err := WhosOnFirstAltSPR(body)

	if err != nil {
		t.Fatalf("Failed to derive SPR for %s, %v", path_feature, err)
	}

	if s.Id() != "102147495" {
		t.Fatalf("Invalid ID")
	}

	if s.Name() != "102147495 alt geometry (quattroshapes)" {
		t.Fatalf("Invalid name: '%s'", s.Name())
	}

	if s.Placetype() != "alt" {
		t.Fatalf("Invalid placetype")
	}

}
