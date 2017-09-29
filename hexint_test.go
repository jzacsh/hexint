package hexint

import "testing"

var tt = []struct {
	Given   byte
	Expect  int
	IsValid bool
}{
	{'0', 0, true},
	// TODO(zacsh) fill out rest of typical cases
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
