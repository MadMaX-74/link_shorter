package statistic

import (
	"go_dev/pkg/db"
	"time"

	"gorm.io/datatypes"
)

type StatisticRepository struct {
	*db.Db
}

func NewStatisticRepository(database *db.Db) *StatisticRepository {
	return &StatisticRepository{
		Db: database,
	}
}

func (repository *StatisticRepository) AddClick(linkId uint) {
	var statistic Statistic
	currentDate := time.Now()
	repository.Db.Find(&statistic, "link_id = ?", linkId, datatypes.Date(currentDate))
	if statistic.ID == 0 {
		repository.Db.Create(&Statistic{
			LinkId: linkId,
			Clicks: 1,
			Data:   datatypes.Date(currentDate),
		})
	} else {
		statistic.Clicks += 1
		repository.Db.Save(&statistic)
	}
}
func (repository *StatisticRepository) GetStatistic(from, to time.Time, by string) []StatisticResponse {
	var statisticList []StatisticResponse
	var selectQuery string
	switch by {
	case FilterByDay:
		selectQuery = "to_char(data, 'YYYY-MM-DD') as period, sum(clicks) as sum"
	case FilterByMonth:
		selectQuery = "to_char(data, 'YYYY-MM') as period, sum(clicks) as sum"
	default:
		return nil
	}
	repository.Db.
		Table("statistic").
		Select(selectQuery).
		Where("data Between ? and ?", from, to).
		Group("period").
		Scan(&statisticList)
	return statisticList
}
