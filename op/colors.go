package op

import "fmt"

// Colors provides access to all Open Props color variables
var Color = &colors{}

type colors struct{}

// colorScale generates a color variable for a given color name and scale (0-12)
func colorScale(name string, scale int) string {
	if scale < 0 || scale > 12 {
		scale = clamp(scale, 0, 12)
	}
	return cssVar(fmt.Sprintf("%s-%d", name, scale))
}

// Gray returns a gray color variable (--gray-{0-12})
func (c *colors) Gray(scale int) string {
	return colorScale("gray", scale)
}

// Stone returns a stone color variable (--stone-{0-12})
func (c *colors) Stone(scale int) string {
	return colorScale("stone", scale)
}

// Red returns a red color variable (--red-{0-12})
func (c *colors) Red(scale int) string {
	return colorScale("red", scale)
}

// Pink returns a pink color variable (--pink-{0-12})
func (c *colors) Pink(scale int) string {
	return colorScale("pink", scale)
}

// Purple returns a purple color variable (--purple-{0-12})
func (c *colors) Purple(scale int) string {
	return colorScale("purple", scale)
}

// Violet returns a violet color variable (--violet-{0-12})
func (c *colors) Violet(scale int) string {
	return colorScale("violet", scale)
}

// Indigo returns an indigo color variable (--indigo-{0-12})
func (c *colors) Indigo(scale int) string {
	return colorScale("indigo", scale)
}

// Blue returns a blue color variable (--blue-{0-12})
func (c *colors) Blue(scale int) string {
	return colorScale("blue", scale)
}

// Cyan returns a cyan color variable (--cyan-{0-12})
func (c *colors) Cyan(scale int) string {
	return colorScale("cyan", scale)
}

// Teal returns a teal color variable (--teal-{0-12})
func (c *colors) Teal(scale int) string {
	return colorScale("teal", scale)
}

// Green returns a green color variable (--green-{0-12})
func (c *colors) Green(scale int) string {
	return colorScale("green", scale)
}

// Lime returns a lime color variable (--lime-{0-12})
func (c *colors) Lime(scale int) string {
	return colorScale("lime", scale)
}

// Yellow returns a yellow color variable (--yellow-{0-12})
func (c *colors) Yellow(scale int) string {
	return colorScale("yellow", scale)
}

// Orange returns an orange color variable (--orange-{0-12})
func (c *colors) Orange(scale int) string {
	return colorScale("orange", scale)
}

// Choco returns a choco color variable (--choco-{0-12})
func (c *colors) Choco(scale int) string {
	return colorScale("choco", scale)
}

// Brown returns a brown color variable (--brown-{0-12})
func (c *colors) Brown(scale int) string {
	return colorScale("brown", scale)
}

// Sand returns a sand color variable (--sand-{0-12})
func (c *colors) Sand(scale int) string {
	return colorScale("sand", scale)
}

// Camo returns a camo color variable (--camo-{0-12})
func (c *colors) Camo(scale int) string {
	return colorScale("camo", scale)
}

// Jungle returns a jungle color variable (--jungle-{0-12})
func (c *colors) Jungle(scale int) string {
	return colorScale("jungle", scale)
}

// Theme colors - these are custom to our framework
func (c *colors) Background() string {
	return cssVar("background")
}

func (c *colors) Surface() string {
	return cssVar("surface")
}

func (c *colors) SurfaceAlt() string {
	return cssVar("surface-alt")
}

func (c *colors) Text() string {
	return cssVar("text")
}

func (c *colors) TextMuted() string {
	return cssVar("text-muted")
}

func (c *colors) Border() string {
	return cssVar("border")
}

func (c *colors) Primary() string {
	return cssVar("primary")
}

func (c *colors) PrimaryHover() string {
	return cssVar("primary-hover")
}

// Common color scale values as constants for convenience
const (
	ColorScaleLightest = 0
	ColorScaleLight    = 3
	ColorScaleMedium   = 6
	ColorScaleDark     = 9
	ColorScaleDarkest  = 12
)