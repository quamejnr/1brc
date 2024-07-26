package main

import (
	"1brcme/brc1"
	"1brcme/brc2"
	"bytes"
	"os"
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

func BenchmarkTestBRC(b *testing.B) {
	b.Run("brc-1", func(b *testing.B) {
		r, err := os.Open("measurements-1.txt")
		if err != nil {
			b.Fatal("error opening file")
		}
		defer r.Close()
		for range b.N {
			brc1.Brc(r)
		}
	})
	b.Run("brc2", func(b *testing.B) {
		r, err := os.Open("measurements-1.txt")
		if err != nil {
			b.Fatal("error opening file")
		}
		for range b.N {
			brc2.Brc(r)
		}
	})
}
