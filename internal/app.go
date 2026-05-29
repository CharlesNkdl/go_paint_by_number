package internal

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/charlesNkdl/go_paint_by_number/internal/config"
	img "github.com/charlesNkdl/go_paint_by_number/internal/image_processing"
	"github.com/charlesNkdl/go_paint_by_number/internal/server"
	"github.com/charlesNkdl/go_paint_by_number/internal/utils"
)

type App struct {
	server *server.Server
	config *config.Config
}

func NewApp() *App {
	return &App{
		server: server.NewServer(),
		config: config.NewConfig(),
	}
}

func Run() error {
	app := NewApp()
	go func() error {
		err := app.server.Start()
		if err != nil {
			return err
		}
		return nil
	}()
	// here is for logic testing
	fmt.Println("Test area", app.config.ImagePath)
	app.LogicTesting()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down...")
	return nil
}

func (a *App) LogicTesting() error {
	handler := img.ImageHandler{}
	imgOpened, err := handler.Open(a.config.ImagePath)
	if err != nil {
		return err
	}
	utils.PrintTypeAndKind(imgOpened)
	return nil
}
