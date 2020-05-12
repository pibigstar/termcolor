package termcolor

import (
	"fmt"
	"testing"
)

func Test0001(t *testing.T) {
	fmt.Println(FgRGB("aaa", 255, 0, 0))
	fmt.Println(BgRGB("aaa", 255, 0, 0))
	fmt.Println(FgRGB("aaa", 0, 0, 0))
	fmt.Println(BgRGB("aaa", 0, 0, 0))
}
