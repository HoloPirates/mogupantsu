package torrentController

import (
	"github.com/HoloPirates/mogupantsu/controllers/middlewares"
	"github.com/HoloPirates/mogupantsu/controllers/router"
)

func init() {
	router.Get().Any("/download/:hash", DownloadTorrent)
	router.Get().Any("/stats/:id", GetStatsHandler)

	torrentRoutes := router.Get().Group("/torrent", middlewares.LoggedInMiddleware())
	{
		torrentRoutes.GET("/", TorrentEditUserPanel)
		torrentRoutes.POST("/", TorrentPostEditUserPanel)
		torrentRoutes.GET("/tag", ViewFormTag)
		torrentRoutes.POST("/tag", ViewFormTag)
		torrentRoutes.GET("/tag/add", AddTag)
		torrentRoutes.GET("/tag/remove", DeleteTag)
		torrentRoutes.POST("/delete", TorrentDeleteUserPanel)
	}
	torrentViewRoutes := router.Get().Group("/view")
	{
		torrentViewRoutes.GET("/:id", ViewHandler)
		torrentViewRoutes.HEAD("/:id", ViewHeadHandler)
		torrentViewRoutes.POST("/:id", PostCommentHandler)
	}
}
