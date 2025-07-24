# Open Props CSS Framework

This project uses the Open Props CSS Framework for Go, which provides:

## Quick Reference

### CSS Import
```go
import "github.com/riclib/open-props-css/uicss"

// Serve CSS
http.Handle("/static/dashboard.css", uicss.CSSHandler())
```

### HTML Setup
```html
<link rel="stylesheet" href="/static/dashboard.css">
```

### Key Features
- **Automatic dark/light themes** - Respects system preferences
- **Semantic HTML styling** - `<nav>`, `<main>`, `<table>` etc. are pre-styled
- **Dashboard layout** - Sidebar navigation + main content area built-in
- **Open Props variables** - Access to 250+ design tokens

### Common Classes
- `.primary`, `.secondary` - Button variants
- `.badge` - Status indicators (+ `.success`, `.warning`, `.danger`)
- `.card` - Content containers
- `.alert` - Notification messages
- `.grid`, `.flex` - Layout utilities
- `.mt-1` to `.mt-4` - Margin utilities
- `.p-1` to `.p-4` - Padding utilities

### Customization
Use Open Props variables in inline styles:
```go
<div style={fmt.Sprintf("background: %s; padding: %s;", op.Gradient(7), op.Size(4))}>
  <p style={fmt.Sprintf("color: %s; font-size: %s;", op.Color.Indigo(9), op.Font.Size(3))}>Custom styled</p>
</div>
```

Or using the Style builder:
```go
<div style={op.NewStyle().Background(op.Gradient(7)).Padding(op.Size(4)).String()}>
  <p style={op.NewStyle().Color(op.Color.Indigo(9)).FontSize(op.Font.Size(3)).String()}>Custom styled</p>
</div>
```

### Go API (op package)
Type-safe access to Open Props variables:
```go
import "github.com/riclib/open-props-css/op"

// Colors with scale 0-12
style := fmt.Sprintf("color: %s; background: %s;", 
    op.Color.Blue(6), 
    op.Color.Surface())

// All color families available:
// Gray, Stone, Red, Pink, Purple, Violet, Indigo, Blue, Cyan, Teal,
// Green, Lime, Yellow, Orange, Choco, Brown, Sand, Camo, Jungle

// Sizes and spacing
padding := op.Size(4)              // var(--size-4)
margin := op.Size(-1)              // var(--size-00)
fluidSpacing := op.SizeFluid(3)    // var(--size-fluid-3)
contentWidth := op.SizeContent(2)  // var(--size-content-2)

// Typography
fontSize := op.Font.Size(3)        // var(--font-size-3)
fontWeight := op.Font.Weight(6)    // var(--font-weight-6) = 600
lineHeight := op.Font.LineHeight(3) // var(--font-lineheight-3)
fontFamily := op.Font.Sans()       // var(--font-sans)

// Animations & Easing
fadeIn := op.Animation.FadeIn()    // var(--animation-fade-in)
easeOut := op.Ease.Out(3)          // var(--ease-out-3)

// Other utilities
shadow := op.Shadow(3)             // var(--shadow-3)
radius := op.Radius(2)             // var(--radius-2)
gradient := op.Gradient(15)        // var(--gradient-15)

// Style builder with all methods
style := op.NewStyle().
    Background(op.Gradient(7)).
    Color(op.Color.Text()).
    Padding(op.Size(4)).
    PaddingTop(op.Size(6)).
    Margin(op.Size(2)).
    BorderRadius(op.Radius(3)).
    BoxShadow(op.Shadow(2)).
    FontSize(op.Font.Size(3)).
    FontWeight(op.Font.Weight(6)).
    Animation(op.Animation.FadeIn()).
    TransitionProperty("all").
    TransitionDuration("200ms").
    TransitionTimingFunction(op.Ease.Out(3)).
    Custom("gap", op.Size(3)).
    String()
```

For full details, see CSSREF.md