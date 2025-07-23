package op

import "fmt"

// Ease provides access to all Open Props easing variables
var Ease = &ease{}

type ease struct{}

// Default ease returns a default ease variable (--ease-{1-5})
func (e *ease) Default(n int) string {
	if n < 1 || n > 5 {
		n = clamp(n, 1, 5)
	}
	return cssVar(fmt.Sprintf("ease-%d", n))
}

// In returns an ease-in variable (--ease-in-{1-5})
func (e *ease) In(n int) string {
	if n < 1 || n > 5 {
		n = clamp(n, 1, 5)
	}
	return cssVar(fmt.Sprintf("ease-in-%d", n))
}

// Out returns an ease-out variable (--ease-out-{1-5})
func (e *ease) Out(n int) string {
	if n < 1 || n > 5 {
		n = clamp(n, 1, 5)
	}
	return cssVar(fmt.Sprintf("ease-out-%d", n))
}

// InOut returns an ease-in-out variable (--ease-in-out-{1-5})
func (e *ease) InOut(n int) string {
	if n < 1 || n > 5 {
		n = clamp(n, 1, 5)
	}
	return cssVar(fmt.Sprintf("ease-in-out-%d", n))
}

// Elastic returns an elastic ease variable (--ease-elastic-{1-5})
func (e *ease) Elastic(n int) string {
	if n < 1 || n > 5 {
		n = clamp(n, 1, 5)
	}
	return cssVar(fmt.Sprintf("ease-elastic-%d", n))
}

// ElasticOut returns an elastic-out ease variable (--ease-elastic-out-{1-5})
func (e *ease) ElasticOut(n int) string {
	if n < 1 || n > 5 {
		n = clamp(n, 1, 5)
	}
	return cssVar(fmt.Sprintf("ease-elastic-out-%d", n))
}

// ElasticIn returns an elastic-in ease variable (--ease-elastic-in-{1-5})
func (e *ease) ElasticIn(n int) string {
	if n < 1 || n > 5 {
		n = clamp(n, 1, 5)
	}
	return cssVar(fmt.Sprintf("ease-elastic-in-%d", n))
}

// ElasticInOut returns an elastic-in-out ease variable (--ease-elastic-in-out-{1-5})
func (e *ease) ElasticInOut(n int) string {
	if n < 1 || n > 5 {
		n = clamp(n, 1, 5)
	}
	return cssVar(fmt.Sprintf("ease-elastic-in-out-%d", n))
}

// Squish returns a squish ease variable (--ease-squish-{1-5})
func (e *ease) Squish(n int) string {
	if n < 1 || n > 5 {
		n = clamp(n, 1, 5)
	}
	return cssVar(fmt.Sprintf("ease-squish-%d", n))
}

// Step returns a step ease variable (--ease-step-{1-5})
func (e *ease) Step(n int) string {
	if n < 1 || n > 5 {
		n = clamp(n, 1, 5)
	}
	return cssVar(fmt.Sprintf("ease-step-%d", n))
}

// Spring returns a spring ease variable (--ease-spring-{1-5})
func (e *ease) Spring(n int) string {
	if n < 1 || n > 5 {
		n = clamp(n, 1, 5)
	}
	return cssVar(fmt.Sprintf("ease-spring-%d", n))
}

// Bounce returns a bounce ease variable (--ease-bounce-{1-5})
func (e *ease) Bounce(n int) string {
	if n < 1 || n > 5 {
		n = clamp(n, 1, 5)
	}
	return cssVar(fmt.Sprintf("ease-bounce-%d", n))
}

// Named easing functions
func (e *ease) CircIn() string {
	return cssVar("ease-circ-in")
}

func (e *ease) CircInOut() string {
	return cssVar("ease-circ-in-out")
}

func (e *ease) CircOut() string {
	return cssVar("ease-circ-out")
}

func (e *ease) CubicIn() string {
	return cssVar("ease-cubic-in")
}

func (e *ease) CubicInOut() string {
	return cssVar("ease-cubic-in-out")
}

func (e *ease) CubicOut() string {
	return cssVar("ease-cubic-out")
}

func (e *ease) ExpoIn() string {
	return cssVar("ease-expo-in")
}

func (e *ease) ExpoInOut() string {
	return cssVar("ease-expo-in-out")
}

func (e *ease) ExpoOut() string {
	return cssVar("ease-expo-out")
}

func (e *ease) QuadIn() string {
	return cssVar("ease-quad-in")
}

func (e *ease) QuadInOut() string {
	return cssVar("ease-quad-in-out")
}

func (e *ease) QuadOut() string {
	return cssVar("ease-quad-out")
}

func (e *ease) QuartIn() string {
	return cssVar("ease-quart-in")
}

func (e *ease) QuartInOut() string {
	return cssVar("ease-quart-in-out")
}

func (e *ease) QuartOut() string {
	return cssVar("ease-quart-out")
}

func (e *ease) QuintIn() string {
	return cssVar("ease-quint-in")
}

func (e *ease) QuintInOut() string {
	return cssVar("ease-quint-in-out")
}

func (e *ease) QuintOut() string {
	return cssVar("ease-quint-out")
}

func (e *ease) SineIn() string {
	return cssVar("ease-sine-in")
}

func (e *ease) SineInOut() string {
	return cssVar("ease-sine-in-out")
}

func (e *ease) SineOut() string {
	return cssVar("ease-sine-out")
}