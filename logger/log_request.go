package logger

import "time"

type RequestLog struct {
	TraceID            string    `json:"trace_id"`
	Method             string    `json:"method"`
	Path               string    `json:"path"`
	QueryString        string    `json:"query_string"`
	UserAgent          string    `json:"user_agent"`
	IpAddress          string    `json:"ip_address"`
	ResponseStatusCode int       `json:"response_status_code"`
	ResponseTimeMs     int64     `json:"response_time_ms"`
	UserId             string    `json:"user_id"`
	Error              string    `json:"error"`
	Time               time.Time `json:"time"`
}
