package statistic

import (
	"go_dev/configs"
	"go_dev/pkg/middleware"
	"net/http"
	"time"
)

const (
	FilterByDay   = "day"
	FilterByMonth = "month"
)

type StatisticHandlerDeps struct {
	Config           *configs.Config
	StatisticService *StatisticService
}

type StatisticHandler struct {
	Config           *configs.Config
	StatisticService *StatisticService
}

func NewStatisticHandler(router *http.ServeMux, deps StatisticHandlerDeps) {
	handler := &StatisticHandler{
		StatisticService: deps.StatisticService,
	}
	router.Handle("GET /statistic", middleware.IsAuth(handler.GetStatistic(), deps.Config))
}

func (handler *StatisticHandler) GetStatistic() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		from, err := time.Parse("yyyy-MM-dd", r.URL.Query().Get("from"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		to, err := time.Parse("yyyy-MM-dd", r.URL.Query().Get("to"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		by := r.URL.Query().Get("by")
		if by != FilterByDay && by != FilterByMonth {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		handler.StatisticService.GetStatistic(from, to, by)
		w.WriteHeader(http.StatusOK)
	}
}
