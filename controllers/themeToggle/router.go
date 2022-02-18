package themeToggleController

import "github.com/HoloPirates/mogupantsu/controllers/router"

func init() {
	router.Get().Any("/dark", toggleThemeHandler)
	router.Get().Any("/dark/*redirect", toggleThemeHandler)
}
