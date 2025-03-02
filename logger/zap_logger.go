package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v7"
)

type AppLogger interface {
	Info(msg string, log *RequestLog)
	Error(msg string, log *RequestLog)
	Debug(msg string, log *RequestLog)
}

type ElasticLogger struct {
	logger *logrus.Logger
}

func NewElasticLogger(esURL, index string) (*ElasticLogger, error) {
	client, err := elastic.NewClient(elastic.SetURL(esURL), elastic.SetSniff(false))
	if err != nil {
		return nil, fmt.Errorf("error connecting to Elasticsearch: %v", err)
	}

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{TimestampFormat: time.RFC3339})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.DebugLevel)

	hook, err := elogrus.NewAsyncElasticHook(client, esURL, logrus.DebugLevel, index)
	if err != nil {
		return nil, fmt.Errorf("error setting up Elasticsearch hook: %v", err)
	}
	logger.AddHook(hook)

	return &ElasticLogger{logger: logger}, nil
}

func (e *ElasticLogger) Info(msg string, log *RequestLog) {
	data := requestLogToMap(log)
	e.logger.WithFields(data).Info(msg)
}

func (e *ElasticLogger) Error(msg string, log *RequestLog) {
	data := requestLogToMap(log)
	e.logger.WithFields(data).Error(msg)
}

func (e *ElasticLogger) Debug(msg string, log *RequestLog) {
	data := requestLogToMap(log)
	e.logger.WithFields(data).Debug(msg)
}

func requestLogToMap(log *RequestLog) map[string]interface{} {
	if log == nil {
		log = &RequestLog{}
	}

	result := make(map[string]interface{})
	// Directly populate the map with field values
	result["trace_id"] = log.TraceID
	result["method"] = log.Method
	result["path"] = log.Path
	result["query_string"] = log.QueryString
	result["user_agent"] = log.UserAgent
	result["ip_address"] = log.IpAddress
	result["response_status_code"] = log.ResponseStatusCode
	result["response_time_ms"] = log.ResponseTimeMs
	result["user_id"] = log.UserId
	result["error"] = log.Error
	result["time"] = log.Time

	return result
}
