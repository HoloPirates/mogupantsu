package oauth

import (
	"github.com/HoloPirates/mogupantsu/utils/log"
	"github.com/HoloPirates/mogupantsu/utils/oauth2"
	"github.com/gin-gonic/gin"
	"github.com/ory/fosite"
)

func introspectionEndpoint(c *gin.Context) {
	ctx := fosite.NewContext()
	mySessionData := oauth2.NewSession("", "")
	ir, err := oauth2.Oauth2.NewIntrospectionRequest(ctx, c.Request, mySessionData)
	if err != nil {
		log.Errorf("Error occurred in NewAuthorizeRequest: %s\n", err)
		oauth2.Oauth2.WriteIntrospectionError(c.Writer, err)
		return
	}

	oauth2.Oauth2.WriteIntrospectionResponse(c.Writer, ir)
}
