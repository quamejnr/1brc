package main

import (
	"bufio"
	"fmt"
	"io"
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

func sort(cities []string) {
	slices.Sort(cities)
}

func brc(r io.Reader) (records, []string) {
	cities := []string{}
	records := records{}
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		city, val, found := strings.Cut(line, ";")
		if found {
			temp, _ := strconv.ParseFloat(val, 64)

			r, ok := records[city]
			if !ok {
				r = &Record{
					max: temp,
					min: temp,
				}
				records[city] = r
				cities = append(cities, city)
			}

			r.max = max(r.max, temp)
			r.min = min(r.min, temp)
			r.count++
			r.sum += temp

		}

	}

	return records, cities

}

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	records, cities := brc(f)
	sort(cities)

	fmt.Fprint(os.Stdout, "{")

	for _, city := range cities {
		if r, ok := records[city]; ok {
			fmt.Fprintf(os.Stdout, "%s=%.1f/%.1f/%.1f,", city, r.min, (r.sum / float64(r.count)), r.max)
		}
	}
	fmt.Fprint(os.Stdout, "}\n")

}
