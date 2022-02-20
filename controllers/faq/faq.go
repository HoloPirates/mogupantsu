package faqController

import (
	"github.com/HoloPirates/mogupantsu/templates"
	"github.com/gin-gonic/gin"
)

// FaqHandler : Controller for FAQ view page
func FaqHandler(c *gin.Context) {
	templates.Static(c, "site/static/faq.jet.html")
}
