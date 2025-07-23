package op

import "fmt"

// Font provides access to all Open Props font/typography variables
var Font = &font{}

type font struct{}

// Size returns a font size variable (--font-size-{0-8})
func (f *font) Size(n int) string {
	if n < 0 || n > 8 {
		n = clamp(n, 0, 8)
	}
	return cssVar(fmt.Sprintf("font-size-%d", n))
}

// SizeFluid returns a fluid font size variable (--font-size-fluid-{0-3})
func (f *font) SizeFluid(n int) string {
	if n < 0 || n > 3 {
		n = clamp(n, 0, 3)
	}
	return cssVar(fmt.Sprintf("font-size-fluid-%d", n))
}

// Weight returns a font weight variable (--font-weight-{1-9})
func (f *font) Weight(n int) string {
	if n < 1 || n > 9 {
		n = clamp(n, 1, 9)
	}
	return cssVar(fmt.Sprintf("font-weight-%d", n))
}

// LineHeight returns a line height variable (--font-lineheight-{0-5})
func (f *font) LineHeight(n int) string {
	if n == 0 {
		return cssVar("font-lineheight-00")
	}
	if n < 0 || n > 5 {
		n = clamp(n, 0, 5)
	}
	return cssVar(fmt.Sprintf("font-lineheight-%d", n))
}

// LetterSpacing returns a letter spacing variable (--font-letterspacing-{0-7})
func (f *font) LetterSpacing(n int) string {
	if n < 0 || n > 7 {
		n = clamp(n, 0, 7)
	}
	return cssVar(fmt.Sprintf("font-letterspacing-%d", n))
}

// Font family helpers
func (f *font) Sans() string {
	return cssVar("font-sans")
}

func (f *font) Serif() string {
	return cssVar("font-serif")
}

func (f *font) Mono() string {
	return cssVar("font-mono")
}

// Specific font family variables
func (f *font) SystemUI() string {
	return cssVar("font-system-ui")
}

func (f *font) Transitional() string {
	return cssVar("font-transitional")
}

func (f *font) OldStyle() string {
	return cssVar("font-old-style")
}

func (f *font) Humanist() string {
	return cssVar("font-humanist")
}

func (f *font) GeometricHumanist() string {
	return cssVar("font-geometric-humanist")
}

func (f *font) ClassicalHumanist() string {
	return cssVar("font-classical-humanist")
}

func (f *font) NeoGrotesque() string {
	return cssVar("font-neo-grotesque")
}

func (f *font) MonospaceSlab() string {
	return cssVar("font-monospace-slab-serif")
}

func (f *font) MonospaceCode() string {
	return cssVar("font-monospace-code")
}

func (f *font) Industrial() string {
	return cssVar("font-industrial")
}

func (f *font) RoundedSans() string {
	return cssVar("font-rounded-sans")
}

func (f *font) SlabSerif() string {
	return cssVar("font-slab-serif")
}

func (f *font) Antique() string {
	return cssVar("font-antique")
}

func (f *font) Didone() string {
	return cssVar("font-didone")
}

func (f *font) Handwritten() string {
	return cssVar("font-handwritten")
}

// Font size constants for common sizes
const (
	FontSizeXS   = 0 // 0.75rem
	FontSizeSM   = 1 // 1rem
	FontSizeMD   = 2 // 1.1rem
	FontSizeLG   = 3 // 1.25rem
	FontSizeXL   = 4 // 1.5rem
	FontSize2XL  = 5 // 2rem
	FontSize3XL  = 6 // 2.5rem
	FontSize4XL  = 7 // 3rem
	FontSize5XL  = 8 // 3.5rem
)

// Font weight constants
const (
	FontWeightThin       = 1 // 100
	FontWeightExtraLight = 2 // 200
	FontWeightLight      = 3 // 300
	FontWeightNormal     = 4 // 400
	FontWeightMedium     = 5 // 500
	FontWeightSemiBold   = 6 // 600
	FontWeightBold       = 7 // 700
	FontWeightExtraBold  = 8 // 800
	FontWeightBlack      = 9 // 900
)