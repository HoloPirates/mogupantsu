package settingsController

import "github.com/HoloPirates/mogupantsu/controllers/router"

func init() {
	router.Get().GET("/settings", SeePublicSettingsHandler)
	router.Get().POST("/settings", ChangePublicSettingsHandler)
}
