package brc2

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"slices"
)

type Record struct {
	max, min, sum, count int
}

type records map[string]*Record

func sort(cities []string) {
	slices.Sort(cities)
}

func Brc(r io.Reader) (records, []string) {
	cities := []string{}
	records := records{}
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Bytes()
		city, val, found := bytes.Cut(line, []byte(";"))
		if found {
			temp := parse(val)
			r, ok := records[string(city)]
			if !ok {
				r = &Record{
					max: temp,
					min: temp,
				}
				records[string(city)] = r
				cities = append(cities, string(city))
			}

			r.max = max(r.max, temp)
			r.min = min(r.min, temp)
			r.count++
			r.sum += temp
		}
	}
	return records, cities

}

func parse(temp []byte) int {
	var res int
	isNegative := false
	if temp[0] == '-' {
		temp = temp[1:]
		isNegative = true
	}
	for _, v := range temp {
		if v != '.' {
			res = (res * 10) + int(v-'0')
		}
	}
	if isNegative {
		res = -res
	}
	return res
}

func PrintBRC(file string) {
	f, _ := os.Open(file)
	defer f.Close()

	records, cities := Brc(f)
	slices.Sort(cities)

	fmt.Fprint(os.Stdout, "{")

	for _, city := range cities {
		if r, ok := records[city]; ok {
			fmt.Fprintf(os.Stdout, "%s=%.1f/%.1f/%.1f,", city, float64(r.min/10), float64(r.sum/r.count)/10, float64(r.max/10))
		}
	}
	fmt.Fprint(os.Stdout, "}\n")

}
