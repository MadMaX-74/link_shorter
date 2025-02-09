package statistic

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Statistic struct {
	gorm.Model
	LinkId uint           `json:"link_id"`
	Clicks uint           `json:"clicks"`
	Data   datatypes.Date `json:"data"`
}
