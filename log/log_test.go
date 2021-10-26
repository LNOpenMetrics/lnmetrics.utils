package log

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestLogLevelParsingDebugOne(t *testing.T) {
	logLevel, err := parseLogLevel("debug")
	if err != nil {
		t.Errorf("Failing with an unexpected error %s", err)
	}

	assert.Equal(t, logrus.DebugLevel, logLevel, "Return level returned it is diffirent")
}

func TestLogLevelParsingDebugTwo(t *testing.T) {
	logLevel, err := parseLogLevel("Debug")
	if err != nil {
		t.Errorf("Failing with an unexpected error %s", err)
	}

	assert.Equal(t, logrus.DebugLevel, logLevel, "Return level returned it is diffirent")
}

func TestLogLevelWrongString(t *testing.T) {
	logLevel, err := parseLogLevel("Vincet")
	if err == nil {
		t.Errorf("Failing with an unexpected error %s", err)
	}

	assert.Equal(t, logrus.InfoLevel, logLevel, "Return level returned it is diffirent")
}
