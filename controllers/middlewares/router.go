package middlewares

import "github.com/HoloPirates/mogupantsu/controllers/router"

func init() {
	router.Get().Use(CSP(), ErrorMiddleware())
}
