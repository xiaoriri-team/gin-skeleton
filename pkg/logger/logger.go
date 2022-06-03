package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/xiaoriri-team/gin-skeleton/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
)

func New(s *setting.LoggerSettingS) (*logrus.Logger, error) {
	log := logrus.New()
	log.Formatter = &logrus.JSONFormatter{}

	switch s.LogType {
	case setting.LogFileType:
		log.Out = &lumberjack.Logger{
			Filename:  s.LogFileSavePath + "/" + s.LogFileName + s.LogFileExt,
			MaxSize:   600,
			MaxAge:    10,
			LocalTime: true,
		}
	}

	return log, nil
}
