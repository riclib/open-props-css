# Open Props CSS Framework - Complete Reference

## Table of Contents
1. [Installation & Setup](#installation--setup)
2. [Theme System](#theme-system)
3. [Layout Components](#layout-components)
4. [Typography](#typography)
5. [Components](#components)
6. [Utility Classes](#utility-classes)
7. [Open Props Variables](#open-props-variables)
8. [Customization Examples](#customization-examples)
9. [Go API (op package)](#go-api-op-package)

## Installation & Setup

### Go Import
```go
import "github.com/riclib/open-props-css/uicss"

// Available functions
uicss.CSSHandler()      // Returns http.Handler for CSS
uicss.StylebookHandler() // Returns http.Handler for stylebook
uicss.CSS()             // Returns raw CSS string
```

### HTML Structure
```html
<!DOCTYPE html>
<html> <!-- Add class="light" or class="dark" to force theme -->
<head>
    <link rel="stylesheet" href="/static/dashboard.css">
</head>
<body>
    <nav>
        <!-- Automatically styled as sidebar -->
    </nav>
    <main>
        <!-- Main content area -->
    </main>
</body>
</html>
```

## Theme System

### Automatic Theme Switching
The framework automatically switches between light and dark themes based on system preferences using `@media (prefers-color-scheme)`.

### Manual Theme Control
```html
<html class="light">  <!-- Force light theme -->
<html class="dark">   <!-- Force dark theme -->
```

### Theme Variables
```css
--background: light-dark(#ffffff, #0a0a0b);
--surface: light-dark(#f8f9fa, #0f0f10);
--surface-alt: light-dark(#e9ecef, #1a1a1b);
--text: light-dark(#212529, #ffffff);
--text-muted: light-dark(#6c757d, #a0a0a0);
--border: light-dark(#dee2e6, #1a1a1b);
--primary: light-dark(#0066cc, #3b82f6);
--primary-hover: light-dark(#0052a3, #2563eb);
```

## Layout Components

### Dashboard Layout (Automatic)
```html
<body>
    <nav>
        <h1>App Name</h1>
        <ul>
            <li><a href="#" class="active">Dashboard</a></li>
            <li><a href="#">Settings</a></li>
        </ul>
    </nav>
    <main>
        <header>
            <h2>Page Title</h2>
            <button class="primary">Action</button>
        </header>
        <!-- Content -->
    </main>
</body>
```

### Responsive Behavior
- Desktop: Sidebar (250px) + Main content
- Mobile: Stacked vertically

## Typography

### Headings (Semantic)
```html
<h1>Heading 1</h1>  <!-- font-size: var(--font-size-7) -->
<h2>Heading 2</h2>  <!-- font-size: var(--font-size-6) -->
<h3>Heading 3</h3>  <!-- font-size: var(--font-size-5) -->
<h4>Heading 4</h4>  <!-- font-size: var(--font-size-4) -->
<h5>Heading 5</h5>  <!-- font-size: var(--font-size-3) -->
<h6>Heading 6</h6>  <!-- font-size: var(--font-size-2) -->
```

### Text Utilities
```html
<p class="text-muted">Muted text</p>
<p class="text-small">Small text</p>
<p class="text-large">Large text</p>
<p class="text-center">Centered</p>
<p class="text-right">Right aligned</p>
<p class="text-primary">Primary color</p>
<p class="font-bold">Bold weight</p>
<p class="font-medium">Medium weight</p>
<p class="font-normal">Normal weight</p>
```

## Components

### Buttons
```html
<!-- Variants -->
<button>Default Button</button>
<button class="primary">Primary Button</button>
<button class="secondary">Secondary Button</button>
<button disabled>Disabled Button</button>

<!-- Sizes -->
<button class="small">Small</button>
<button class="large">Large</button>
```

### Forms
```html
<!-- Form group -->
<div class="form-group">
    <label for="input">Label</label>
    <input type="text" id="input" placeholder="Enter text">
</div>

<!-- Input types (all styled) -->
<input type="text">
<input type="email">
<input type="password">
<textarea rows="3"></textarea>
<select>
    <option>Option 1</option>
    <option>Option 2</option>
</select>

<!-- Input group -->
<div class="input-group">
    <input type="text" placeholder="Search...">
    <button class="primary">Search</button>
</div>
```

### Cards
```html
<div class="card">
    <h3>Card Title</h3>
    <p>Card content goes here.</p>
    <button class="primary">Action</button>
</div>
```

### Alerts
```html
<div class="alert info">Info message</div>
<div class="alert success">Success message</div>
<div class="alert warning">Warning message</div>
<div class="alert danger">Error message</div>
```

### Badges
```html
<span class="badge">Default</span>
<span class="badge primary">Primary</span>
<span class="badge success">Success</span>
<span class="badge warning">Warning</span>
<span class="badge danger">Danger</span>
```

### Tables
```html
<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Status</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>John Doe</td>
            <td><span class="badge success">Active</span></td>
        </tr>
    </tbody>
</table>
```

### Loading Spinner
```html
<div class="spinner"></div>
```

## Utility Classes

### Spacing
```css
/* Margin */
.m-0 to .m-4    /* All sides */
.mt-0 to .mt-4  /* Top */
.mb-0 to .mb-4  /* Bottom */
.m-auto         /* Auto margins */

/* Padding */
.p-0 to .p-4    /* All sides */
.px-1 to .px-4  /* Horizontal */
.py-1 to .py-4  /* Vertical */
```

### Display
```css
.hidden         /* display: none */
.block          /* display: block */
.inline         /* display: inline */
.inline-block   /* display: inline-block */
.flex           /* display: flex */
.inline-flex    /* display: inline-flex */
.grid           /* display: grid */
```

### Flexbox
```css
.flex-row       /* flex-direction: row */
.flex-col       /* flex-direction: column */
.flex-wrap      /* flex-wrap: wrap */
.items-start    /* align-items: flex-start */
.items-center   /* align-items: center */
.items-end      /* align-items: flex-end */
.justify-start  /* justify-content: flex-start */
.justify-center /* justify-content: center */
.justify-end    /* justify-content: flex-end */
.justify-between /* justify-content: space-between */
.gap-1 to .gap-4 /* gap spacing */
```

### Grid
```css
.grid           /* display: grid */
.grid-cols-2    /* 2 columns */
.grid-cols-3    /* 3 columns */
.grid-cols-4    /* 4 columns */
/* Responsive: stacks on mobile */
```

### Border & Shadow
```css
.border         /* 1px solid border */
.border-0       /* Remove border */
.border-t/b/l/r /* Directional borders */
.rounded        /* border-radius: var(--radius-2) */
.rounded-full   /* Fully rounded */
.shadow         /* Box shadow */
.shadow-lg      /* Large shadow */
```

### Width & Overflow
```css
.w-full         /* width: 100% */
.w-auto         /* width: auto */
.overflow-hidden
.overflow-auto
.overflow-x-auto
.overflow-y-auto
```

### Other
```css
.cursor-pointer
.cursor-not-allowed
.select-none
.select-text
.sr-only        /* Screen reader only */
```

## Open Props Variables

### Color Scales (0-12 for each)
```css
--gray-[0-12]    /* Neutral grays */
--stone-[0-12]   /* Warm grays */
--red-[0-12]
--pink-[0-12]
--purple-[0-12]
--violet-[0-12]
--indigo-[0-12]
--blue-[0-12]
--cyan-[0-12]
--teal-[0-12]
--green-[0-12]
--lime-[0-12]
--yellow-[0-12]
--orange-[0-12]
--brown-[0-12]
```

### Spacing
```css
--size-1   /* 0.25rem */
--size-2   /* 0.5rem */
--size-3   /* 1rem */
--size-4   /* 1.25rem */
--size-5   /* 1.5rem */
--size-6   /* 1.75rem */
--size-7   /* 2rem */
--size-8   /* 3rem */
--size-9   /* 4rem */
--size-10  /* 5rem */
/* ... up to --size-15 */
```

### Typography
```css
/* Font sizes */
--font-size-0  /* 0.75rem */
--font-size-1  /* 1rem */
--font-size-2  /* 1.1rem */
--font-size-3  /* 1.25rem */
--font-size-4  /* 1.5rem */
--font-size-5  /* 2rem */
--font-size-6  /* 2.5rem */
--font-size-7  /* 3rem */
--font-size-8  /* 3.5rem */

/* Font weights */
--font-weight-1 to --font-weight-9 /* 100 to 900 */

/* Line heights */
--font-lineheight-0 to --font-lineheight-5
```

### Radius
```css
--radius-1     /* 2px */
--radius-2     /* 5px */
--radius-3     /* 1rem */
--radius-4     /* 2rem */
--radius-5     /* 4rem */
--radius-6     /* 8rem */
--radius-round /* 1e5px (fully round) */
```

### Shadows
```css
--shadow-1 to --shadow-6
--inner-shadow-0 to --inner-shadow-4
```

### Gradients
```css
--gradient-1 to --gradient-30
/* Examples: */
--gradient-7   /* Blue gradient */
--gradient-15  /* Soft pastel */
--gradient-22  /* Cool blue */
```

### Animations
```css
--animation-fade-in
--animation-fade-out
--animation-scale-up
--animation-scale-down
--animation-slide-in-[up/down/left/right]
--animation-slide-out-[up/down/left/right]
--animation-shake-[x/y/z]
--animation-spin
--animation-ping
--animation-blink
--animation-float
--animation-bounce
--animation-pulse
```

### Easing Functions
```css
--ease-1 to --ease-5
--ease-in-1 to --ease-in-5
--ease-out-1 to --ease-out-5
--ease-in-out-1 to --ease-in-out-5
```

## Customization Examples

### Custom Colored Elements
```html
<!-- Purple button -->
<button style="background-color: var(--purple-6); border-color: var(--purple-7);">
    Purple Button
</button>

<!-- Gradient background -->
<div style="background: var(--gradient-15); padding: var(--size-5);">
    Gradient content
</div>

<!-- Custom card -->
<div class="card" style="border: 2px solid var(--indigo-3); 
                         box-shadow: var(--shadow-4);">
    Custom styled card
</div>
```

### Custom Spacing & Typography
```html
<div style="padding: var(--size-7); margin-block: var(--size-5);">
    <h2 style="font-size: var(--font-size-6); 
               font-weight: var(--font-weight-7);
               line-height: var(--font-lineheight-1);">
        Custom heading
    </h2>
    <p style="color: var(--gray-6); font-size: var(--font-size-0);">
        Small muted text
    </p>
</div>
```

### Animations
```html
<button style="animation: var(--animation-pulse);">
    Pulsing button
</button>

<div style="animation: var(--animation-fade-in);">
    Fade in content
</div>
```

### Complex Compositions
```html
<div style="background: var(--gradient-22);
            padding: var(--size-6);
            border-radius: var(--radius-3);
            box-shadow: var(--shadow-3);
            animation: var(--animation-fade-in);">
    <h3 style="color: white; margin-bottom: var(--size-3);">
        Gradient Card
    </h3>
    <p style="color: rgba(255,255,255,0.9);">
        Fully customized with Open Props
    </p>
</div>
```

## Go API (op package)

The `op` package provides type-safe Go access to all Open Props CSS variables. This allows you to generate dynamic styles in your Go code while maintaining consistency with the design system.

### Installation
```go
import "github.com/riclib/open-props-css/op"
```

### Colors

#### Color Scales
All Open Props colors are available with scale values 0-12:
```go
// Basic usage
op.Color.Gray(6)    // "var(--gray-6)"
op.Color.Blue(9)    // "var(--blue-9)"
op.Color.Red(5)     // "var(--red-5)"

// Available colors:
// Gray, Stone, Red, Pink, Purple, Violet, Indigo, Blue, Cyan, Teal,
// Green, Lime, Yellow, Orange, Choco, Brown, Sand, Camo, Jungle

// Theme colors
op.Color.Background()    // "var(--background)"
op.Color.Surface()       // "var(--surface)"
op.Color.SurfaceAlt()    // "var(--surface-alt)"
op.Color.Text()          // "var(--text)"
op.Color.TextMuted()     // "var(--text-muted)"
op.Color.Border()        // "var(--border)"
op.Color.Primary()       // "var(--primary)"
op.Color.PrimaryHover()  // "var(--primary-hover)"
```

### Sizes and Spacing
```go
// Size scale -2 to 15
op.Size(4)         // "var(--size-4)"
op.Size(-1)        // "var(--size-00)"
op.Size(-2)        // "var(--size-000)"
op.Size(15)        // "var(--size-15)"

// Pixel-based sizes -2 to 15
op.SizePx(10)      // "var(--size-px-10)"
op.SizePx(-1)      // "var(--size-px-00)"

// Relative sizes 0-15
op.SizeRelative(5) // "var(--size-relative-5)"

// Content sizes 1-3
op.SizeContent(1)  // "var(--size-content-1)" = 20ch
op.SizeContent(2)  // "var(--size-content-2)" = 45ch
op.SizeContent(3)  // "var(--size-content-3)" = 60ch

// Header sizes 1-3
op.SizeHeader(1)   // "var(--size-header-1)" = 20ch
op.SizeHeader(2)   // "var(--size-header-2)" = 25ch
op.SizeHeader(3)   // "var(--size-header-3)" = 35ch

// Fluid sizes 0-10
op.SizeFluid(0)    // "var(--size-fluid-0)"
op.SizeFluid(4)    // "var(--size-fluid-4)"
op.SizeFluid(10)   // "var(--size-fluid-10)"

// Layer/z-index values 1-5
op.Layer(1)        // "var(--layer-1)"
op.Layer(5)        // "var(--layer-5)"
op.LayerImportant() // "var(--layer-important)"

// Aspect ratios
op.Ratio("square")     // "var(--ratio-square)" = 1
op.Ratio("landscape")  // "var(--ratio-landscape)" = 4/3
op.Ratio("portrait")   // "var(--ratio-portrait)" = 3/4
op.Ratio("widescreen") // "var(--ratio-widescreen)" = 16/9
op.Ratio("ultrawide")  // "var(--ratio-ultrawide)" = 18/5
op.Ratio("golden")     // "var(--ratio-golden)" = 1.618
```

### Typography
```go
// Font sizes 0-8
op.Font.Size(3)           // "var(--font-size-3)"
op.Font.SizeFluid(2)      // "var(--font-size-fluid-2)"

// Font weights 1-9 (100-900)
op.Font.Weight(1)         // "var(--font-weight-1)" = 100
op.Font.Weight(4)         // "var(--font-weight-4)" = 400 (normal)
op.Font.Weight(6)         // "var(--font-weight-6)" = 600 (semibold)
op.Font.Weight(7)         // "var(--font-weight-7)" = 700 (bold)
op.Font.Weight(9)         // "var(--font-weight-9)" = 900 (black)

// Line heights 0-5
op.Font.LineHeight(0)     // "var(--font-lineheight-00)"
op.Font.LineHeight(3)     // "var(--font-lineheight-3)"

// Letter spacing 0-7
op.Font.LetterSpacing(2)  // "var(--font-letterspacing-2)"

// Generic font families
op.Font.Sans()            // "var(--font-sans)"
op.Font.Serif()           // "var(--font-serif)"
op.Font.Mono()            // "var(--font-mono)"

// Specific font family stacks
op.Font.SystemUI()        // "var(--font-system-ui)"
op.Font.Transitional()    // "var(--font-transitional)"
op.Font.OldStyle()        // "var(--font-old-style)"
op.Font.Humanist()        // "var(--font-humanist)"
op.Font.GeometricHumanist() // "var(--font-geometric-humanist)"
op.Font.ClassicalHumanist() // "var(--font-classical-humanist)"
op.Font.NeoGrotesque()    // "var(--font-neo-grotesque)"
op.Font.MonospaceSlab()   // "var(--font-monospace-slab-serif)"
op.Font.MonospaceCode()   // "var(--font-monospace-code)"
op.Font.Industrial()      // "var(--font-industrial)"
op.Font.RoundedSans()     // "var(--font-rounded-sans)"
op.Font.SlabSerif()       // "var(--font-slab-serif)"
op.Font.Antique()         // "var(--font-antique)"
op.Font.Didone()          // "var(--font-didone)"
op.Font.Handwritten()     // "var(--font-handwritten)"
```

### Borders and Radius
```go
// Border sizes 1-5
op.Border.Size(2)         // "var(--border-size-2)"

// Border radius 1-6
op.Radius(3)              // "var(--radius-3)"
op.RadiusRound()          // "var(--radius-round)"
op.RadiusBlob(2)          // "var(--radius-blob-2)"
op.RadiusConditional(4)   // "var(--radius-conditional-4)"
```

### Shadows
```go
// Shadow scale 1-6
op.Shadow(3)       // "var(--shadow-3)"

// Inner shadows 0-4
op.InnerShadow(2)  // "var(--inner-shadow-2)"
```

### Gradients
```go
// Gradient scale 1-30
op.Gradient(7)     // "var(--gradient-7)"
```

### Animations
```go
// Fade animations
op.Animation.FadeIn()        // "var(--animation-fade-in)"
op.Animation.FadeInBloom()   // "var(--animation-fade-in-bloom)"
op.Animation.FadeOut()       // "var(--animation-fade-out)"
op.Animation.FadeOutBloom()  // "var(--animation-fade-out-bloom)"

// Scale animations
op.Animation.ScaleUp()       // "var(--animation-scale-up)"
op.Animation.ScaleDown()     // "var(--animation-scale-down)"

// Slide animations (in and out for all directions)
op.Animation.SlideInUp()     // "var(--animation-slide-in-up)"
op.Animation.SlideInDown()   // "var(--animation-slide-in-down)"
op.Animation.SlideInLeft()   // "var(--animation-slide-in-left)"
op.Animation.SlideInRight()  // "var(--animation-slide-in-right)"
op.Animation.SlideOutUp()    // "var(--animation-slide-out-up)"
op.Animation.SlideOutDown()  // "var(--animation-slide-out-down)"
op.Animation.SlideOutLeft()  // "var(--animation-slide-out-left)"
op.Animation.SlideOutRight() // "var(--animation-slide-out-right)"

// Shake animations
op.Animation.ShakeX()        // "var(--animation-shake-x)"
op.Animation.ShakeY()        // "var(--animation-shake-y)"
op.Animation.ShakeZ()        // "var(--animation-shake-z)"

// Other animations
op.Animation.Spin()          // "var(--animation-spin)"
op.Animation.Ping()          // "var(--animation-ping)"
op.Animation.Blink()         // "var(--animation-blink)"
op.Animation.Float()         // "var(--animation-float)"
op.Animation.Bounce()        // "var(--animation-bounce)"
op.Animation.Pulse()         // "var(--animation-pulse)"
```

### Easing Functions
```go
// Basic easing 1-5
op.Ease.Default(3)        // "var(--ease-3)"
op.Ease.In(2)             // "var(--ease-in-2)"
op.Ease.Out(4)            // "var(--ease-out-4)"
op.Ease.InOut(3)          // "var(--ease-in-out-3)"

// Elastic easing 1-5
op.Ease.Elastic(2)        // "var(--ease-elastic-2)"
op.Ease.ElasticIn(1)      // "var(--ease-elastic-in-1)"
op.Ease.ElasticOut(3)     // "var(--ease-elastic-out-3)"
op.Ease.ElasticInOut(4)   // "var(--ease-elastic-in-out-4)"

// Special easing 1-5
op.Ease.Squish(3)         // "var(--ease-squish-3)"
op.Ease.Spring(4)         // "var(--ease-spring-4)"
op.Ease.Bounce(2)         // "var(--ease-bounce-2)"
op.Ease.Step(5)           // "var(--ease-step-5)"

// Named easing functions
op.Ease.CircIn()          // "var(--ease-circ-in)"
op.Ease.CircOut()         // "var(--ease-circ-out)"
op.Ease.CircInOut()       // "var(--ease-circ-in-out)"
op.Ease.CubicIn()         // "var(--ease-cubic-in)"
op.Ease.CubicOut()        // "var(--ease-cubic-out)"
op.Ease.CubicInOut()      // "var(--ease-cubic-in-out)"
op.Ease.ExpoIn()          // "var(--ease-expo-in)"
op.Ease.ExpoOut()         // "var(--ease-expo-out)"
op.Ease.ExpoInOut()       // "var(--ease-expo-in-out)"
op.Ease.QuadIn()          // "var(--ease-quad-in)"
op.Ease.QuadOut()         // "var(--ease-quad-out)"
op.Ease.QuadInOut()       // "var(--ease-quad-in-out)"
op.Ease.QuartIn()         // "var(--ease-quart-in)"
op.Ease.QuartOut()        // "var(--ease-quart-out)"
op.Ease.QuartInOut()      // "var(--ease-quart-in-out)"
op.Ease.QuintIn()         // "var(--ease-quint-in)"
op.Ease.QuintOut()        // "var(--ease-quint-out)"
op.Ease.QuintInOut()      // "var(--ease-quint-in-out)"
op.Ease.SineIn()          // "var(--ease-sine-in)"
op.Ease.SineOut()         // "var(--ease-sine-out)"
op.Ease.SineInOut()       // "var(--ease-sine-in-out)"
```

### Style Builder
The style builder provides a fluent interface for creating inline styles:
```go
style := op.NewStyle().
    // Colors and backgrounds
    Background(op.Gradient(15)).
    Color(op.Color.Blue(9)).
    
    // Spacing - general and specific
    Padding(op.Size(4)).
    PaddingTop(op.Size(6)).
    PaddingBottom(op.Size(4)).
    PaddingLeft(op.Size(3)).
    PaddingRight(op.Size(3)).
    Margin(op.Size(2)).
    MarginTop(op.Size(4)).
    MarginBottom(op.Size(4)).
    MarginLeft("auto").
    MarginRight("auto").
    
    // Visual effects
    BorderRadius(op.Radius(3)).
    BoxShadow(op.Shadow(4)).
    
    // Typography
    FontSize(op.Font.Size(3)).
    FontWeight(op.Font.Weight(6)).
    
    // Animations and transitions
    Animation(op.Animation.FadeIn()).
    Transition("all 200ms ease").
    TransitionProperty("transform").
    TransitionDuration("300ms").
    TransitionTimingFunction(op.Ease.Out(3)).
    
    // Custom CSS properties
    Custom("gap", op.Size(3)).
    Custom("grid-template-columns", "1fr 2fr").
    
    // Generate the CSS string
    String()

// Result: "background: var(--gradient-15); color: var(--blue-9); padding: var(--size-4); ..."
```

#### Available Style Builder Methods
- `Background(value)` - Set background
- `Color(value)` - Set text color
- `Padding(value)` - Set all padding
- `PaddingTop(value)`, `PaddingBottom(value)`, `PaddingLeft(value)`, `PaddingRight(value)` - Set specific padding
- `Margin(value)` - Set all margins
- `MarginTop(value)`, `MarginBottom(value)`, `MarginLeft(value)`, `MarginRight(value)` - Set specific margins
- `BorderRadius(value)` - Set border radius
- `BoxShadow(value)` - Set box shadow
- `FontSize(value)` - Set font size
- `FontWeight(value)` - Set font weight
- `Animation(value)` - Set animation
- `Transition(value)` - Set transition shorthand
- `TransitionProperty(value)` - Set transition property
- `TransitionDuration(value)` - Set transition duration
- `TransitionTimingFunction(value)` - Set transition timing function
- `Custom(property, value)` - Set any custom CSS property
- `String()` - Generate the final CSS string

### Practical Examples

#### Dynamic Component Styling
```go
func ButtonStyle(variant string, size int) string {
    style := op.NewStyle().
        Padding(op.Size(size)).
        BorderRadius(op.Radius(2)).
        FontWeight(op.Font.Weight(6))
    
    switch variant {
    case "primary":
        style.Background(op.Color.Primary()).
            Color("white").
            Custom("border", "none")
    case "secondary":
        style.Background("transparent").
            Color(op.Color.Text()).
            Custom("border", fmt.Sprintf("1px solid %s", op.Color.Border()))
    }
    
    return style.String()
}
```

#### Theme-Aware Cards
```go
func CardStyle(elevated bool) string {
    style := op.NewStyle().
        Background(op.Color.Surface()).
        Padding(op.Size(5)).
        BorderRadius(op.Radius(3))
    
    if elevated {
        style.BoxShadow(op.Shadow(3))
    } else {
        style.Custom("border", fmt.Sprintf("1px solid %s", op.Color.Border()))
    }
    
    return style.String()
}
```

#### Responsive Spacing
```go
func ResponsiveSection() string {
    return op.NewStyle().
        Padding(op.Size(4)).
        Custom("padding-block", op.SizeFluid(3)).
        Custom("max-width", op.SizeContent(3)).
        Margin("0 auto").
        String()
}
```

### Constants

The op package provides useful constants for common values:

#### Size Constants
```go
// Common sizes
op.SizeXS  = 1   // 0.25rem
op.SizeSM  = 2   // 0.5rem
op.SizeMD  = 3   // 1rem
op.SizeLG  = 4   // 1.25rem
op.SizeXL  = 5   // 1.5rem
op.Size2XL = 6   // 1.75rem
op.Size3XL = 7   // 2rem
op.Size4XL = 8   // 3rem
op.Size5XL = 9   // 4rem
op.Size6XL = 10  // 5rem
```

#### Color Scale Constants
```go
op.ColorScaleLightest = 0
op.ColorScaleLight    = 3
op.ColorScaleMedium   = 6
op.ColorScaleDark     = 9
op.ColorScaleDarkest  = 12
```

#### Font Size Constants
```go
op.FontSizeXS  = 0  // 0.75rem
op.FontSizeSM  = 1  // 1rem
op.FontSizeMD  = 2  // 1.1rem
op.FontSizeLG  = 3  // 1.25rem
op.FontSizeXL  = 4  // 1.5rem
op.FontSize2XL = 5  // 2rem
op.FontSize3XL = 6  // 2.5rem
op.FontSize4XL = 7  // 3rem
op.FontSize5XL = 8  // 3.5rem
```

#### Font Weight Constants
```go
op.FontWeightThin       = 1  // 100
op.FontWeightExtraLight = 2  // 200
op.FontWeightLight      = 3  // 300
op.FontWeightNormal     = 4  // 400
op.FontWeightMedium     = 5  // 500
op.FontWeightSemiBold   = 6  // 600
op.FontWeightBold       = 7  // 700
op.FontWeightExtraBold  = 8  // 800
op.FontWeightBlack      = 9  // 900
```

#### Border Size Constants
```go
op.BorderSizeThin    = 1  // 1px
op.BorderSizeDefault = 2  // 2px
op.BorderSizeMedium  = 3  // 5px
op.BorderSizeThick   = 4  // 10px
op.BorderSizeHeavy   = 5  // 25px
```

#### Radius Constants
```go
op.RadiusXS  = 1  // 2px
op.RadiusSM  = 2  // 5px
op.RadiusMD  = 3  // 1rem
op.RadiusLG  = 4  // 2rem
op.RadiusXL  = 5  // 4rem
op.Radius2XL = 6  // 8rem
```

## Best Practices

1. **Use semantic HTML** - The framework styles semantic elements by default
2. **Prefer built-in classes** - Use utility classes before inline styles
3. **Use Open Props variables** - For consistency across your customizations
4. **Test both themes** - Ensure your customizations work in light and dark modes
5. **Mobile-first** - The framework is responsive by default

## Resources

- [Open Props Documentation](https://open-props.style/)
- [Project Repository](https://github.com/riclib/open-props-css)
- Interactive Stylebook: `/stylebook` route in your app