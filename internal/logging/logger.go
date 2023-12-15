package logging

import (
	"context"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LogLevel holds an AtomicLevel that can be used to change the logging
// level of Default.
var LogLevel = zap.NewAtomicLevel()

// Default holds the logger returned by Logger when there is no logger in
// the context. If replacing Default with a new Logger then consider
// using &LogLevel as the LevelEnabler so that SetLevel can still be used
// to dynamically change the logging level.
var Default = NewLogger()

func NewLogger() *zap.Logger {
	return zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(zapcore.EncoderConfig{
				MessageKey:  "msg",
				LevelKey:    "level",
				TimeKey:     "ts",
				EncodeLevel: zapcore.LowercaseLevelEncoder,
				EncodeTime:  zapcore.ISO8601TimeEncoder,
			}),
			os.Stdout,
			&LogLevel,
		),
	)
}

func Error(ctx context.Context, msg string, fields ...zap.Field) {
	WithContext(ctx).Error(msg, fields...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	WithContext(ctx).Sugar().Errorf(format, args...)
}

func Errorw(ctx context.Context, msg string, args ...interface{}) {
	WithContext(ctx).Sugar().Errorw(msg, args...)
}

func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	WithContext(ctx).Warn(msg, fields...)
}

func Warnf(ctx context.Context, format string, args ...interface{}) {
	WithContext(ctx).Sugar().Warnf(format, args...)
}

func Warnw(ctx context.Context, msg string, args ...interface{}) {
	WithContext(ctx).Sugar().Warnw(msg, args...)
}

func Info(ctx context.Context, msg string, fields ...zap.Field) {
	WithContext(ctx).Info(msg, fields...)
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	WithContext(ctx).Sugar().Infof(format, args...)
}

func Infow(ctx context.Context, msg string, args ...interface{}) {
	WithContext(ctx).Sugar().Infow(msg, args...)
}

func Debug(ctx context.Context, msg string, fields ...zap.Field) {
	WithContext(ctx).Debug(msg, fields...)
}

func Debugf(ctx context.Context, format string, args ...interface{}) {
	WithContext(ctx).Sugar().Debugf(format, args...)
}

func Debugw(ctx context.Context, msg string, args ...interface{}) {
	WithContext(ctx).Sugar().Debugw(msg, args...)
}

func Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	WithContext(ctx).Fatal(msg, fields...)
}

func Fatalf(ctx context.Context, format string, args ...interface{}) {
	WithContext(ctx).Sugar().Fatalf(format, args...)
}

func Fatalw(ctx context.Context, msg string, args ...interface{}) {
	WithContext(ctx).Sugar().Fatalw(msg, args...)
}

func Panic(ctx context.Context, msg string, fields ...zap.Field) {
	WithContext(ctx).Fatal(msg, fields...)
}

func Panicf(ctx context.Context, format string, args ...interface{}) {
	WithContext(ctx).Sugar().Panicf(format, args...)
}

func Panicw(ctx context.Context, msg string, args ...interface{}) {
	WithContext(ctx).Sugar().Panicw(msg, args...)
}
