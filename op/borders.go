package op

import "fmt"

// Border provides access to border-related variables
var Border = &border{}

type border struct{}

// Size returns a border size variable (--border-size-{1-5})
func (b *border) Size(n int) string {
	if n < 1 || n > 5 {
		n = clamp(n, 1, 5)
	}
	return cssVar(fmt.Sprintf("border-size-%d", n))
}

// Radius returns a radius variable (--radius-{1-6})
func Radius(n int) string {
	if n < 1 || n > 6 {
		n = clamp(n, 1, 6)
	}
	return cssVar(fmt.Sprintf("radius-%d", n))
}

// RadiusRound returns the round radius variable
func RadiusRound() string {
	return cssVar("radius-round")
}

// RadiusBlob returns a blob radius variable (--radius-blob-{1-5})
func RadiusBlob(n int) string {
	if n < 1 || n > 5 {
		n = clamp(n, 1, 5)
	}
	return cssVar(fmt.Sprintf("radius-blob-%d", n))
}

// RadiusConditional returns a conditional radius variable (--radius-conditional-{1-6})
func RadiusConditional(n int) string {
	if n < 1 || n > 6 {
		n = clamp(n, 1, 6)
	}
	return cssVar(fmt.Sprintf("radius-conditional-%d", n))
}

// Border size constants
const (
	BorderSizeThin    = 1 // 1px
	BorderSizeDefault = 2 // 2px
	BorderSizeMedium  = 3 // 5px
	BorderSizeThick   = 4 // 10px
	BorderSizeHeavy   = 5 // 25px
)

// Radius constants
const (
	RadiusXS     = 1 // 2px
	RadiusSM     = 2 // 5px
	RadiusMD     = 3 // 1rem
	RadiusLG     = 4 // 2rem
	RadiusXL     = 5 // 4rem
	Radius2XL    = 6 // 8rem
)