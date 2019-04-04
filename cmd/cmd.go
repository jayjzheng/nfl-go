package cmd

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"strings"
)

func WriteJSON(w io.Writer, v interface{}, pretty bool) error {
	e := json.NewEncoder(w)

	if pretty {
		e.SetIndent("", "\t")
	}

	return e.Encode(v)
}

func SplitInput(s, sep string) ([]string, error) {
	m := make(map[string]bool)

	piped, err := pipedInput()
	if err != nil {
		return nil, err
	}
	if piped {
		for _, l := range scanStdin() {
			split(l, sep, m)
		}
	}
	split(s, sep, m)

	var ss []string
	for k := range m {
		ss = append(ss, k)
	}

	return ss, nil
}

func split(s, sep string, m map[string]bool) {
	for _, s := range strings.Split(s, sep) {
		k := strings.TrimSpace(s)
		if len(k) == 0 {
			continue
		}
		m[k] = true
	}
}

func pipedInput() (bool, error) {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false, err
	}

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return true, nil
	}

	return false, nil
}

func scanStdin() (lines []string) {
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		lines = append(lines, string(s.Bytes()))
	}
	if s.Err() != nil {
		panic(s.Err())
	}
	return lines
}
