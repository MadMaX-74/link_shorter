package statistic

type StatisticResponse struct {
	Period string `json:"period"`
	Sum    int    `json:"sum"`
}
