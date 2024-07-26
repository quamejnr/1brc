package brc1

import (
	"bytes"
	"reflect"
	"testing"
)

func makeData() *bytes.Buffer {
	d := `
Halifax;12.9
Cabo San Lucas;14.9
Adelaide;15.0
Pittsburgh;9.7
Karachi;15.4
Dodoma;22.2
  `
	r := bytes.NewBufferString(d)
	return r

}

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
	r = makeData()

	t.Run("test brc", func(t *testing.T) {
		got, _ := Brc(r)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v got %v\n", want, got)
		}
	})

}
