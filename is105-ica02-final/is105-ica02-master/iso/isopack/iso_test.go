package iso

import "testing"

func TestGreetingsASCII(t *testing.T) {
	extASCII := GreetingExtendedASCII()
	for _, C := range extASCII {
		if C > 255 {
			t.Fail()
		}
	}
}
