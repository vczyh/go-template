package log

import (
	"os"
	"testing"
)

func TestLoggerWithStdout(t *testing.T) {
	logger := NewLogger("TestLogger", "DEBUG", os.Stdout)

	logger.Debug("log content")
	logger.Debug("log content with key and value", "key", "val")
}
