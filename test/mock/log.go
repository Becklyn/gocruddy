package mock

import (
	"context"
	"github.com/ao-concepts/logging"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm/logger"
	"time"
)

// Log mock
type Log struct {
	mock.Mock
}

func (l *Log) Fatal(s string, args ...interface{}) {
	l.Called(s, args)
	panic("fatal log")
}

func (l *Log) ErrFatal(err error) {
	l.Called(err)
	panic("fatal log")
}

func (l *Log) Error(s string, args ...interface{}) {
	l.Called(s, args)
}

func (l *Log) ErrError(err error) {
	l.Called(err)
}

func (l *Log) Warn(s string, args ...interface{}) {
	l.Called(s, args)
}

func (l *Log) ErrWarn(err error) {
	l.Called(err)
}

func (l *Log) Info(s string, args ...interface{}) {
	l.Called(s, args)
}

func (l *Log) ErrInfo(err error) {
	l.Called(err)
}

func (l *Log) Debug(s string, args ...interface{}) {
	l.Called(s, args)
}

func (l *Log) ErrDebug(err error) {
	l.Called(err)
}

func (l *Log) CreateGormLogger() logger.Interface {
	return l.Called().Get(0).(logger.Interface)
}
func (l *Log) Write(p []byte) (n int, err error) {
	args := l.Called(err)
	return len(p), args.Error(0)
}
func (l *Log) GetLevel() logging.Level {
	return l.Called().Get(0).(logging.Level)
}

// GormLogger gorm logger
type GormLogger struct {
	mock.Mock
}

func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	l.Called(level)
	return l
}

func (l *GormLogger) Info(ctx context.Context, s string, args ...interface{}) {
	l.Called(ctx, s, args)
}

func (l *GormLogger) Warn(ctx context.Context, s string, args ...interface{}) {
	l.Called(ctx, s, args)
}

func (l *GormLogger) Error(ctx context.Context, s string, args ...interface{}) {
	l.Called(ctx, s, args)
}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	l.Called(ctx, begin, fc, err)
}
