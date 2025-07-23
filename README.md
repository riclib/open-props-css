# Open Props CSS Framework for Go

A minimal, classless CSS framework for Go web applications, built on [Open Props](https://open-props.style/) design tokens. Create beautiful dashboard interfaces with semantic HTML—no classes required.

## Features

- 🎨 **Dual theme support**: Automatic light/dark theme switching based on system preferences
- 📦 **Zero JavaScript**: Pure CSS solution, no runtime dependencies
- 🏗️ **Classless by default**: Semantic HTML looks great out of the box
- 📱 **Responsive**: Mobile-first design that works on all devices
- 🚀 **Go-native**: Embedded CSS files with simple HTTP handlers
- 🎯 **Dashboard-ready**: Pre-styled navigation and layout components

## Installation

```bash
go get github.com/riclib/open-props-css
```

## Quick Start

```go
package main

import (
    "net/http"
    "github.com/riclib/open-props-css/uicss"
)

func main() {
    // Serve the CSS file
    http.Handle("/static/dashboard.css", uicss.CSSHandler())
    
    // Serve the interactive stylebook
    http.Handle("/stylebook", uicss.StylebookHandler())
    
    // Your application routes...
    http.HandleFunc("/", dashboardHandler)
    
    http.ListenAndServe(":8080", nil)
}
```

In your HTML:

```html
<!DOCTYPE html>
<html>
<head>
    <link rel="stylesheet" href="/static/dashboard.css">
</head>
<body>
    <nav>
        <h1>My App</h1>
        <ul>
            <li><a href="#" class="active">Dashboard</a></li>
            <li><a href="#">Settings</a></li>
        </ul>
    </nav>
    <main>
        <header>
            <h2>Overview</h2>
            <button class="primary">New Item</button>
        </header>
        
        <table>
            <thead>
                <tr>
                    <th>Name</th>
                    <th>Status</th>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <td>Project Alpha</td>
                    <td><span class="badge success">Active</span></td>
                </tr>
            </tbody>
        </table>
    </main>
</body>
</html>
```

## API

### `uicss.CSSHandler() http.Handler`
Returns an HTTP handler that serves the compiled CSS file with appropriate caching headers.

### `uicss.StylebookHandler() http.Handler`
Returns an HTTP handler that serves an interactive stylebook showcasing all available components and styles.

### `uicss.CSS() string`
Returns the raw CSS as a string for inline embedding or custom processing.

## Theme Support

The framework automatically respects the user's system theme preference. You can also manually control the theme:

```html
<!-- Auto theme (default) -->
<html>

<!-- Force light theme -->
<html class="light">

<!-- Force dark theme -->
<html class="dark">
```


## Semantic HTML Structure

The framework is designed to work with semantic HTML:

- `<nav>` - Automatically styled as a sidebar
- `<main>` - Content area with proper spacing
- `<header>` - Page headers with flex layout
- `<table>` - Styled data tables with hover states
- `<button>` - Styled buttons with hover/active states
- `<form>` elements - Inputs, selects, and textareas

## Optional CSS Classes

While the framework is classless by default, some optional classes are available:

### Buttons
- `.primary` - Primary action button
- `.secondary` - Secondary action button
- `.small`, `.large` - Size modifiers

### Components
- `.badge` - Inline status indicators
- `.card` - Content containers
- `.alert` - Notification messages
- `.active` - Active navigation state

### Utilities
- `.text-muted` - Muted text color
- `.text-center`, `.text-right` - Text alignment
- `.flex`, `.grid` - Layout helpers
- `.mt-1` to `.mt-4` - Margin top spacing
- `.gap-1` to `.gap-4` - Flex/grid gap spacing

## Development

To modify the CSS:

1. Clone the repository
2. Install dependencies: `npm install`
3. Edit CSS files in `src/`
4. Build: `npm run build`
5. Test: `go run cmd/demo/main.go`

## Demo

Run the demo server to see all components:

```bash
go run cmd/demo/main.go
```

Then visit:
- Dashboard demo: http://localhost:8080/
- Interactive stylebook: http://localhost:8080/stylebook

## License

MIT

## Credits

Built on top of [Open Props](https://open-props.style/) by Adam Argyle.