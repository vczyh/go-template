package log

import (
	"os"
	"testing"
)

func TestLoggerWithStdout(t *testing.T) {
	l := New("TestLogger", "DEBUG", os.Stdout)

	l.Debug("log content")
	l.Debug("log content with key and value", "key1", "val1", "key12", "val2")
}
