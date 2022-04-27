package mediao

import (
	"fmt"
	"testing"
)

func TestNetflix(*testing.T) {
	f := NewNetflixVerifier()
	v, err := f.Verify()
	fmt.Println("result info:", err, v)
}
