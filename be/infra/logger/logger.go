package logger

import (
	"fmt"
	"log/slog"
	"os"
	"runtime"
)

// Debug
// デバックモード時のログ出力
func Debug(msg string) {
	logger := slog.New(createJSONHandler())
	_, file, line, _ := runtime.Caller(1)
	logger.Debug(msg, "path", fmt.Sprintf("%s:%d", file, line))
}

// Info
// infoレベルのログ出力
func Info(msg string) {
	logger := slog.New(createJSONHandler())
	_, file, line, _ := runtime.Caller(1)
	logger.Info(msg, "path", fmt.Sprintf("%s:%d", file, line))
}

func Error(msg string) {
	logger := slog.New(createJSONHandler())
	_, file, line, _ := runtime.Caller(1)
	logger.Error(msg, "path", fmt.Sprintf("%s:%d", file, line))
}

func createJSONHandler() *slog.JSONHandler {
	ops := slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	h := slog.NewJSONHandler(os.Stdout, &ops)

	return h
}