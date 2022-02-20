package moderatorController

import (
	"github.com/HoloPirates/mogupantsu/models"
	"github.com/HoloPirates/mogupantsu/models/comments"
	"github.com/HoloPirates/mogupantsu/models/reports"
	"github.com/HoloPirates/mogupantsu/models/torrents"
	"github.com/HoloPirates/mogupantsu/models/users"
	"github.com/HoloPirates/mogupantsu/templates"
	"github.com/gin-gonic/gin"
)

// IndexModPanel : Controller for showing index page of Mod Panel
func IndexModPanel(c *gin.Context) {
	offset := 10
	torrents, _, _ := torrents.FindAllForAdminsOrderBy("torrent_id DESC", offset, 0)
	users, _ := users.FindUsersForAdmin(offset, 0)
	comments, _ := comments.FindAll(offset, 0, "", "")
	torrentReports, _, _ := reports.GetAll(offset, 0)

	templates.PanelAdmin(c, torrents, models.TorrentReportsToJSON(torrentReports), users, comments)
}

func GuidelinesModPanel(c *gin.Context) {
	templates.Static(c, "admin/guidelines.jet.html")
}
