package captchaController

import (
	"github.com/HoloPirates/mogupantsu/controllers/router"
	"github.com/HoloPirates/mogupantsu/utils/captcha"
)

func init() {
	router.Get().Any("/captcha/*hash", captcha.ServeFiles)
}
