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