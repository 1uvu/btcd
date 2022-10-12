package common

import (
	bitlog_types "github.com/1uvu/bitlog/parser/types"
)

// 初始化和管理 bitlog 的三种 Log

// chain 和 status 的全局实例，实时获取和更新，发生更改时替换，
// event 默认每次都会创建新的，无需全局实例
var (
	txChangeLog      *bitlog_types.ChangeLog
	blockChangeLog   *bitlog_types.ChangeLog
	chainChangeLog   *bitlog_types.ChangeLog
	networkChangeLog *bitlog_types.ChangeLog

	chainStatusLog   *bitlog_types.StatusLog
	networkStatusLog *bitlog_types.StatusLog
)

func NewRawLog(logType bitlog_types.RawLogType, raw []byte) *bitlog_types.RawLog {
	return nil
}

// NewEventLog 根据当前 status 实例生成新的 event
func NewEventLog(rawLog *bitlog_types.RawLog) *bitlog_types.EventLog {
	return nil
}

// NewChangeLog 传入的参数生成新的 change
func NewChangeLog(rawLog *bitlog_types.RawLog) *bitlog_types.ChangeLog {
	return nil
}

// GetChangeLog 获得当前 change
func GetChangeLog(changeLogType bitlog_types.RawLogType) *bitlog_types.ChangeLog {
	return nil
}

// UpdateChangeLog 传入参数更新当前 change
func UpdateChangeLog(rawLog *bitlog_types.RawLog) *bitlog_types.ChangeLog {
	return nil
}

// NewStatusLog 传入的参数生成新的 status
func NewStatusLog(rawLog *bitlog_types.RawLog) *bitlog_types.StatusLog {
	return nil
}

// GetStatusLog 获得当前 status
func GetStatusLog(statusLogType bitlog_types.RawLogType) *bitlog_types.StatusLog {
	return nil
}

// UpdateStatusLog 传入参数更新当前 status
func UpdateStatusLog(rawLog *bitlog_types.RawLog) *bitlog_types.StatusLog {
	return nil
}
