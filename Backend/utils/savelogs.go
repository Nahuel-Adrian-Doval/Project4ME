package utils

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func SaveLog(message string) {
	date := time.Now().Format("2006-01-02")
	logDir := "logs/" + date

	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err = os.MkdirAll(logDir, 0755)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to create directory")
			panic("Failed to create directory")
		}
	}

	logFileName := time.Now().Format("2006-01-02T15-04-05") + ".log"
	logFile, err := os.Create(logDir + "/" + logFileName)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create log file")
		panic("Failed to create log file")
	}

	// Asegúrate de cerrar el archivo al final
	defer logFile.Close()

	log.Logger = zerolog.New(logFile).With().Timestamp().Logger().Level(zerolog.FatalLevel)

	log.Info().Msg(message)
}
