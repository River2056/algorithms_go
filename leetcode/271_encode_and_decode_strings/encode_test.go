package encodeanddecodestrings

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	input := []string{"C#", "&", "~Xp|F", "R4QBf9g=_"}
	codec := Codec{}
	encode := codec.Encode(input)
	fmt.Println(encode)
	decode := codec.Decode(encode)
	fmt.Println(decode)
}
