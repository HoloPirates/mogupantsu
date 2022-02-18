package databasedumpsController

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/HoloPirates/mogupantsu/models"
	"github.com/HoloPirates/mogupantsu/templates"
	"github.com/HoloPirates/mogupantsu/utils/log"
	"github.com/HoloPirates/mogupantsu/utils/metainfo"
	"github.com/gin-gonic/gin"
)

const (
	// DatabaseDumpPath : Location of database dumps
	DatabaseDumpPath = "./public/dumps"
	// GPGPublicKeyPath : Location of the GPG key
	GPGPublicKeyPath = "./public/gpg/gpg.key"
)

// DatabaseDumpHandler : Controller for getting database dumps
func DatabaseDumpHandler(c *gin.Context) {
	// db params url
	// TODO Use config from cli
	files, _ := filepath.Glob(filepath.Join(DatabaseDumpPath, "*.torrent"))

	var dumpsJSON []models.DatabaseDumpJSON
	// TODO Filter *.torrent files
	for _, f := range files {
		// TODO Use config from cli
		file, err := os.Open(f)
		if err != nil {
			continue
		}
		var tf metainfo.TorrentFile
		err = tf.Decode(file)
		if err != nil {
			log.CheckError(err)
			fmt.Println(err)
			continue
		}
		dump := models.DatabaseDump{
			Date:        time.Now(),
			Filesize:    int64(tf.TotalSize()),
			Name:        tf.TorrentName(),
			TorrentLink: "/dbdumps/" + file.Name()}
		dumpsJSON = append(dumpsJSON, dump.ToJSON())
	}

	templates.DatabaseDump(c, dumpsJSON, "/gpg/gpg.pub")
}
