package hexint

import "testing"

var tt = []struct {
	Given  byte
	Expect int
	IsHex  bool
}{
	{'0', 0, true}, {'1', 1, true}, {'2', 2, true}, {'3', 3, true}, {'4', 4, true},
	{'5', 5, true}, {'6', 6, true}, {'7', 7, true}, {'8', 8, true}, {'9', 9, true},
	{'A', 10, true}, {'B', 11, true}, {'C', 12, true}, {'D', 13, true},
	{'E', 14, true}, {'F', 15, true},
	{'a', 10, true}, {'b', 11, true}, {'c', 12, true}, {'d', 13, true},
	{'e', 14, true}, {'f', 15, true},

	{'x', -1, false}, {' ', -1, false}, {'g', -1, false},
}

func Test_IsHex(t *testing.T) {
	for _, tc := range tt {
		t.Run(string(tc.Given), func(it *testing.T) {
			isHex := IsHex(tc.Given)
			if isHex == tc.IsHex {
				return
			}

			validity := "in"
			if tc.IsHex {
				validity = ""
			}
			it.Fatalf("expected '%c' to be %svalid hex", tc.Given, validity)
		})
	}
}

func Test_DecodeInt(t *testing.T) {
	for _, tc := range tt {
		t.Run(string(tc.Given), func(it *testing.T) {
			zint, e := DecodeInt(tc.Given)
			isHex := e == nil
			if tc.IsHex != isHex {
				validity := "in"
				if tc.IsHex {
					validity = ""
				}
				it.Fatalf("expected '%c' to be %svalid hex", tc.Given, validity)
			}

			if zint == tc.Expect {
				return
			}
			it.Fatalf("expected '%c' to decode to %d, but got %d", tc.Given, tc.Expect, zint)
		})
	}
}

func Test_MustDecodeInt(t *testing.T) {
	for _, tc := range tt {
		t.Run(string(tc.Given), func(it *testing.T) {
			defer func() {
				r := recover()
				wasPanicless := r == nil
				if tc.IsHex == wasPanicless {
					return
				}
				it.Fatalf("expected panic for '%c', but got non", tc.Given)
			}()

			zint := MustDecodeInt(tc.Given)

			if zint == tc.Expect {
				return
			}
			it.Fatalf("expected '%c' to decode to %d, but got %d", tc.Given, tc.Expect, zint)
		})
	}
}
