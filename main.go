package main

import (
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"os"
	"strings"
	"time"
)

var allowRussianQuotes bool

func isRussian(s string) bool {
	x := 0
	for _, c := range s {
		if int(c) > 127 {
			x++
			if x > 12 {
				return true
			}
		}
	}
	return false
}

func splitString(s string, maxLength int) []string {
	ru := isRussian(s)
	if ru {
		maxLength *= 7
		maxLength /= 4
	}
	var arr []string
	var suffix string
	c := strings.Index(s, "(C)")
	if c > 0 {
		suffix = s[c:]
		s = s[:c-1]
	}
	for len(s) > maxLength {
		j := maxLength - 1
		if len(s) <= j {
			j = len(s)
		}
		for s[j] != ' ' && j < len(s) {
			j--
		}
		arr = append(arr, s[:j])
		s = s[j+1:]
	}

	if len(s) > 0 {
		arr = append(arr, s)
	}
	if len(suffix) > 0 {
		arr = append(arr, suffix)
	}
	return arr
}

func getRandomQuote() (res string) {
	q := getQuotes()
	seconds := time.Now().Unix()
	t := time.Unix(seconds, 0)
	s := t.String()[:14] // change every hour
	seed := []byte(s)
	for {
		h := sha1.Sum([]byte(seed))
		seed = h[:]
		i := binary.BigEndian.Uint64(h[:8])
		r := int(i % uint64(len(q)))
		res = q[r]
		if allowRussianQuotes || !isRussian(res) {
			break
		}
	}
	return res
}

func processFlags() {
	if len(os.Args) > 1 {
		f := os.Args[1]
		allowRussianQuotes = strings.Contains(f, "r")
		if strings.Contains(f, "h") || strings.Contains(f, "?") {
			paint(createIgel(), getHelpText())
			os.Exit(0)
		}
	}
}

func paint(pic []string, lines []string) {
	numEmptyLines := (len(pic) - len(lines)) / 2
	if numEmptyLines > 3 {
		numEmptyLines--
	}

	cur := 0
	for j, ps := range pic {
		fmt.Printf(ps)
		if j >= numEmptyLines && cur < len(lines) {
			fmt.Printf(lines[cur])
			cur++
		}
		fmt.Println()
	}
}

func main() {
	processFlags()
	pic := createIgel()
	quote := getRandomQuote()
	lines := splitString(quote, 80)
	paint(pic, lines)
}
