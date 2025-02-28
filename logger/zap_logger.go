package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	elogrus "gopkg.in/sohlich/elogrus.v7"
)

type AppLogger interface {
	Info(msg string, jsonString string)
	Error(msg string, jsonString string)
	Debug(msg string, jsonString string)
}

// ElasticLogger
type ElasticLogger struct {
	logger *logrus.Logger
}

// NewElasticLogger
func NewElasticLogger(esURL, index string) (*ElasticLogger, error) {
	client, err := elastic.NewClient(elastic.SetURL(esURL), elastic.SetSniff(false))
	if err != nil {
		return nil, fmt.Errorf("error connecting to Elasticsearch: %v", err)
	}

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{TimestampFormat: time.RFC3339})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.DebugLevel)

	hook, err := elogrus.NewAsyncElasticHook(client, "localhost", logrus.DebugLevel, index)
	if err != nil {
		return nil, fmt.Errorf("error setting up Elasticsearch hook: %v", err)
	}
	logger.AddHook(hook)

	return &ElasticLogger{logger: logger}, nil
}

func parseJSONString(jsonString string) (map[string]interface{}, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (e *ElasticLogger) Info(msg string, jsonString string) {
	data, err := parseJSONString(jsonString)
	if err != nil {
		e.logger.WithFields(logrus.Fields{"error": err.Error()}).Error("Failed to parse JSON input")
		return
	}
	e.logger.WithFields(data).Info(msg)
}

func (e *ElasticLogger) Error(msg string, jsonString string) {
	data, err := parseJSONString(jsonString)
	if err != nil {
		e.logger.WithFields(logrus.Fields{"error": err.Error()}).Error("Failed to parse JSON input")
		return
	}
	e.logger.WithFields(data).Error(msg)
}

func (e *ElasticLogger) Debug(msg string, jsonString string) {
	data, err := parseJSONString(jsonString)
	if err != nil {
		e.logger.WithFields(logrus.Fields{"error": err.Error()}).Error("Failed to parse JSON input")
		return
	}
	e.logger.WithFields(data).Debug(msg)
}
