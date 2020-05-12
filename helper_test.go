package termcolor

import (
	"fmt"
	"testing"
)

func TestBgColor(t *testing.T) {
	fmt.Println(BgColor(RedColor, "sss" ))
}
