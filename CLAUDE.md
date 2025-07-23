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
```html
<div style="background: var(--gradient-7); padding: var(--size-4);">
  <p style="color: var(--indigo-9); font-size: var(--font-size-3);">Custom styled</p>
</div>
```

For full details, see CSSREF.md