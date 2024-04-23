package core

import (
	"io"
	"log/slog"
	"os"
)

type Config struct {
	CheckInterval int
	LogFileDest   string
	Logger        *slog.Logger
}

func NewConfig(interval int, logFile string) *Config {
	return &Config{CheckInterval: interval, LogFileDest: logFile}
}

func (cfg *Config) Init() error {
	logger, err := cfg.setupLogger()
	if err != nil {
		return err
	}
	cfg.Logger = logger

	return nil
}

func (cfg *Config) setupLogger() (*slog.Logger, error) {
	var logFile io.Writer = os.Stdout
	if cfg.LogFileDest != "" {
		file, err := os.OpenFile(cfg.LogFileDest, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return nil, err
		}
		// defer file.Close()
		logFile = file
	}

	jsonHandler := slog.NewJSONHandler(logFile, nil)
	logger := slog.New(jsonHandler)

	return logger, nil
}
