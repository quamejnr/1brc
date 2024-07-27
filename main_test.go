package main

import (
	"1brcme/brc1"
	"1brcme/brc2"
	"1brcme/brc3"
	"os"
	"testing"
)

func BenchmarkTestBRC(b *testing.B) {
	f := "measurements-1.txt"
	// silence print statements
	null, _ := os.Open(os.DevNull)
	os.Stdout = null
	null.Close()

	b.Run("brc-1", func(b *testing.B) {
		for range b.N {
			brc1.PrintBRC(f)
		}
	})
	b.Run("brc2", func(b *testing.B) {
		for range b.N {
			brc2.PrintBRC(f)
		}
	})
	b.Run("brc3", func(b *testing.B) {
		for range b.N {
			brc3.PrintBRC(f)
		}
	})
}
