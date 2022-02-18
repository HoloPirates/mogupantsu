package controllers

import (
	"net/http"

	_ "github.com/HoloPirates/mogupantsu/controllers/activities"    // activities controller
	_ "github.com/HoloPirates/mogupantsu/controllers/api"           // api controller
	_ "github.com/HoloPirates/mogupantsu/controllers/captcha"       // captcha controller
	_ "github.com/HoloPirates/mogupantsu/controllers/databasedumps" // databasedumps controller
	_ "github.com/HoloPirates/mogupantsu/controllers/faq"           // faq controller
	_ "github.com/HoloPirates/mogupantsu/controllers/feed"          // feed controller
	_ "github.com/HoloPirates/mogupantsu/controllers/middlewares"   // middlewares
	_ "github.com/HoloPirates/mogupantsu/controllers/moderator"     // moderator controller
	_ "github.com/HoloPirates/mogupantsu/controllers/oauth"         // oauth2 controller
	_ "github.com/HoloPirates/mogupantsu/controllers/pprof"         // pprof controller
	_ "github.com/HoloPirates/mogupantsu/controllers/report"        // report controller
	"github.com/HoloPirates/mogupantsu/controllers/router"
	_ "github.com/HoloPirates/mogupantsu/controllers/search"   // search controller
	_ "github.com/HoloPirates/mogupantsu/controllers/settings" // settings controller
	_ "github.com/HoloPirates/mogupantsu/controllers/static"   // static files
	_ "github.com/HoloPirates/mogupantsu/controllers/themeToggle" // themeToggle controller
	_ "github.com/HoloPirates/mogupantsu/controllers/torrent"  // torrent controller
	_ "github.com/HoloPirates/mogupantsu/controllers/upload"   // upload controller
	_ "github.com/HoloPirates/mogupantsu/controllers/user"     // user controller
	"github.com/justinas/nosurf"
)

// CSRFRouter : CSRF protection for Router variable for exporting the route configuration
var CSRFRouter *nosurf.CSRFHandler

func init() {
	CSRFRouter = nosurf.New(router.Get())
	CSRFRouter.ExemptRegexp("/api(?:/.+)*")
	CSRFRouter.ExemptRegexp("/mod(?:/.+)*")
	CSRFRouter.ExemptPath("/upload")
	CSRFRouter.ExemptPath("/user/login")
	CSRFRouter.ExemptPath("/oauth2/token")
	CSRFRouter.SetFailureHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Invalid CSRF tokens", http.StatusBadRequest)
	}))
	CSRFRouter.SetBaseCookie(http.Cookie{
		Path:   "/",
		MaxAge: nosurf.MaxAge,
	})

}
