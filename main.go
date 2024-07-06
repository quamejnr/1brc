package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Record struct {
	count         int
	max, min, sum float64
}

type records map[string]*Record

func sort(data records) []string {
	cities := []string{}

	for c := range data {
		cities = append(cities, c)
	}

	slices.Sort(cities)
	return cities
}

func brc(r io.Reader) records {
	records := records{}
	s := bufio.NewScanner(r)
	for s.Scan() {
		if line := s.Text(); strings.TrimSpace(line) != "" {
			d := strings.Split(line, ";")
			city := d[0]
			temp, _ := strconv.ParseFloat(d[1], 32)

			r, ok := records[city]
			if !ok {
				r = &Record{
					max: temp,
					min: temp,
				}
				records[city] = r
				// pass new cities into a slice to be sorted later
			}

			r.max = math.Max(r.max, temp)
			r.min = math.Min(r.min, temp)
			r.count++
			r.sum += temp
		}
	}

	return records

}

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	records := brc(f)
	cities := sort(records)

	fmt.Fprint(os.Stdout, "{")

	for _, city := range cities {
		if r, ok := records[city]; ok {
			fmt.Fprintf(os.Stdout, "%s=%.1f/%.1f/%.1f,", city, r.min, (r.sum / float64(r.count)), r.max)
		}
	}
	fmt.Fprint(os.Stdout, "}\n")

}
