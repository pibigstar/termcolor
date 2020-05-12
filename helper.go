package termcolor

type RGB interface {
	RGB()(r,g,b uint8)
}

type Color int

const (
	BlackColor Color = 1
	RedColor  Color = 2
)

func (c Color) RGB()(r,g,b uint8) {
	switch c {
	case BlackColor:
		return 0,0,0
	case RedColor:
		return 255,0,0
	default:
		return 255,255,255
	}
}

func BgColor(color RGB, str string) string {
	r, g, b := color.RGB()
	return BgRGB(str, r, g, b)
}

func FgColor(color RGB, str string) string {
	r, g, b := color.RGB()
	return FgRGB(str, r, g, b)
}

func BgRed(str string) string {
	r, g, b := RedColor.RGB()
	return BgRGB(str, r, g, b)
}

func BgBlack(str string) string {
	r, g, b := BlackColor.RGB()
	return BgRGB(str, r, g, b)
}