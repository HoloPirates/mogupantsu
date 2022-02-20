package search

import "github.com/HoloPirates/mogupantsu/models"

// TorrentCache torrent cache struct
type TorrentCache struct {
	Torrents []models.Torrent
	Count    int
}
