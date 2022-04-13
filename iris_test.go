package iris

import "testing"

func TestIris(t *testing.T) {
	c := New(FgRed, BgWhite)
	c.Printf("Hello, %s\n", "Iris")
	Red("THIS IS A FATAL ERROR!!!11")
}
