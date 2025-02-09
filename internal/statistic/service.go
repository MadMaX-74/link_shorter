package statistic

import (
	"go_dev/pkg/event"
	"log"
	"time"
)

type StatisticServiceDeps struct {
	StatisticRepository *StatisticRepository
	EventBus            *event.EventBus
}
type StatisticService struct {
	StatisticRepository *StatisticRepository
	EventBus            *event.EventBus
}

func NewStatisticService(deps *StatisticServiceDeps) *StatisticService {
	return &StatisticService{
		StatisticRepository: deps.StatisticRepository,
		EventBus:            deps.EventBus,
	}
}
func (service *StatisticService) AddClick(id uint) {
	service.StatisticRepository.AddClick(id)
}
func (service *StatisticService) GetStatistic(from, to time.Time, by string) []StatisticResponse {
	return service.StatisticRepository.GetStatistic(from, to, by)
}
func (service *StatisticService) EventAddClick() {
	for msg := range service.EventBus.Subscribe() {
		if msg.Type == event.EventLinkVisited {
			id, ok := msg.Data.(uint)
			if !ok {
				log.Fatalln("Event Link visited error", msg.Data)
				continue
			}
			service.StatisticRepository.AddClick(id)
		}
	}
}
