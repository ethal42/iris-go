package iris

import (
	"fmt"
	"testing"
)

func TestIris(t *testing.T) {
	c := New(Bold, FgMagenta)
	fmt.Println(c.Sprintf("this is test of %s", "Color.Sprintf()"))
}
