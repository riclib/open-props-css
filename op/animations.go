package op

// Animation provides access to all Open Props animation variables
var Animation = &animation{}

type animation struct{}

// FadeIn returns the fade-in animation
func (a *animation) FadeIn() string {
	return cssVar("animation-fade-in")
}

// FadeInBloom returns the fade-in-bloom animation
func (a *animation) FadeInBloom() string {
	return cssVar("animation-fade-in-bloom")
}

// FadeOut returns the fade-out animation
func (a *animation) FadeOut() string {
	return cssVar("animation-fade-out")
}

// FadeOutBloom returns the fade-out-bloom animation
func (a *animation) FadeOutBloom() string {
	return cssVar("animation-fade-out-bloom")
}

// ScaleUp returns the scale-up animation
func (a *animation) ScaleUp() string {
	return cssVar("animation-scale-up")
}

// ScaleDown returns the scale-down animation
func (a *animation) ScaleDown() string {
	return cssVar("animation-scale-down")
}

// SlideOutUp returns the slide-out-up animation
func (a *animation) SlideOutUp() string {
	return cssVar("animation-slide-out-up")
}

// SlideOutDown returns the slide-out-down animation
func (a *animation) SlideOutDown() string {
	return cssVar("animation-slide-out-down")
}

// SlideOutRight returns the slide-out-right animation
func (a *animation) SlideOutRight() string {
	return cssVar("animation-slide-out-right")
}

// SlideOutLeft returns the slide-out-left animation
func (a *animation) SlideOutLeft() string {
	return cssVar("animation-slide-out-left")
}

// SlideInUp returns the slide-in-up animation
func (a *animation) SlideInUp() string {
	return cssVar("animation-slide-in-up")
}

// SlideInDown returns the slide-in-down animation
func (a *animation) SlideInDown() string {
	return cssVar("animation-slide-in-down")
}

// SlideInRight returns the slide-in-right animation
func (a *animation) SlideInRight() string {
	return cssVar("animation-slide-in-right")
}

// SlideInLeft returns the slide-in-left animation
func (a *animation) SlideInLeft() string {
	return cssVar("animation-slide-in-left")
}

// ShakeX returns the shake-x animation
func (a *animation) ShakeX() string {
	return cssVar("animation-shake-x")
}

// ShakeY returns the shake-y animation
func (a *animation) ShakeY() string {
	return cssVar("animation-shake-y")
}

// ShakeZ returns the shake-z animation
func (a *animation) ShakeZ() string {
	return cssVar("animation-shake-z")
}

// Spin returns the spin animation
func (a *animation) Spin() string {
	return cssVar("animation-spin")
}

// Ping returns the ping animation
func (a *animation) Ping() string {
	return cssVar("animation-ping")
}

// Blink returns the blink animation
func (a *animation) Blink() string {
	return cssVar("animation-blink")
}

// Float returns the float animation
func (a *animation) Float() string {
	return cssVar("animation-float")
}

// Bounce returns the bounce animation
func (a *animation) Bounce() string {
	return cssVar("animation-bounce")
}

// Pulse returns the pulse animation
func (a *animation) Pulse() string {
	return cssVar("animation-pulse")
}