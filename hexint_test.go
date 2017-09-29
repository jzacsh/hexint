package hexint

import "testing"

var tt = []struct {
	Given   byte
	Expect  int
	IsValid bool
}{
	{'0', 0, true}, {'1', 1, true}, {'2', 2, true}, {'3', 3, true}, {'4', 4, true},
	{'5', 5, true}, {'6', 6, true}, {'7', 7, true}, {'8', 8, true}, {'9', 9, true},
	{'A', 10, true}, {'B', 11, true}, {'C', 12, true}, {'D', 13, true},
	{'E', 14, true}, {'F', 15, true},
	{'a', 10, true}, {'b', 11, true}, {'c', 12, true}, {'d', 13, true},
	{'e', 14, true}, {'f', 15, true},

	{'x', -1, false}, {' ', -1, false}, {'g', -1, false},
}

func Test_IsValidHex(t *testing.T) {
	for _, tc := range tt {
		t.Run(string(tc.Given), func(it *testing.T) {
			isValid := IsHex(tc.Given)
			if isValid != tc.IsValid {
				validity := "in"
				if tc.IsValid {
					validity = ""
				}
				it.Fatalf("expected '%c' to be %svalid", tc.Given, validity)
			}
		})
	}
}

func Test_MustParseInt(t *testing.T) {
	t.Fatalf("not yet tested") // TODO(zacsh)
}

func Test_ParseInt(t *testing.T) {
	t.Fatalf("not yet tested") // TODO(zacsh)
}
