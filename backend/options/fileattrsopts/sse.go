package fileattrsopts

import (
	"strings"

	"github.com/c2fo/vfs/v6"
)

const OptionNameSSE = "SSE"

// WithSSE allows 'NONE' or 'AES256' (default)
func WithSSE(algo string) vfs.FileAttrsOption {
	if !strings.EqualFold(algo, "NONE") {
		algo = "AES256"
	}
	return SSE{
		algorithm: algo,
	}
}

type SSE struct {
	algorithm string
}

func (s SSE) FileAttrsOptionName() string {
	return OptionNameSSE
}

func (s SSE) GetAlgorithm() string {
	return s.algorithm
}
