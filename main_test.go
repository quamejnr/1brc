package main

import (
	"1brcme/brc1"
	"bytes"
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
	r := makeData()
	b.Run("brc1", func(b *testing.B) {
		for range b.N {
			brc1.Brc(r)
		}
	})
}
