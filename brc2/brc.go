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
			temp := parseTemp(val)
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

// parseTemp takes bytes and returns an integer
func parseTemp(temp []byte) int {
	var res int
	isNegative := false
	// Given a string like "-12.3". We know the first byte will carry the sign
	// We grab that if the first element is "-" and set isNegative flag to true
	if temp[0] == '-' {
		temp = temp[1:]
		isNegative = true
	}
	// The encoding bytes of characters is in sequence like 'abcd' or '0123'.
	// This means if I have `a` having an u8 value of 2 then b will be 3
	// Using that same logic if I have 0 being 48 then 1 will be 49
	// This means I can get the int value of 1 by subtracting it from 0 thus 49-48=1
	// Also, given a value of 123, this is the same as 1*100 + 2*10 + 3.
	// This means I can multiply each value by 10 and add to the next value
	// Thus in a loop, you'll have:
	//
	// | iteration | value          |
	// |-----------|----------------|
	// | 1st       | (0*10)+1=1     |
	// | 2nd       | (1*10)+2=12    |
	// | 3rd       | (12*10)+3=123  |
	//
	// This gives us the int value we need.
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
