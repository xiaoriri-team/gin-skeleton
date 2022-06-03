package global

import (
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/xiaoriri-team/gin-skeleton/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	RedisSetting    *setting.RedisSettingS
	AliossSetting   *setting.AliossSettingS
	JWTSetting      *setting.JWTSettingS
	LoggerSetting   *setting.LoggerSettingS
	Logger          *logrus.Logger
	Mutex           *sync.Mutex
)
