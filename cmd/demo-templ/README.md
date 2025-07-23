# Op Package Demo with Templ

This demo showcases the `op` package (type-safe Open Props API) together with the CSS framework using [templ](https://templ.guide/).

## Running the Demo

1. Install templ CLI:
```bash
go install github.com/a-h/templ/cmd/templ@latest
```

2. Generate the templ files:
```bash
cd cmd/demo-templ
templ generate
```

3. Run the demo:
```bash
go run .
```

4. Open http://localhost:8081 in your browser

## Features Demonstrated

### Dynamic Styling with op Package
- **Buttons** - Size and color variants using `op.NewStyle()`
- **Cards** - Theme-aware styling with `op.Color.Surface()`
- **Spacing** - Visual guide showing `op.Size()` scale
- **Colors** - Interactive color scales using `op.Color.Blue()` etc.
- **Animations** - All Open Props animations via `op.Animation`

### Code Examples

#### Dynamic Button Style
```go
style := op.NewStyle().
    Padding(op.Size(size)).
    BorderRadius(op.Radius(2)).
    Background(op.Color.Primary()).
    Color("white").
    String()
```

#### Theme-Aware Card
```go
style := op.NewStyle().
    Background(op.Color.Surface()).
    Padding(op.Size(5)).
    BorderRadius(op.Radius(3)).
    BoxShadow(op.Shadow(3)).
    String()
```

## Why This Demo is Separate

This demo has its own `go.mod` file to ensure that `templ` doesn't become a dependency for users of the main `open-props-css` packages. Users can import and use `uicss` and `op` packages without pulling in templ as a dependency.