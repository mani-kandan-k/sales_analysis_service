package common

import "time"

type ResponseStatus string

const (
	SUCCESS ResponseStatus = "S"
	ERROR   ResponseStatus = "E"
)

func GetCurrentDateTime() string {
	return time.Now().Format(time.DateTime)
}
