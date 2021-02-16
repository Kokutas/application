/**
 * @Author Kokutas
 * @Description //TODO
 * @Date 2021/2/15 0:40
 **/
package themes

import (
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2"
	"image/color"
)

type myTheme struct{
	variant fyne.ThemeVariant
	regular  fyne.Resource
}
var _ fyne.Theme = (*myTheme)(nil)

// return bundled font resource
func (*myTheme) Font(s fyne.TextStyle) fyne.Resource {
	if s.Monospace {
		return theme.DefaultTheme().Font(s)
	}
	if s.Bold {
		if s.Italic {
			return theme.DefaultTheme().Font(s)
		}
		return regular
	}
	if s.Italic {
		return theme.DefaultTheme().Font(s)
	}
	return regular
}

func (*myTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	colors := darkPalette

	if n == ColorNamePrimary {
		return PrimaryColorNamed(fyne.CurrentApp().Settings().PrimaryColor())
	} else if n == ColorNameFocus {
		return focusColorNamed(fyne.CurrentApp().Settings().PrimaryColor())
	}

	if c, ok := colors[n]; ok {
		return c
	}
	return color.Transparent
	return theme.DefaultTheme().Color(n, v)
}

func (*myTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (*myTheme) Size(n fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(n)
}
const (
	ColorRed = "red"
	ColorOrange = "orange"
	ColorYellow = "yellow"
	ColorGreen = "green"
	ColorBlue = "blue"
	ColorPurple = "purple"
	ColorBrown = "brown"
	ColorGray = "gray"
	ColorNameBackground fyne.ThemeColorName = "background"
	ColorNameButton fyne.ThemeColorName = "button"
	ColorNameDisabledButton fyne.ThemeColorName = "disabledButton"
	ColorNameDisabled fyne.ThemeColorName = "disabled"
	ColorNameError fyne.ThemeColorName = "error"
	ColorNameFocus fyne.ThemeColorName = "focus"
	ColorNameForeground fyne.ThemeColorName = "foreground"
	ColorNameHover fyne.ThemeColorName = "hover"
	ColorNameInputBackground fyne.ThemeColorName = "inputBackground"
	ColorNamePlaceHolder fyne.ThemeColorName = "placeholder"
	ColorNamePressed fyne.ThemeColorName = "pressed"
	ColorNamePrimary fyne.ThemeColorName = "primary"
	ColorNameScrollBar fyne.ThemeColorName = "scrollBar"
	ColorNameShadow fyne.ThemeColorName = "shadow"
	VariantDark fyne.ThemeVariant = 0
)

var(
	errorColor  = color.NRGBA{0xf4, 0x43, 0x36, 0xff}
	darkPalette = map[fyne.ThemeColorName]color.Color{
		ColorNameBackground:      color.NRGBA{0x30, 0x30, 0x30, 0xff},
		ColorNameButton:          color.Transparent,
		ColorNameDisabled:        color.NRGBA{0xff, 0xff, 0xff, 0x42},
		ColorNameDisabledButton:  color.NRGBA{0x26, 0x26, 0x26, 0xff},
		ColorNameError:           errorColor,
		ColorNameForeground:      color.NRGBA{0xff, 0xff, 0xff, 0xff},
		ColorNameHover:           color.NRGBA{0xff, 0xff, 0xff, 0x0f},
		ColorNameInputBackground: color.NRGBA{0xff, 0xff, 0xff, 0x19},
		ColorNamePlaceHolder:     color.NRGBA{0xb2, 0xb2, 0xb2, 0xff},
		ColorNamePressed:         color.NRGBA{0xff, 0xff, 0xff, 0x66},
		ColorNameScrollBar:       color.NRGBA{0x0, 0x0, 0x0, 0x99},
		ColorNameShadow:          color.NRGBA{0x0, 0x0, 0x0, 0x66},
	}
	primaryColors = map[string]color.Color{
		ColorRed:    color.NRGBA{R: 0xf4, G: 0x43, B: 0x36, A: 0xff},
		ColorOrange: color.NRGBA{R: 0xff, G: 0x98, B: 0x00, A: 0xff},
		ColorYellow: color.NRGBA{R: 0xff, G: 0xeb, B: 0x3b, A: 0xff},
		ColorGreen:  color.NRGBA{R: 0x8b, G: 0xc3, B: 0x4a, A: 0xff},
		ColorBlue:   color.NRGBA{R: 0x21, G: 0x96, B: 0xf3, A: 0xff},
		ColorPurple: color.NRGBA{R: 0x9c, G: 0x27, B: 0xb0, A: 0xff},
		ColorBrown:  color.NRGBA{R: 0x79, G: 0x55, B: 0x48, A: 0xff},
		ColorGray:   color.NRGBA{R: 0x9e, G: 0x9e, B: 0x9e, A: 0xff},
	}
	focusColors = map[string]color.Color{
		ColorRed:    color.NRGBA{R: 0xf4, G: 0x43, B: 0x36, A: 0x7f},
		ColorOrange: color.NRGBA{R: 0xff, G: 0x98, B: 0x00, A: 0x7f},
		ColorYellow: color.NRGBA{R: 0xff, G: 0xeb, B: 0x3b, A: 0x7f},
		ColorGreen:  color.NRGBA{R: 0x8b, G: 0xc3, B: 0x4a, A: 0x7f},
		ColorBlue:   color.NRGBA{R: 0x21, G: 0x96, B: 0xf3, A: 0x7f},
		ColorPurple: color.NRGBA{R: 0x9c, G: 0x27, B: 0xb0, A: 0x7f},
		ColorBrown:  color.NRGBA{R: 0x79, G: 0x55, B: 0x48, A: 0x7f},
		ColorGray:   color.NRGBA{R: 0x9e, G: 0x9e, B: 0x9e, A: 0x7f},
	}
)

func DarkTheme() fyne.Theme {
	theme := &myTheme{variant: VariantDark,regular: regular}
	return theme
}
func PrimaryColorNamed(name string) color.Color {
	col, ok := primaryColors[name]
	if !ok {
		return primaryColors[ColorBlue]
	}
	return col
}

func focusColorNamed(name string) color.Color {
	col, ok := focusColors[name]
	if !ok {
		return focusColors[ColorBlue]
	}
	return col
}