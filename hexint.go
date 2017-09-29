package hexint

import "fmt"

func IsHex(rn byte) bool {
	_, e := DecodeInt(rn)
	return e == nil
}

func MustDecodeInt(hex byte) int {
	resp, e := DecodeInt(hex)
	if e != nil {
		panic(e)
	}
	return resp
}

// Given a hex char, returns its corresponding integer: a value in [0,15]
func DecodeInt(hex byte) (int, error) {
	// Char-to-rune table:
	//
	// | char    code point
	// +-------------------
	// | [0,9] = [48,57]
	// | [A,F] = [65,70]
	// | [a,f] = [97,102]

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
		return int(rn) - 65 + 10
		// TODO(zacsh): cannot comma-case these last two cases?
	default:
		return -1
	}
}
