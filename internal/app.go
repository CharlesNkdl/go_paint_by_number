package internal

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
	"github.com/charlesNkdl/go_paint_by_number/internal/calculation"
	"github.com/charlesNkdl/go_paint_by_number/internal/config"
	"github.com/charlesNkdl/go_paint_by_number/internal/image_processing"
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
	imgHandler := image_processing.ImageHandler{}
	imgOpened, err := imgio.Open(a.config.ImagePath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	resized := transform.Resize(imgOpened, 1600, 900, transform.Linear)
	utils.PrintTypeAndKind(imgOpened)
	pixels := imgHandler.ExtractPixels(resized)
	// change to flag
	numberOfColors, limitIter := 16, 50
	km := calculation.NewKMeans(numberOfColors, limitIter)
	km.Fit(pixels)
	//quantized := km.Quantize(resized)
	if err := imgio.Save("output.png", resized, imgio.PNGEncoder()); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
