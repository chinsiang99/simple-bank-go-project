package logger

import (
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/chinsiang99/simple-bank-go-project/internal/config"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var log *logrus.Logger

// Init initializes a single structured logger
func Init(cfg *config.LogConfig) {
	// Ensure log directories exist
	dirs := []string{
		filepath.Dir(cfg.AppPath),
		filepath.Dir(cfg.ErrPath),
	}
	for _, d := range dirs {
		_ = os.MkdirAll(d, os.ModePerm)
	}

	// Base formatter (JSON for machine readability)
	formatter := &logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	}

	log = logrus.New()
	log.SetFormatter(formatter)
	log.SetLevel(parseLevel(cfg.Level))

	// Setup outputs
	if cfg.LogToFile {
		// App log rotation
		appWriter := &lumberjack.Logger{
			Filename:   cfg.AppPath,
			MaxSize:    cfg.MaxSize,
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAge,
			Compress:   cfg.Compress,
		}

		// Error log rotation
		errWriter := &lumberjack.Logger{
			Filename:   cfg.ErrPath,
			MaxSize:    cfg.MaxSize,
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAge,
			Compress:   cfg.Compress,
		}

		// Custom hook: send all logs to appWriter, errors+ to errWriter
		log.SetOutput(io.MultiWriter(os.Stdout, appWriter))
		log.AddHook(&errorFileHook{Writer: io.MultiWriter(os.Stderr, errWriter)})
	} else {
		// Console only
		log.SetOutput(os.Stdout)
	}
}

// errorFileHook sends error+ logs to error log file
type errorFileHook struct {
	Writer io.Writer
}

func (h *errorFileHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel}
}

func (h *errorFileHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	_, err = h.Writer.Write([]byte(line))
	return err
}

// parseLevel converts string → logrus.Level
func parseLevel(level string) logrus.Level {
	switch level {
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	default:
		return logrus.InfoLevel
	}
}

// --- Public API ---

func Debug(args ...interface{}) { log.Debug(args...) }
func Info(args ...interface{})  { log.Info(args...) }
func Warn(args ...interface{})  { log.Warn(args...) }
func Error(args ...interface{}) { log.Error(args...) }
func Fatal(args ...interface{}) { log.Fatal(args...) }

// --- Public API (printf-style) ---

func Debugf(format string, args ...interface{}) { log.Debugf(format, args...) }
func Infof(format string, args ...interface{})  { log.Infof(format, args...) }
func Warnf(format string, args ...interface{})  { log.Warnf(format, args...) }
func Errorf(format string, args ...interface{}) { log.Errorf(format, args...) }
func Fatalf(format string, args ...interface{}) { log.Fatalf(format, args...) }

func WithField(key string, value interface{}) *logrus.Entry {
	return log.WithField(key, value)
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return log.WithFields(fields)
}

// Extra helpers
func LogRequest(method, path, clientIP string, statusCode int, duration time.Duration) {
	log.WithFields(logrus.Fields{
		"method":      method,
		"path":        path,
		"client_ip":   clientIP,
		"status_code": statusCode,
		"duration":    duration,
		"type":        "http_request",
	}).Info("HTTP Request")
}

func LogError(err error, context map[string]interface{}) {
	fields := logrus.Fields{
		"error": err.Error(),
		"type":  "application_error",
	}
	for k, v := range context {
		fields[k] = v
	}
	log.WithFields(fields).Error("Application Error")
}
