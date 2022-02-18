package uploadController

import "github.com/HoloPirates/mogupantsu/controllers/router"

func init() {
	router.Get().Any("/upload", UploadHandler)
	router.Get().Any("/upload/status/:id", multiUploadStatus)
}
