package log

import (
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
	"time"
)

var(
	Logger zerolog.Logger
)

func InitLog(){
	Logger=zerolog.New(&lumberjack.Logger{
		Filename: viper.GetString("log.path"),
		MaxSize: viper.GetInt("log.filemaxsize"),
		MaxBackups: viper.GetInt("filemaxbackups"),
		MaxAge: viper.GetInt("filemaxage"),
		Compress: true,
	}).With().Timestamp().Logger()

	zerolog.TimeFieldFormat=time.RFC3339
	logLevel,_:=zerolog.ParseLevel(viper.GetString("log.level"))
	zerolog.SetGlobalLevel(logLevel)

	Logger.Info().Msg("log init success")

}
