package logger

// Logger setup and helper functions

import (
	"log"
)

func initLogger() log.Logger {
	return *log.Default() //temp
}
