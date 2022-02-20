package tags

import "github.com/HoloPirates/mogupantsu/models"

func FindAll(tagType string, torrentID uint) ([]models.Tag, error) {
	tags := []models.Tag{}
	if err := models.ORM.Where("type = ? AND torrent_id = ?", tagType, torrentID).Find(&tags).Error; err != nil {
		return tags, err
	}
	return tags, nil
}
