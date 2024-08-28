package brc4

import (
	"fmt"
	"io"
	"os"
	"sort"
)

type Record struct {
	max, min, sum, count int
}

type records map[string]*Record

// The lines being returned are in the form of city;xx.x
// We are assured the xx.x will be max 4 bytes but the city names can be longer
// Thus, instead of starting from the start of the line to find the delimiter,
// we start from the end being assured, we'll have to check max 4 bytes before we find our delimiter.
func cut(d []byte, delim byte) (before, after []byte, found bool) {
	for i := len(d) - 1; i >= 0; i-- {
		if d[i] == delim {
			return d[:i], d[i+1:], true
		}
	}
	return nil, nil, false
}

func Brc(r io.Reader) (records, []string) {
	cities := []string{}
  records := records{}

	// Create a buffer size of 1MB.
	buf := make([]byte, 1<<20)
	readStart := 0
	// We're reading 1MB into our buffer and then splitting the read data into new lines.
	// We process them and then copy the remaining data into the buffer and read the rest of the data.
	// This means we may read data into our buffer like [Lisbon;1.23\nAcc]
	// What we do is we'll process the data to the new line character: '\n' until there are no more new lines. Then copy the remainder back into the buffer.
	// Keep in mind the buffer still contains the data and we're just using slices of them. So we'll still have buf containing [Lisbon;1.23\nAcc] before we copy the remainder.
	// Ergo when we copy the remainder `Acc` into the buffer, knowing the behavior of Go slices, what we'll have is [Accbon;1.23\nAcc].
	// `Acc` will replace the first 3 indices of the buffer slice.
	// This is because the array the buffer slice points to still contain our buffer data.
	// This means when reading again, we'll have to let our buffer read after `Acc` thus we adjust our readStart to be the index of the remaining bytes copied.
	for {
		_, err := r.Read(buf[readStart:])
		if err == io.EOF {
			break
		}
		lineStart := 0
		for i, v := range buf {
			if v == '\n' {
				// Line is index from where we started reading to finding a newline
				line := buf[lineStart:i]
				// New line starts after '\n' thus i + 1
				lineStart = i + 1
				// readStart becomes our lineStart.
				// This helps keep track of where we last read to so we know where our remaining data comes from.
				readStart = lineStart

				city, val, found := cut(line, ';')
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

		}
		// copy the remaining chunk of data
		readStart = copy(buf, buf[readStart:])

	}

	return records, cities

}

// parseTemp takes bytes and returns an integer
func parseTemp(temp []byte) int {
	var res int
	sign := 1
	// Given a string like "-12.3". We know the first byte will carry the sign
	// We grab that if the first element is "-" and set isNegative flag to true
	if temp[0] == '-' {
		temp = temp[1:]
		sign = -1
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
	for _, v := range temp {
		if v != '.' {
			res = (res * 10) + int(v-'0')
		}
	}
	return res * sign
}

func PrintBRC(file string) {
	f, _ := os.Open(file)
	defer f.Close()

	records, cities := Brc(f)
	sort.Strings(cities)

	fmt.Fprint(os.Stdout, "{")

	for _, city := range cities {
		if r, ok := records[city]; ok {
			fmt.Fprintf(os.Stdout, "%s=%.1f/%.1f/%.1f,", city, float64(r.min/10), float64(r.sum/r.count)/10, float64(r.max/10))
		}
	}
	fmt.Fprint(os.Stdout, "}\n")

}
