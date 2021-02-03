package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/euwas/rebootmadgo/internal/config"
	"github.com/euwas/rebootmadgo/pkg/controller"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if true {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}

	log.Logger = zerolog.New(output).With().Timestamp().Logger()
	log.Info().Str("foo", "bar").Msg("Hello world")

	log.Info().Msg("Starting RebootMadGo 🚀")

	config, err := config.ReadConfig()
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Cannot load configuration, does config.yaml exist?")
	}
	log.Debug().Object("config", config).Msg("Loaded config")

	dw := controller.DeviceWatcherFromConfig(config)

	log.Debug().Int("device_count", dw.Length()).Msg("Loaded all devices")
	dw.Update()
}
