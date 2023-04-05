package logger

import (
	"encoding/json"
	"log"

	"app/config"

	"go.uber.org/zap"
)

func NewLogger(config *config.Config) *zap.Logger {
	rawJSON := []byte(`{
		"level": "debug",
		"encoding": "json",
		"outputPaths": ["stdout", "./logs"],
		"errorOutputPaths": ["stderr"],
		"encoderConfig": {
		  "messageKey": "message",
		  "levelKey": "level",
		  "levelEncoder": "lowercase"
		}	
	  }`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		log.Fatalln(err)
	}
	logger := zap.Must(cfg.Build())
	defer logger.Sync()

	return logger
}
