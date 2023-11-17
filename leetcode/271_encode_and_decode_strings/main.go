package encodeanddecodestrings

import (
	"regexp"
	"strings"
)

type Codec struct {
}

func (codec *Codec) Encode(strs []string) string {
	sep := "\n"
	return strings.Join(strs, sep)
}

func (codec *Codec) Decode(strs string) []string {
	r := regexp.MustCompile("\n")
	return r.Split(strs, -1)
}
