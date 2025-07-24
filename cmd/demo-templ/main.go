package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/riclib/open-props-css/uicss"
)

func main() {
	// Serve the CSS
	http.Handle("/static/dashboard.css", uicss.CSSHandler())

	// Routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		HomePage().Render(r.Context(), w)
	})

	http.HandleFunc("/colors", func(w http.ResponseWriter, r *http.Request) {
		ColorsPage().Render(r.Context(), w)
	})

	http.HandleFunc("/spacing", func(w http.ResponseWriter, r *http.Request) {
		SpacingPage().Render(r.Context(), w)
	})

	http.HandleFunc("/animations", func(w http.ResponseWriter, r *http.Request) {
		AnimationsPage().Render(r.Context(), w)
	})

	http.HandleFunc("/icons", func(w http.ResponseWriter, r *http.Request) {
		IconsPage().Render(r.Context(), w)
	})

	fmt.Println("Op Package Demo running at http://localhost:8081")
	fmt.Println("- Components: http://localhost:8081/")
	fmt.Println("- Colors: http://localhost:8081/colors")
	fmt.Println("- Spacing: http://localhost:8081/spacing")
	fmt.Println("- Animations: http://localhost:8081/animations")
	fmt.Println("- Icons: http://localhost:8081/icons")

	log.Fatal(http.ListenAndServe(":8081", nil))
}