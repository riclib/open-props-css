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
// Size scale 1-15 (plus 00, 000)
op.Size(4)         // "var(--size-4)"
op.Size(-1)        // "var(--size-00)"
op.Size(-2)        // "var(--size-000)"

// Pixel-based sizes
op.SizePx(10)      // "var(--size-px-10)"

// Relative sizes
op.SizeRelative(5) // "var(--size-relative-5)"

// Content sizes
op.SizeContent(2)  // "var(--size-content-2)"
op.SizeHeader(3)   // "var(--size-header-3)"

// Fluid sizes
op.SizeFluid(4)    // "var(--size-fluid-4)"
```

### Typography
```go
// Font sizes 0-8
op.Font.Size(3)           // "var(--font-size-3)"
op.Font.SizeFluid(2)      // "var(--font-size-fluid-2)"

// Font weights 1-9
op.Font.Weight(6)         // "var(--font-weight-6)"

// Line heights 0-5
op.Font.LineHeight(3)     // "var(--font-lineheight-3)"

// Letter spacing 0-7
op.Font.LetterSpacing(2)  // "var(--font-letterspacing-2)"

// Font families
op.Font.Sans()            // "var(--font-sans)"
op.Font.Serif()           // "var(--font-serif)"
op.Font.Mono()            // "var(--font-mono)"
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
op.ShadowColor(4)  // "var(--shadow-color-4)"
op.ShadowStrength(50) // "var(--shadow-strength-50)"
```

### Gradients
```go
// Gradient scale 1-30
op.Gradient(7)     // "var(--gradient-7)"
```

### Animations
```go
// Animation presets
op.Animation.FadeIn()        // "var(--animation-fade-in)"
op.Animation.FadeOut()       // "var(--animation-fade-out)"
op.Animation.ScaleUp()       // "var(--animation-scale-up)"
op.Animation.ScaleDown()     // "var(--animation-scale-down)"
op.Animation.SlideInUp()     // "var(--animation-slide-in-up)"
op.Animation.SlideInDown()   // "var(--animation-slide-in-down)"
op.Animation.SlideInLeft()   // "var(--animation-slide-in-left)"
op.Animation.SlideInRight()  // "var(--animation-slide-in-right)"
op.Animation.Bounce()        // "var(--animation-bounce)"
op.Animation.Pulse()         // "var(--animation-pulse)"
op.Animation.Spin()          // "var(--animation-spin)"
// ... and many more
```

### Easing Functions
```go
// Basic easing 1-5
op.Ease.Default(3)     // "var(--ease-3)"
op.Ease.In(2)          // "var(--ease-in-2)"
op.Ease.Out(4)         // "var(--ease-out-4)"
op.Ease.InOut(3)       // "var(--ease-in-out-3)"

// Elastic easing
op.Ease.Elastic(2)     // "var(--ease-elastic-2)"
op.Ease.Squish(3)      // "var(--ease-squish-3)"
op.Ease.Spring(4)      // "var(--ease-spring-4)"

// Named easing
op.Ease.CircIn()       // "var(--ease-circ-in)"
op.Ease.CubicInOut()   // "var(--ease-cubic-in-out)"
op.Ease.ExpoOut()      // "var(--ease-expo-out)"
```

### Style Builder
The style builder provides a fluent interface for creating inline styles:
```go
style := op.NewStyle().
    Background(op.Gradient(15)).
    Color(op.Color.Blue(9)).
    Padding(op.Size(4)).
    PaddingTop(op.Size(6)).
    Margin(op.Size(2)).
    BorderRadius(op.Radius(3)).
    BoxShadow(op.Shadow(4)).
    FontSize(op.Font.Size(3)).
    FontWeight(op.Font.Weight(6)).
    Animation(op.Animation.FadeIn()).
    TransitionProperty("all").
    TransitionDuration("200ms").
    TransitionTimingFunction(op.Ease.Out(3)).
    String()

// Result: "background: var(--gradient-15); color: var(--blue-9); ..."
```

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