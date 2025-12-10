package monitor

import (
	"strings"
)

type TemplateVars struct {
	Address    string
	Chain      string
	FirstBlock uint64
	LastBlock  uint64
	BlockCount uint64
}

func ExpandTemplate(input string, vars TemplateVars) string {
	result := input
	result = strings.ReplaceAll(result, "{address}", vars.Address)
	result = strings.ReplaceAll(result, "{chain}", vars.Chain)
	result = strings.ReplaceAll(result, "{first_block}", formatUint64(vars.FirstBlock))
	result = strings.ReplaceAll(result, "{last_block}", formatUint64(vars.LastBlock))
	result = strings.ReplaceAll(result, "{block_count}", formatUint64(vars.BlockCount))
	return result
}

func formatUint64(n uint64) string {
	if n == 0 {
		return "0"
	}
	var result strings.Builder
	str := []byte{}
	for n > 0 {
		str = append(str, byte('0'+n%10))
		n /= 10
	}
	for i := len(str) - 1; i >= 0; i-- {
		result.WriteByte(str[i])
	}
	return result.String()
}
