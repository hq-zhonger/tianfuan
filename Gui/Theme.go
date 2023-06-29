package Gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

// CustomTheme 主题定义
type CustomTheme struct {
	Theme string
}

func (m CustomTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch m.Theme {
	case "Dark":
		variant = theme.VariantDark
		switch name {
		case theme.ColorNameBackground:
			return color.NRGBA{R: 0x30, G: 0x30, B: 0x30, A: 0xff}
		case theme.ColorNameButton:
			return color.Transparent
		case theme.ColorNameDisabled:
			// return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x42}
			// return color.NRGBA{R: 19, G: 161, B: 14, A: 255}
			// 248,248,255
			return color.NRGBA{0xf8, 0xf8, 0xff, 0x66}
		case theme.ColorNameDisabledButton:
			return color.NRGBA{R: 0x26, G: 0x26, B: 0x26, A: 0xff}
		case theme.ColorNameError:
			return theme.ErrorColor()
		case theme.ColorNameForeground:
			return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
		case theme.ColorNameHover:
			return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x0f}
		case theme.ColorNameInputBackground:
			return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x19}
		case theme.ColorNamePlaceHolder:
			return color.NRGBA{R: 0xb2, G: 0xb2, B: 0xb2, A: 0xff}
		case theme.ColorNamePressed:
			return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x66}
		case theme.ColorNameScrollBar:
			return color.NRGBA{A: 0x99}
		case theme.ColorNameShadow:
			return color.NRGBA{A: 0x66}
		}

	case "Light":
		variant = theme.VariantLight
		switch name {
		case theme.ColorNameDisabled:
			return color.NRGBA{R: 0xab, G: 0xab, B: 0xab, A: 0xff}
			//case theme.ColorNameInputBorder:
			//	return color.NRGBA{R: 0xf3, G: 0xf3, B: 0xf3, A: 0xff}
		case theme.ColorNameDisabledButton:
			return color.NRGBA{R: 0xe5, G: 0xe5, B: 0xe5, A: 0xff}
		}
	}

	return theme.DefaultTheme().Color(name, variant)
}

func (m CustomTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m CustomTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m CustomTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
