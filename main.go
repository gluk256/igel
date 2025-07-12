package main

import (
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"os"
	"strings"
	"time"
)

const textFieldSize = 80

var allowRussian, testMode bool

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

func splitString(s string) []string {
	maxLength := textFieldSize
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

func getRandomAphorism() (res string) {
	arr := getAphorismList()
	seconds := time.Now().Unix()
	t := time.Unix(seconds, 0)
	s := t.String()
	if !testMode {
		s = s[:14] // change every hour
	}
	seed := []byte(s)
	for {
		hash := sha1.Sum([]byte(seed))
		seed = hash[:]
		i := binary.BigEndian.Uint64(hash[:8])
		rnd := int(i % uint64(len(arr)))
		res = arr[rnd]
		if allowRussian || !isRussian(res) {
			break
		}
	}
	return res
}

func processFlags() {
	if len(os.Args) > 1 {
		f := os.Args[1]
		testMode = strings.Contains(f, "t")
		allowRussian = strings.Contains(f, "r") // todo: fix before commit

		if strings.Contains(f, "d") {
			runDebug()
			os.Exit(0)
		}
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
	for j, p := range pic {
		var text string
		if j >= numEmptyLines && cur < len(lines) {
			text = lines[cur]
			cur++
		}
		fmt.Println(p + text)
	}
}

func main() {
	processFlags()
	pic := createIgel()
	quote := getRandomAphorism()
	lines := splitString(quote)
	paint(pic, lines)
}

func runDebug() {
	arr := getAphorismList()
	quote := arr[54] // (line_num - 6) / 2
	fmt.Println("running in debug mode")
	lines := splitString(quote)
	paint(createIgel(), lines)
}
