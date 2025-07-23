# Open Props CSS Framework for Go - Detailed Build Plan

## Project Overview
Building a static CSS framework specifically designed for Go web applications, with a focus on creating beautiful dashboard interfaces with minimal effort. The framework combines Open Props design tokens with a shadcn-inspired dark theme aesthetic.

## Visual Design Target
- **Dual theme support**: Both light and dark themes
- **Automatic theme switching**: Respects system preferences
- **Consistent design language**: Same aesthetics in both themes
- **Clean typography**: System fonts with clear hierarchy
- **Subtle borders**: Low contrast dividers in both themes
- **Minimal shadows**: Used sparingly for depth
- **Generous spacing**: Breathing room between elements

## Technical Architecture

### 1. Go Module Structure
```
github.com/riclib/open-props-css/
├── go.mod                  # Go module definition
├── uicss/                  # Main Go package
│   ├── ui.go              # Exports CSSHandler() and StylebookHandler()
│   ├── dashboard.css      # Embedded CSS file (built from src/)
│   └── stylebook.html     # Embedded interactive stylebook
├── cmd/demo/              # Demo server
│   └── main.go           # Example usage
├── src/                   # CSS source files
│   ├── tokens.css        # Open Props + custom variables
│   ├── reset.css         # Minimal modern reset
│   ├── base.css          # Default HTML element styles
│   ├── components.css    # Optional component classes
│   └── utilities.css     # Helper classes
├── package.json          # NPM deps for CSS building only
├── build.js              # Build script
└── README.md
```

### 2. Go Package API
```go
package uicss

// CSSHandler() - Serves the dashboard.css file
// StylebookHandler() - Serves an interactive stylebook
// CSS() - Returns the raw CSS as a string
```

### 3. CSS Philosophy
- **Classless by default**: Plain HTML elements look great
- **Progressive enhancement**: Add classes only when needed
- **Semantic HTML**: `<nav>` becomes sidebar, `<main>` becomes content area
- **90% use case**: Default styles cover most needs

## Implementation Details

### Phase 1: Project Setup
1. Initialize Go module: `go mod init github.com/riclib/open-props-css`
2. Create directory structure
3. Setup NPM for CSS tooling (PostCSS, Open Props)
4. Configure build pipeline

### Phase 2: Design System
1. **Tokens**: Import Open Props, add custom theme variables
2. **Theme Strategy**:
   - CSS custom properties that adapt to theme
   - `@media (prefers-color-scheme: dark)` for automatic switching
   - Optional `.light` and `.dark` classes for manual control
3. **Color System**:
   ```css
   /* Semantic variables that switch automatically */
   --background: light-dark(#ffffff, #0a0a0b);
   --surface: light-dark(#f8f9fa, #0f0f10);
   --surface-alt: light-dark(#e9ecef, #1a1a1b);
   --text: light-dark(#212529, #ffffff);
   --text-muted: light-dark(#6c757d, #a0a0a0);
   --border: light-dark(#dee2e6, #1a1a1b);
   --primary: light-dark(#0066cc, #3b82f6);
   ```
4. **Typography**: System font stack, modular scale
5. **Spacing**: Open Props spacing scale (--size-1 through --size-10)

### Phase 3: Base Styles
1. **Reset**: Modern CSS reset (minimal, not aggressive)
2. **HTML Elements**:
   - `body`: Dark background, light text, system fonts
   - `h1-h6`: Clear hierarchy with appropriate spacing
   - `p`, `ul`, `ol`: Readable line-height and spacing
   - `a`: Subtle blue, hover states
   - `button`: Dark with border, hover/active states
   - `input`, `textarea`, `select`: Dark backgrounds, subtle borders
   - `table`: Alternating row colors, clean borders

### Phase 4: Layout System
1. **Dashboard Layout**:
   ```css
   nav {
     /* Automatic sidebar: 250px wide, full height, dark bg */
   }
   main {
     /* Content area: flex-grow, padding, scrollable */
   }
   body {
     /* Flex container, horizontal layout */
   }
   ```
2. **Responsive**: Stack on mobile, side-by-side on desktop

### Phase 5: Components
Optional classes for enhanced styling:
- `.primary`, `.secondary`: Button variants
- `.small`, `.large`: Size modifiers
- `.badge`: Inline status indicators
- `.card`: Content containers
- `.active`: Navigation state

### Phase 6: Go Integration
1. Build CSS to `uicss/dashboard.css`
2. Create embed directives in Go
3. Implement HTTP handlers
4. Create demo server

### Phase 7: Documentation
1. **Stylebook**: Interactive HTML showing all elements
2. **README**: Installation, usage, examples
3. **Demo**: Full dashboard example

## Usage Example
```go
package main

import (
    "net/http"
    "github.com/riclib/open-props-css/uicss"
)

func main() {
    // Serve the CSS
    http.Handle("/static/dashboard.css", uicss.CSSHandler())
    
    // Serve the stylebook
    http.Handle("/stylebook", uicss.StylebookHandler())
    
    // Your app...
    http.HandleFunc("/", dashboardHandler)
    
    http.ListenAndServe(":8080", nil)
}
```

```html
<!DOCTYPE html>
<html>
<head>
    <link rel="stylesheet" href="/static/dashboard.css">
</head>
<body>
    <!-- Automatically uses system theme preference -->
    <!-- Or add class="dark" or class="light" to force a theme -->
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
            <!-- Beautiful by default in both themes -->
        </table>
    </main>
</body>
</html>
```

## Success Criteria
1. **Zero JavaScript dependencies** for end users
2. **Single CSS file** import
3. **Beautiful defaults** without any classes
4. **Easy dashboard creation** with semantic HTML
5. **Both light and dark themes** with automatic switching
6. **Matches target aesthetic** (modern, clean, professional)
7. **Go-native distribution** via embedded files

## Development Workflow
1. Edit CSS sources in `src/`
2. Run `npm run build` to compile
3. Test with `go run cmd/demo/main.go`
4. Commit all files including built CSS
5. Tag releases for Go module versioning