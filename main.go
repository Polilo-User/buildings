package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Polilo-User/buildings/services/apartaments"
	auth "github.com/Polilo-User/buildings/services/authorization"
	"github.com/Polilo-User/buildings/services/buildings"
	"github.com/Polilo-User/buildings/services/favorites"
	news "github.com/Polilo-User/buildings/services/news"

	"github.com/Polilo-User/buildings/functions/logging"
	authsvc "github.com/Polilo-User/buildings/services/authorization/implementation"
	authrepo "github.com/Polilo-User/buildings/services/authorization/repository"
	authtransport "github.com/Polilo-User/buildings/services/authorization/transport"

	apartsvc "github.com/Polilo-User/buildings/services/apartaments/implementation"
	apartrepo "github.com/Polilo-User/buildings/services/apartaments/repository"
	aparttransport "github.com/Polilo-User/buildings/services/apartaments/transport"

	buildingssvc "github.com/Polilo-User/buildings/services/buildings/implementation"
	buildingsrepo "github.com/Polilo-User/buildings/services/buildings/repository"
	buildingstransport "github.com/Polilo-User/buildings/services/buildings/transport"

	favoritessvc "github.com/Polilo-User/buildings/services/favorites/implementation"
	favoritesrepo "github.com/Polilo-User/buildings/services/favorites/repository"
	favoritestransport "github.com/Polilo-User/buildings/services/favorites/transport"

	newssvc "github.com/Polilo-User/buildings/services/news/implementation"
	newsrepo "github.com/Polilo-User/buildings/services/news/repository"
	newstransport "github.com/Polilo-User/buildings/services/news/transport"

	"github.com/Polilo-User/buildings/config"
	postgresdb "github.com/Polilo-User/buildings/functions/postgresDB"
	_ "github.com/jackc/pgx/v5/stdlib"
	//_ "github.com/go-sql-driver/mysql" // драйвер - расширение библиотеки sql для работы с mysql
)

// @title github.com/Polilo-User/buildings Microservice Server Documentation
// @version 0.1
// @description github.com/Polilo-User/buildings API
// @termsOfService http://swagger.io/terms/

func main() {
	//Инициализируем логгер
	logging.Init(false)
	logger := logging.GetLogger()

	logger.Info("Services started")
	defer logger.Info("Services ended")

	// Создаем коннект к БД
	cfg := config.GetConfig()
	db := postgresdb.NewClient(cfg.Repository)

	var authSvc auth.AuthService
	{
		repository := authrepo.New(db, logger)
		authSvc = authsvc.NewService(repository, logger)
	}

	var apartSvc apartaments.ApartService
	{
		repository := apartrepo.New(db, logger)
		apartSvc = apartsvc.NewService(repository, logger)
	}

	var buildingsSvc buildings.BuildingsService
	{
		repository := buildingsrepo.New(db, logger)
		buildingsSvc = buildingssvc.NewService(repository, logger)
	}

	var favoritesSvc favorites.FavoritesService
	{
		repository := favoritesrepo.New(db, logger)
		favoritesSvc = favoritessvc.NewService(repository, logger)
	}

	var newsSvc news.NewsService
	{
		repository := newsrepo.New(db, logger)
		newsSvc = newssvc.NewService(repository, logger)
	}

	//Для экономии
	var authEndpoints authtransport.Endpoints
	{
		authEndpoints = authtransport.MakeEndpoints(authSvc)
	}

	var apartEndpoints aparttransport.Endpoints
	{
		apartEndpoints = aparttransport.MakeEndpoints(apartSvc)
	}

	var buildingsEndpoints buildingstransport.Endpoints
	{
		buildingsEndpoints = buildingstransport.MakeEndpoints(buildingsSvc)
	}

	var favoriteEndpoint favoritestransport.Endpoints
	{
		favoriteEndpoint = favoritestransport.MakeEndpoints(favoritesSvc)
	}

	var newsEndpoint newstransport.Endpoints
	{
		newsEndpoint = newstransport.MakeEndpoints(newsSvc)
	}

	// Наш внешний роутер. Будет перенаправлять в нужный сервис, в нужный эндпоинт
	mux := http.NewServeMux()

	mux.Handle("/authorization/", authtransport.NewService(authEndpoints, logger))
	mux.Handle("/apartaments/", aparttransport.NewService(apartEndpoints, logger))
	mux.Handle("/buildings/", buildingstransport.NewService(buildingsEndpoints, logger))
	mux.Handle("/favorites/", favoritestransport.NewService(favoriteEndpoint, logger))
	mux.Handle("/news/", newstransport.NewService(newsEndpoint, logger))
	// mux.Handle("/swagger/", httpSwagger.WrapHandler)

	errs := make(chan error)
	// Зафигачим получение сигналов, как я понял будем ловить сигналы(ошибки) от UNIX'а
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	// Запустим наш сервер
	go func() {
		//level.Info(logger).Log("transport", "HTTP", "addr", *httpAddr)
		server := &http.Server{
			Addr:    cfg.Listen.Port, //*httpAddr,
			Handler: CORS(mux),
		}
		errs <- server.ListenAndServe()
	}()
	err := <-errs
	if err != nil {
		logger.Fatal(<-errs)
	}
}

func CORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		w.Header().Set("Access-Control-Allow-Origin", origin)
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			return
		} else {
			h.ServeHTTP(w, r)
		}
	})
}
