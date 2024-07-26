package brc2

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
	want := records{
		"Adelaide":       &Record{min: 150, max: 150, count: 1, sum: 150},
		"Cabo San Lucas": &Record{min: 149, max: 149, count: 1, sum: 149},
		"Dodoma":         &Record{min: 222, max: 222, count: 1, sum: 222},
		"Halifax":        &Record{min: 129, max: 129, count: 1, sum: 129},
		"Karachi":        &Record{min: 154, max: 154, count: 1, sum: 154},
		"Pittsburgh":     &Record{min: 97, max: 97, count: 1, sum: 97},
	}

	r := makeData()

	t.Run("test brc", func(t *testing.T) {
		got, _ := Brc(r)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v got %v\n", want, got)
		}
	})

}
