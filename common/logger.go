package common

import (
	bitlog_common "github.com/1uvu/bitlog/pkg/common"
)

// 初始化 logger

var (
	logger    bitlog_common.Logger
	constants = map[bitlog_common.ConstantKey]string{
		bitlog_common.ROOT_DIR:   "", // TODO 从 env 读取下面的值
		bitlog_common.CLIENT_DIR: "",
		bitlog_common.LOG_DIR:    "",
		bitlog_common.CONFIG_DIR: "",
	}
)

func init() {
	bitlog_common.InitConstants(constants)
	logger = bitlog_common.GetLogger("btcd_logger", bitlog_common.GetConstants(bitlog_common.LOG_DIR))
}

func Info(format string, msg ...interface{}) {
	logger.Info(format, msg)
}

func Warn(format string, msg ...interface{}) {
	logger.Warn(format, msg)
}

func Error(format string, msg ...interface{}) {
	logger.Error(format, msg)
}

func Fatal(format string, msg ...interface{}) {
	logger.Fatal(format, msg)
}
