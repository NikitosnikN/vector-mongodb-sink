package common

import "time"

type LogRecord struct {
	Timestamp time.Time   `bson:"timestamp" json:"timestamp"`
	Message   interface{} `bson:"message" json:"message"`
}
