package termcolor

type RGB interface {
	RGB() (r, g, b uint8)
}

type Color int

const (
	BlackColor  Color = 1
	RedColor    Color = 2
	YellowColor Color = 3
	WhiteColor  Color = 4
	GreenColor  Color = 5
	PurpleColor Color = 6
	BlueColor   Color = 7
)

func (c Color) RGB() (r, g, b uint8) {
	switch c {
	case BlackColor:
		return 0, 0, 0
	case RedColor:
		return 255, 0, 0
	case YellowColor:
		return 255, 255, 0
	case WhiteColor:
		return 255, 255, 255
	case GreenColor:
		return 0, 255, 0
	case PurpleColor:
		return 160, 32, 240
	case BlueColor:
		return 0, 0, 255
	default:
		return 255, 255, 255
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

func FgRed(str string) string {
	r, g, b := RedColor.RGB()
	return FgRGB(str, r, g, b)
}

func BgBlack(str string) string {
	r, g, b := BlackColor.RGB()
	return BgRGB(str, r, g, b)
}

func FgBlack(str string) string {
	r, g, b := BlackColor.RGB()
	return FgRGB(str, r, g, b)
}

func BgYellow(str string) string {
	r, g, b := YellowColor.RGB()
	return BgRGB(str, r, g, b)
}

func FgYellow(str string) string {
	r, g, b := YellowColor.RGB()
	return FgRGB(str, r, g, b)
}

func BgWhite(str string) string {
	r, g, b := WhiteColor.RGB()
	return BgRGB(str, r, g, b)
}

func FgWhite(str string) string {
	r, g, b := WhiteColor.RGB()
	return FgRGB(str, r, g, b)
}

func BgGreen(str string) string {
	r, g, b := GreenColor.RGB()
	return BgRGB(str, r, g, b)
}

func FgGreen(str string) string {
	r, g, b := GreenColor.RGB()
	return FgRGB(str, r, g, b)
}

func BgPurple(str string) string {
	r, g, b := PurpleColor.RGB()
	return BgRGB(str, r, g, b)
}

func FgPurple(str string) string {
	r, g, b := PurpleColor.RGB()
	return FgRGB(str, r, g, b)
}

func BgBlue(str string) string {
	r, g, b := BlueColor.RGB()
	return BgRGB(str, r, g, b)
}

func FgBlue(str string) string {
	r, g, b := BlueColor.RGB()
	return FgRGB(str, r, g, b)
}
