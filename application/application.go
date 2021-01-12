package application

import (
	"encoding/json"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var storage *gorm.DB
var logger *zap.Logger

func Logger() *zap.Logger {
	return logger
}

func InitLogger() {
	rawJSON := []byte(`{
		"level": "debug",
		"encoding": "json",
		"outputPaths": ["stdout", "./logs"],
		"errorOutputPaths": ["stderr", "./logs"],
		"encoderConfig": {
			"messageKey": "message",
			"levelKey": "level",
			"levelEncoder": "lowercase",
			"timeKey": "time",
      "timeEncoder": "SyslogTimeEncoder",
      "CallerKey": "caller",
      "callerEncoder": "zapcore.FullCallerEncoder"
		}
	}`)
	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		Logger().Info(err.Error())
	}
	var er error
	logger, er = cfg.Build()
	if er != nil {
		Logger().Info(er.Error())
	}
	defer logger.Sync()
}

func Storage() *gorm.DB {
	return storage
}

func InitDB() {
	//TODO: properties should come from config
	dsn := "host=localhost user=memories_user password=memories_password dbname=memories_dev port=5432 sslmode=disable"
	var err error
	storage, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		Logger().Info(err.Error())
	}
}
