package hexint

import "fmt"

// IsHex determines if the given rune, `rn` is a hex character that can be
// decoded by `DecodeInt`.
func IsHex(rn byte) bool {
	_, e := DecodeInt(rn)
	return e == nil
}

// MustDecodeInt thin panic-wrapper for DecodeInt().
func MustDecodeInt(hex byte) int {
	resp, e := DecodeInt(hex)
	if e != nil {
		panic(e)
	}
	return resp
}

// DecodeInt returns an integer in [0,15] given `hex` byte as extracted from a
// string, case-insensitive. Specifically, follows this table mapping chars to
// rune-codepoints:
//   char  | code point | decoded to
//  --------------------------------
//   [0,9] |  [48,57]   |  [0,9]
//   [A,F] |  [65,70]   | [10,15]
//   [a,f] |  [97,102]  | [10,15]
func DecodeInt(hex byte) (int, error) {
	runePoint := rune(hex)
	val := decodeInt(runePoint)
	if val == -1 {
		return -1, fmt.Errorf("invalid hex, got: '%c'", runePoint)
	}
	return val, nil // TODO(zacsh) finish converting to int [0,15]
}

func decodeInt(rn rune) int {
	switch {
	case (rn >= 48 && rn <= 57): /*[0,9]*/
		return int(rn) - 48
	case (rn >= 65 && rn <= 70): /*[A,F]*/
		return int(rn) - 65 + 10
	case (rn >= 97 && rn <= 102): /*[a,f]*/
		return int(rn) - 97 + 10
		// TODO(zacsh): cannot comma-case these last two cases?
	default:
		return -1
	}
}
