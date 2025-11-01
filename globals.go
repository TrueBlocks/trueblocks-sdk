package sdk

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strings"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/v5/pkg/logger"
)

// Globals is a subset of globally available options from the command line
// that make sense in the SDK context
type Globals struct {
	Ether   bool   `json:"ether,omitempty"`
	Cache   bool   `json:"cache,omitempty"`
	Decache bool   `json:"decache,omitempty"`
	Verbose bool   `json:"verbose,omitempty"`
	Chain   string `json:"chain,omitempty"`
	Output  string `json:"output,omitempty"`
	Append  bool   `json:"append,omitempty"`
	// Probably can't support
	// --file
	// Global things ignored in the SDK
	// --help
	// --wei
	// --fmt
	// --version
	// --noop
	// --nocolor
	// --no_header
}

func (g Globals) String() string {
	bytes, _ := json.Marshal(g)
	return string(bytes)
}

type CacheOp uint8

const (
	CacheOn CacheOp = iota
	CacheOff
	Decache
)

func (g *Globals) Caching(op CacheOp) {
	switch op {
	case CacheOn:
		g.Cache = true
		g.Decache = false
	case CacheOff:
		g.Cache = false
		g.Decache = false
	case Decache:
		g.Cache = false
		g.Decache = true
	}
}

type Cacher interface {
	Caching(op CacheOp)
}

func convertObjectToArray(field, strIn string) string {
	// We’ll look for this pattern: `"field":`
	pattern := []byte("\"" + field + "\":")
	in := []byte(strIn)

	// Output buffer (pre-allocate at least len(in) to reduce re-allocation).
	out := make([]byte, 0, len(in))

	// Current scan position
	i := 0
	for {
		// Find the next occurrence of `"field":` starting from i.
		idx := bytes.Index(in[i:], pattern)
		if idx == -1 {
			// No more occurrences; copy everything left to output.
			out = append(out, in[i:]...)
			break
		}
		absIdx := i + idx // absolute index of the match in `in`

		// 1. Copy everything *before* the match to output.
		out = append(out, in[i:absIdx]...)

		// 2. Copy the `"field":` pattern itself to output.
		out = append(out, pattern...)
		// Advance i past `"field":`
		i = absIdx + len(pattern)

		// 3. Preserve any whitespace right after the colon in the output.
		for i < len(in) && isWhitespace(in[i]) {
			out = append(out, in[i])
			i++
		}

		// 4. If the next character is not `{`, nothing to bracket. Continue scanning.
		if i >= len(in) || in[i] != '{' {
			continue
		}

		// 5. Brace matching. If we find a matching `}`, we bracket the substring.
		braceStart := i
		braceCount := 0
		matchFound := false

		for j := i; j < len(in); j++ {
			if in[j] == '{' {
				braceCount++
			} else if in[j] == '}' {
				braceCount--
				if braceCount == 0 {
					// Found the matching closing brace at j
					// Write `[ { ... } ]` to output
					out = append(out, '[')
					out = append(out, in[braceStart:j+1]...)
					out = append(out, ']')

					// Advance i to the char after `}`
					i = j + 1
					matchFound = true
					break
				}
			}
		}

		// 6. If unmatched braces, copy the rest as-is and stop.
		if !matchFound {
			out = append(out, in[braceStart:]...)
			break
		}
	}

	return string(out)
}

func isWhitespace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\r'
}

func convertEmptyStrToZero(field, strIn string) string {
	convertToZero := func(field, str string) (string, bool) {
		find := "\"" + field + "\": \"\""
		start := strings.Index(str, find)
		if start == -1 {
			return str, false
		}

		end := start + len(find)
		for i := end; i < len(str); i++ {
			if str[i] == ',' || str[i] == '}' {
				end = i
				break
			}
		}

		beforeB := str[:start+len(find)-2] // Adjust to include '""'
		afterB := str[end:]                // after ","
		return beforeB + "\"0\"" + afterB, strings.Contains(str, find)
	}

	str := strIn
	for {
		var more bool
		str, more = convertToZero(field, str)
		if !more {
			break
		}
	}

	return str
}

func debugPrint(str string, t any, err error) {
	logger.Error("======================================")
	logger.Error(err)
	logger.Error(reflect.TypeOf(t))
	max := min(2000, len(str))
	logger.Error(str[:max])
	logger.Error("======================================")
}
