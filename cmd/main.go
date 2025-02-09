package main

import (
	"fmt"
	"go_dev/configs"
	"go_dev/internal/auth"
	"go_dev/internal/link"
	"go_dev/internal/statistic"
	"go_dev/internal/user"
	"go_dev/pkg/db"
	"go_dev/pkg/event"
	"go_dev/pkg/middleware"
	"net/http"
)

func App() http.Handler {
	conf := configs.LoadConfig()
	dbase := db.NewDb(conf)
	router := http.NewServeMux()
	eventBus := event.NewEventBus()
	// repositories
	linkRepository := link.NewLinkRepository(dbase)
	userRepository := user.NewUserRepository(dbase)
	statisticRepository := statistic.NewStatisticRepository(dbase)

	//services
	authService := auth.NewAuthService(userRepository)
	statisticService := statistic.NewStatisticService(&statistic.StatisticServiceDeps{
		EventBus:            eventBus,
		StatisticRepository: statisticRepository,
	})

	//handlers
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
		EventBus:       eventBus,
		Config:         conf,
	})
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	statistic.NewStatisticHandler(router, statistic.StatisticHandlerDeps{
		Config:           conf,
		StatisticService: statisticService,
	})
	go statisticService.EventAddClick()

	//middlewares
	stack := middleware.Chain(
		middleware.Cors,
		middleware.Logging,
	)
	return stack(router)
}

func main() {
	app := App()

	server := http.Server{
		Addr:    ":8081",
		Handler: app,
	}
	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
