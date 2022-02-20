package activitiesController

import "github.com/HoloPirates/mogupantsu/controllers/router"

func init() {
	router.Get().Any("/activities", ActivityListHandler)
	router.Get().Any("/activities/p/:page", ActivityListHandler)
}
