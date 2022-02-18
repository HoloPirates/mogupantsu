package userController

import (
	"github.com/HoloPirates/mogupantsu/templates"
	"github.com/HoloPirates/mogupantsu/utils/email"
	msg "github.com/HoloPirates/mogupantsu/utils/messages"
	"github.com/gin-gonic/gin"
)

// UserVerifyEmailHandler : Controller when verifying email, needs a token
func UserVerifyEmailHandler(c *gin.Context) {
	token := c.Param("token")
	messages := msg.GetMessages(c)

	_, errEmail := email.EmailVerification(token, c)
	if errEmail != nil {
		messages.ImportFromError("errors", errEmail)
	}
	templates.Static(c, "site/static/verify_success.jet.html")
}
