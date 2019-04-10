// Testpakke for pakken ascii
package ascii

import "testing"
import "fmt"

// Testfunksjon for funksjonen GreetingsASCII.
func TestGreetingsASCII(t *testing.T) {
	ascii := GreetingASCII()
	for _, C := range ascii {
		if C > 127 {
			fmt.Println("Not only ascii! ")
			t.Fail()
		}

	}
}
