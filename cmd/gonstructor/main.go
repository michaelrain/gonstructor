package main

import (
	"context"
	"errors"
	"fmt"
	"gonstructor/internal/configs"
	"gonstructor/internal/domain"
	"gonstructor/internal/repository"
	"gonstructor/internal/scenario"
	"gonstructor/internal/sources"
	"gonstructor/internal/system"
	"os"
	"os/signal"
	"syscall"

	"github.com/rosedblabs/rosedb/v2"
	"github.com/sirupsen/logrus"
)

const (
	errParseScenario = "error reading scenario"
)

func main() {

	path := os.Getenv("CONFIG_PATH")
	logger := logrus.New()

	if path == "" {
		path = "."
	}

	cfg, err := configs.NewAppConfig(path)

	if err != nil {
		return
	}

	botScenrioPath := os.Args[1]

	a, err := scenario.ParseActions(botScenrioPath)

	if err != nil {
		logger.Error(errors.Join(errors.New(errParseScenario), err))
		return
	}

	logLVL, err := logrus.ParseLevel(cfg.LogLevel)

	if err != nil {
		logger.Errorf("invalid log level %s", cfg.LogLevel)
	}

	logger.SetLevel(logLVL)

	options := rosedb.DefaultOptions
	options.DirPath = cfg.RoseDB.Path

	db, err := rosedb.Open(options)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	roseDB := repository.NewRoseDBStateRepository(db)

	if err != nil {
		fmt.Println(err.Error())
	}

	responders := make(map[string]domain.SourceResponder)

	reqeustCh := make(chan domain.Request)

	source, responder, err := sources.NewTGResource(cfg.TG, reqeustCh)

	if err != nil {
		println(err.Error())
	}

	responders["tg"] = responder

	sys, err := system.NewSystem(roseDB, a, responders, reqeustCh, logger)

	if err != nil {
		println(err.Error())
	}

	ctx := context.Background()

	go source.Listen()
	go sys.Listen(ctx)

	if err != nil {
		println(err.Error())
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	func() {
		for {
			select {
			case <-shutdown:
				logger.Trace("bye")
				return
			}
		}
	}()

	/*cli := sources.CLI{}
	cli.LoadSystem(sys)
	cli.Listen()*/
}
