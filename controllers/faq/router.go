package faqController

import "github.com/HoloPirates/mogupantsu/controllers/router"

func init() {
	router.Get().Any("/faq", FaqHandler)
}
