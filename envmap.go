package envmap

import (
	"os"
	"strings"
)

const separator = "="

// Envmap is a mapping of environment keys to values
type Envmap map[string]string

// Export exports the keys and values defined in this Envmap to the
// actual environment
func (e Envmap) Export() {

	for k, v := range e {
		os.Setenv(k, v)
	}

}

// Pop removes prefixes from all environment variable keys that match the
// filter. Matching keys that have no prefixes anymore, get dropped.
func (e Envmap) Pop(prefix string, filter func(string) bool) Envmap {

	var m Envmap = make(map[string]string)

	for k, v := range e {

		var (
			parts = strings.Split(k, prefix)
			depth = len(parts) - 1
			last  = parts[depth]
		)

		if filter(last) {

			if depth > 0 {
				m[k[1:len(k)]] = v
			}

		} else {
			m[k] = v
		}

	}

	return m

}

// Push prefixes all environment variable keys that match the filter
// with the given prefix
func (e Envmap) Push(prefix string, filter func(string) bool) Envmap {

	var m Envmap = make(map[string]string)

	for k, v := range e {

		var (
			parts = strings.Split(k, prefix)
			depth = len(parts) - 1
			last  = parts[depth]
		)

		if filter(last) {
			m[prefix+k] = v
		} else {
			m[k] = v
		}

	}

	return m

}

// ToEnv converts a map of environment variables to a slice
// of key=value strings
func (e Envmap) ToEnv() (env []string) {

	for k, v := range e {
		env = append(env, Join(k, v))
	}

	return

}

// Import creates an Envmap from the actual environment.
func Import() Envmap {
	return ToMap(os.Environ())
}

// Join builds a environment variable declaration out of seperate
// key and value strings
func Join(k, v string) string {
	return strings.Join([]string{k, v}, separator)
}

// ToMap converts a slice of environment variables to a map
// of environment variables
func ToMap(env []string) (m Envmap) {

	m = make(map[string]string)

	for _, e := range env {

		s := strings.Split(e, separator)

		var (
			key = s[0]
			val = strings.Join(s[1:], separator)
		)

		m[key] = val

	}

	return

}
