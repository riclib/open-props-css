package op

import (
	"strings"
	"testing"
)

func TestCssVar(t *testing.T) {
	result := cssVar("test-var")
	expected := "var(--test-var)"
	if result != expected {
		t.Errorf("cssVar() = %v, want %v", result, expected)
	}
}

func TestSize(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{-2, "var(--size-000)"},
		{-1, "var(--size-00)"},
		{1, "var(--size-1)"},
		{5, "var(--size-5)"},
		{15, "var(--size-15)"},
		{20, "var(--size-15)"}, // Should clamp to 15
	}

	for _, tt := range tests {
		result := Size(tt.input)
		if result != tt.expected {
			t.Errorf("Size(%d) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestColorGray(t *testing.T) {
	tests := []struct {
		scale    int
		expected string
	}{
		{0, "var(--gray-0)"},
		{6, "var(--gray-6)"},
		{12, "var(--gray-12)"},
		{15, "var(--gray-12)"}, // Should clamp to 12
	}

	for _, tt := range tests {
		result := Color.Gray(tt.scale)
		if result != tt.expected {
			t.Errorf("Color.Gray(%d) = %v, want %v", tt.scale, result, tt.expected)
		}
	}
}

func TestFontSize(t *testing.T) {
	tests := []struct {
		size     int
		expected string
	}{
		{0, "var(--font-size-0)"},
		{3, "var(--font-size-3)"},
		{8, "var(--font-size-8)"},
		{10, "var(--font-size-8)"}, // Should clamp to 8
	}

	for _, tt := range tests {
		result := Font.Size(tt.size)
		if result != tt.expected {
			t.Errorf("Font.Size(%d) = %v, want %v", tt.size, result, tt.expected)
		}
	}
}

func TestAnimation(t *testing.T) {
	tests := []struct {
		name     string
		fn       func() string
		expected string
	}{
		{"FadeIn", Animation.FadeIn, "var(--animation-fade-in)"},
		{"Bounce", Animation.Bounce, "var(--animation-bounce)"},
		{"SlideInUp", Animation.SlideInUp, "var(--animation-slide-in-up)"},
		{"Pulse", Animation.Pulse, "var(--animation-pulse)"},
	}

	for _, tt := range tests {
		result := tt.fn()
		if result != tt.expected {
			t.Errorf("Animation.%s() = %v, want %v", tt.name, result, tt.expected)
		}
	}
}

func TestGradient(t *testing.T) {
	tests := []struct {
		n        int
		expected string
	}{
		{1, "var(--gradient-1)"},
		{15, "var(--gradient-15)"},
		{30, "var(--gradient-30)"},
		{35, "var(--gradient-30)"}, // Should clamp to 30
	}

	for _, tt := range tests {
		result := Gradient(tt.n)
		if result != tt.expected {
			t.Errorf("Gradient(%d) = %v, want %v", tt.n, result, tt.expected)
		}
	}
}

func TestStyleBuilder(t *testing.T) {
	style := NewStyle().
		Background(Gradient(7)).
		Color(Color.Indigo(9)).
		Padding(Size(4)).
		BorderRadius(Radius(3)).
		BoxShadow(Shadow(3)).
		Animation(Animation.FadeIn()).
		FontSize(Font.Size(3)).
		FontWeight(Font.Weight(6))

	result := style.String()

	// Check that all properties are included
	expectedParts := []string{
		"background: var(--gradient-7)",
		"color: var(--indigo-9)",
		"padding: var(--size-4)",
		"border-radius: var(--radius-3)",
		"box-shadow: var(--shadow-3)",
		"animation: var(--animation-fade-in)",
		"font-size: var(--font-size-3)",
		"font-weight: var(--font-weight-6)",
	}

	for _, part := range expectedParts {
		if !strings.Contains(result, part) {
			t.Errorf("Style.String() missing %v", part)
		}
	}

	// Check the overall format
	if !strings.Contains(result, "; ") {
		t.Error("Style.String() should separate properties with '; '")
	}
}

func TestBorderSize(t *testing.T) {
	tests := []struct {
		size     int
		expected string
	}{
		{1, "var(--border-size-1)"},
		{3, "var(--border-size-3)"},
		{5, "var(--border-size-5)"},
		{10, "var(--border-size-5)"}, // Should clamp to 5
	}

	for _, tt := range tests {
		result := Border.Size(tt.size)
		if result != tt.expected {
			t.Errorf("Border.Size(%d) = %v, want %v", tt.size, result, tt.expected)
		}
	}
}

func TestRadius(t *testing.T) {
	// Test regular radius
	result := Radius(3)
	expected := "var(--radius-3)"
	if result != expected {
		t.Errorf("Radius(3) = %v, want %v", result, expected)
	}

	// Test round radius
	result = RadiusRound()
	expected = "var(--radius-round)"
	if result != expected {
		t.Errorf("RadiusRound() = %v, want %v", result, expected)
	}

	// Test blob radius
	result = RadiusBlob(2)
	expected = "var(--radius-blob-2)"
	if result != expected {
		t.Errorf("RadiusBlob(2) = %v, want %v", result, expected)
	}
}

func TestEasing(t *testing.T) {
	tests := []struct {
		name     string
		fn       func() string
		expected string
	}{
		{"Default", func() string { return Ease.Default(3) }, "var(--ease-3)"},
		{"In", func() string { return Ease.In(2) }, "var(--ease-in-2)"},
		{"Out", func() string { return Ease.Out(4) }, "var(--ease-out-4)"},
		{"Elastic", func() string { return Ease.Elastic(1) }, "var(--ease-elastic-1)"},
		{"CircIn", Ease.CircIn, "var(--ease-circ-in)"},
		{"CubicInOut", Ease.CubicInOut, "var(--ease-cubic-in-out)"},
	}

	for _, tt := range tests {
		result := tt.fn()
		if result != tt.expected {
			t.Errorf("%s = %v, want %v", tt.name, result, tt.expected)
		}
	}
}

func TestThemeColors(t *testing.T) {
	tests := []struct {
		name     string
		fn       func() string
		expected string
	}{
		{"Background", Color.Background, "var(--background)"},
		{"Surface", Color.Surface, "var(--surface)"},
		{"Text", Color.Text, "var(--text)"},
		{"Primary", Color.Primary, "var(--primary)"},
		{"Border", Color.Border, "var(--border)"},
	}

	for _, tt := range tests {
		result := tt.fn()
		if result != tt.expected {
			t.Errorf("Color.%s() = %v, want %v", tt.name, result, tt.expected)
		}
	}
}