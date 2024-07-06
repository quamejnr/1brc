package main

import (
	"bytes"
	"testing"
)

func TestBRC(t *testing.T) {
	d := `
Halifax;12.9
Cabo San Lucas;14.9
Adelaide;15.0
Pittsburgh;9.7
Karachi;15.4
Dodoma;22.2
  `
	want := records{
		"Adelaide":       &Record{min: 15.0, max: 15.0, count: 1, sum: 15.0},
		"Cabo San Lucas": &Record{min: 14.9, max: 14.9, count: 1, sum: 14.9},
		"Dodoma":         &Record{min: 22.2, max: 22.2, count: 1, sum: 22.2},
		"Halifax":        &Record{min: 12.9, max: 12.9, count: 1, sum: 12.9},
		"Karachi":        &Record{min: 15.4, max: 15.4, count: 1, sum: 15.4},
		"Pittsburgh":     &Record{min: 9.7, max: 9.7, count: 1, sum: 9.7},
	}

	r := bytes.NewBufferString(d)

	got, _ := brc(r)

	if compareMaps(got, want) {
		t.Errorf("want %v got %v\n", want, got)
	}

}

// compareMaps compares two maps for equality
func compareMaps(map1, map2 records) bool {
	// If both maps are nil, they are equal
	if map1 == nil && map2 == nil {
		return true
	}

	// If one of them is nil and the other is not, they are not equal
	if map1 == nil || map2 == nil {
		return false
	}

	// Compare lengths
	if len(map1) != len(map2) {
		return false
	}

	// Compare each key and value
	for k, v := range map1 {
		if val, ok := map2[k]; !ok || val != v {
			return false
		}
	}

	return true
}
