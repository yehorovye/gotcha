/* simple config parser, by elisiei. published under CC0. */

package config

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func Parse(r io.Reader) (map[string]string, error) {
	cfg := make(map[string]string)
	sc := bufio.NewScanner(r)

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		cfg[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
	}

	return cfg, sc.Err()
}

func LoadFile(path string) (map[string]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return Parse(f)
}
