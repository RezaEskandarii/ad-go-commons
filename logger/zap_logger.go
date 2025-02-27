package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/RezaEskandarii/ad-go-commons/env_manager"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type AppLogger interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Fatal(msg string, args ...interface{})
}

type Logger struct {
	sugar *zap.SugaredLogger
}

func NewLogger(index string) AppLogger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), esWriter(index)),
		zap.DebugLevel,
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	return &Logger{sugar: logger.Sugar()}
}

// esWriter
func esWriter(index string) zapcore.WriteSyncer {
	client, err := elastic.NewClient(elastic.SetURL(env_manager.Load("elasticsearch_url")), elastic.SetSniff(false))
	if err != nil {
		fmt.Println("Failed to connect to Elasticsearch:", err)
		return zapcore.AddSync(os.Stdout)
	}
	return zapcore.AddSync(&elasticWriter{client: client, index: index})
}

// elasticWriter write logs in Elasticsearch
type elasticWriter struct {
	client *elastic.Client
	index  string
}

func (w *elasticWriter) Write(p []byte) (n int, err error) {
	if w.client == nil {
		return 0, fmt.Errorf("elasticsearch client is nil")
	}

	var logEntry RequestLog
	var logResult interface{}

	err = json.Unmarshal(p, &logEntry)
	if err != nil {
		fmt.Println("Failed to unmarshal log entry:", err)
		logResult = string(p)

	} else {
		logResult = logEntry
	}

	_, err = w.client.Index().
		Index(w.index).
		BodyJson(logResult).
		Do(context.Background())

	if err != nil {
		fmt.Println("Failed to log to Elasticsearch:", err)
		return 0, err
	}

	return len(p), nil
}

func (w *elasticWriter) Sync() error {
	return nil
}

func (l *Logger) Debug(msg string, args ...interface{}) {
	l.sugar.Debugf(msg, args...)
}

func (l *Logger) Info(msg string, args ...interface{}) {
	l.sugar.Infof(msg, args...)
}

func (l *Logger) Warn(msg string, args ...interface{}) {
	l.sugar.Warnf(msg, args...)
}

func (l *Logger) Error(msg string, args ...interface{}) {
	l.sugar.Errorf(msg, args...)
}

func (l *Logger) Fatal(msg string, args ...interface{}) {
	l.sugar.Fatalf(msg, args...)
}
