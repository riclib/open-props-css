// Package op provides a type-safe Go API for Open Props CSS variables.
// It allows easy access to all Open Props design tokens through a clean, idiomatic Go interface.
package op

import (
	"fmt"
	"strings"
)

// cssVar returns a CSS variable reference string
func cssVar(name string) string {
	return fmt.Sprintf("var(--%s)", name)
}

// clamp ensures a value is within min and max bounds
func clamp(val, min, max int) int {
	if val < min {
		return min
	}
	if val > max {
		return max
	}
	return val
}

// Size returns a size variable (--size-{n})
// Valid range: 1-15 (also supports negative values 000, 00)
func Size(n int) string {
	switch n {
	case -2:
		return cssVar("size-000")
	case -1:
		return cssVar("size-00")
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
		return cssVar(fmt.Sprintf("size-%d", n))
	default:
		// Clamp to valid range
		clamped := clamp(n, 1, 15)
		return cssVar(fmt.Sprintf("size-%d", clamped))
	}
}

// SizePx returns a pixel-based size variable (--size-px-{n})
func SizePx(n int) string {
	switch n {
	case -2:
		return cssVar("size-px-000")
	case -1:
		return cssVar("size-px-00")
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
		return cssVar(fmt.Sprintf("size-px-%d", n))
	default:
		clamped := clamp(n, 1, 15)
		return cssVar(fmt.Sprintf("size-px-%d", clamped))
	}
}

// SizeFluid returns a fluid size variable (--size-fluid-{n})
func SizeFluid(n int) string {
	if n < 1 || n > 10 {
		n = clamp(n, 1, 10)
	}
	return cssVar(fmt.Sprintf("size-fluid-%d", n))
}

// SizeContent returns a content size variable (--size-content-{n})
func SizeContent(n int) string {
	if n < 1 || n > 3 {
		n = clamp(n, 1, 3)
	}
	return cssVar(fmt.Sprintf("size-content-%d", n))
}

// SizeHeader returns a header size variable (--size-header-{n})
func SizeHeader(n int) string {
	if n < 1 || n > 3 {
		n = clamp(n, 1, 3)
	}
	return cssVar(fmt.Sprintf("size-header-%d", n))
}

// Shadow returns a shadow variable (--shadow-{n})
func Shadow(n int) string {
	if n < 1 || n > 6 {
		n = clamp(n, 1, 6)
	}
	return cssVar(fmt.Sprintf("shadow-%d", n))
}

// InnerShadow returns an inner shadow variable (--inner-shadow-{n})
func InnerShadow(n int) string {
	if n < 0 || n > 4 {
		n = clamp(n, 0, 4)
	}
	return cssVar(fmt.Sprintf("inner-shadow-%d", n))
}

// Gradient returns a gradient variable (--gradient-{n})
func Gradient(n int) string {
	if n < 1 || n > 30 {
		n = clamp(n, 1, 30)
	}
	return cssVar(fmt.Sprintf("gradient-%d", n))
}

// Layer returns a layer/z-index variable (--layer-{n})
func Layer(n int) string {
	if n < 1 || n > 5 {
		n = clamp(n, 1, 5)
	}
	return cssVar(fmt.Sprintf("layer-%d", n))
}

// LayerImportant returns the important layer variable
func LayerImportant() string {
	return cssVar("layer-important")
}

// Ratio returns an aspect ratio variable
func Ratio(name string) string {
	validRatios := map[string]bool{
		"square":     true,
		"landscape":  true,
		"portrait":   true,
		"widescreen": true,
		"ultrawide":  true,
		"golden":     true,
	}
	
	if !validRatios[name] {
		return cssVar("ratio-square") // default
	}
	return cssVar(fmt.Sprintf("ratio-%s", name))
}

// Style provides a builder pattern for creating CSS styles
type Style struct {
	props []string
}

// NewStyle creates a new style builder
func NewStyle() *Style {
	return &Style{
		props: make([]string, 0),
	}
}

// Background adds a background property
func (s *Style) Background(value string) *Style {
	s.props = append(s.props, fmt.Sprintf("background: %s", value))
	return s
}

// Color adds a color property
func (s *Style) Color(value string) *Style {
	s.props = append(s.props, fmt.Sprintf("color: %s", value))
	return s
}

// Padding adds a padding property
func (s *Style) Padding(value string) *Style {
	s.props = append(s.props, fmt.Sprintf("padding: %s", value))
	return s
}

// Margin adds a margin property
func (s *Style) Margin(value string) *Style {
	s.props = append(s.props, fmt.Sprintf("margin: %s", value))
	return s
}

// BorderRadius adds a border-radius property
func (s *Style) BorderRadius(value string) *Style {
	s.props = append(s.props, fmt.Sprintf("border-radius: %s", value))
	return s
}

// BoxShadow adds a box-shadow property
func (s *Style) BoxShadow(value string) *Style {
	s.props = append(s.props, fmt.Sprintf("box-shadow: %s", value))
	return s
}

// Animation adds an animation property
func (s *Style) Animation(value string) *Style {
	s.props = append(s.props, fmt.Sprintf("animation: %s", value))
	return s
}

// FontSize adds a font-size property
func (s *Style) FontSize(value string) *Style {
	s.props = append(s.props, fmt.Sprintf("font-size: %s", value))
	return s
}

// FontWeight adds a font-weight property
func (s *Style) FontWeight(value string) *Style {
	s.props = append(s.props, fmt.Sprintf("font-weight: %s", value))
	return s
}

// Custom adds a custom CSS property
func (s *Style) Custom(property, value string) *Style {
	s.props = append(s.props, fmt.Sprintf("%s: %s", property, value))
	return s
}

// String returns the CSS string
func (s *Style) String() string {
	return strings.Join(s.props, "; ")
}

// Common preset sizes
const (
	SizeXS  = 1  // 0.25rem
	SizeSM  = 2  // 0.5rem
	SizeMD  = 3  // 1rem
	SizeLG  = 4  // 1.25rem
	SizeXL  = 5  // 1.5rem
	Size2XL = 6  // 1.75rem
	Size3XL = 7  // 2rem
	Size4XL = 8  // 3rem
	Size5XL = 9  // 4rem
	Size6XL = 10 // 5rem
)