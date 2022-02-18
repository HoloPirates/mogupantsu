package databasedumpsController

import "github.com/HoloPirates/mogupantsu/controllers/router"

func init() {
	router.Get().Any("/dumps", DatabaseDumpHandler)
}
