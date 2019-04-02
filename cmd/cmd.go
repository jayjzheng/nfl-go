package cmd

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"strings"
)

func PipedInput() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return true
	}

	return false
}

func ScanStdin() []byte {
	s := bufio.NewScanner(os.Stdin)

	var b []byte
	for s.Scan() {
		b = append(b, s.Bytes()...)
	}
	if s.Err() != nil {
		panic(s.Err())
	}
	return b
}

func StringSlice(s, sep string) []string {
	m := make(map[string]bool)

	for _, s := range strings.Split(s, sep) {
		k := strings.TrimSpace(s)
		if len(k) == 0 {
			continue
		}
		k = strings.ToUpper(k)
		m[k] = true
	}

	var ss []string
	for k := range m {
		ss = append(ss, k)
	}

	return ss
}

func WriteJSON(w io.Writer, v interface{}, pretty bool) error {
	e := json.NewEncoder(w)

	if pretty {
		e.SetIndent("", "\t")
	}

	return e.Encode(v)
}
